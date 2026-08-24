package c2q55

import (
	"errors"
	"runtime"
	"sync/atomic"
)

var (
	ErrRingFull        = errors.New("c2q55/ring: capacity exceeded (backpressure)")
	ErrRingEmpty       = errors.New("c2q55/ring: no ready messages")
	ErrSlotNotReserved = errors.New("c2q55/ring: slot is not reserved by caller")
	ErrPayloadTooLarge = errors.New("c2q55/ring: payload exceeds 16 bytes (must use SlabArena with FlagExternalPayload)")
)

// ProducerCursor regroupe le compteur d'écriture isolé sur 128 octets.
type ProducerCursor struct {
	head atomic.Uint64
	_    [120]byte
}

// ConsumerCursor regroupe le compteur de lecture isolé sur 128 octets.
type ConsumerCursor struct {
	tail atomic.Uint64
	_    [120]byte
}

// Ring représente un anneau MPMC lock-free partitionné.
type Ring struct {
	producer ProducerCursor
	consumer ConsumerCursor
	drops    atomic.Uint64
	capacity uint64
	mask     uint64
	slots    []Slot
}

// NewRing alloue un anneau dont la capacité est arrondie à la puissance de deux supérieure.
func NewRing(capacity uint64) *Ring {
	if capacity < 16 {
		capacity = 16
	}

	capPow2 := uint64(1)
	for capPow2 < capacity {
		capPow2 <<= 1
	}

	return &Ring{
		capacity: capPow2,
		mask:     capPow2 - 1,
		slots:    make([]Slot, capPow2),
	}
}

// Capacity retourne la capacité totale de l'anneau.
func (r *Ring) Capacity() uint64 {
	return r.capacity
}

// Drops retourne le nombre total de messages rejetés par contre-pression.
func (r *Ring) Drops() uint64 {
	return r.drops.Load()
}

// Enqueue réserve un créneau via CAS et publie le message de façon lock-free.
// Refuse strictement toute charge > 16 octets sans FlagExternalPayload (zéro troncature silencieuse).
func (r *Ring) Enqueue(idLow, idHigh, visibleAtNs, leaseExpireNs uint64, topicID, flags uint32, payload []byte) (uint64, error) {
	pLen := len(payload)
	if pLen > 16 && (flags&FlagExternalPayload == 0) {
		return 0, ErrPayloadTooLarge
	}

	for {
		head := r.producer.head.Load()
		tail := r.consumer.tail.Load()

		if int64(head-tail) >= int64(r.capacity) {
			r.drops.Add(1)
			return 0, ErrRingFull
		}

		if r.producer.head.CompareAndSwap(head, head+1) {
			idx := head & r.mask
			slot := &r.slots[idx]

			// Attente active bornée si le tour précédent n'est pas libéré
			spins := 0
			for slot.State.Load() != SlotFree && slot.State.Load() != SlotAcked && slot.State.Load() != SlotDLQ {
				spins++
				if spins > 30 {
					runtime.Gosched()
				}
				if spins > 10000 {
					// Protection absolue contre l'écrasement silencieux d'un slot réservé
					r.drops.Add(1)
					return 0, ErrRingFull
				}
			}

			slot.IDLow = idLow
			slot.IDHigh = idHigh
			slot.VisibleAtNs = visibleAtNs
			slot.LeaseExpire.Store(leaseExpireNs)
			slot.TopicID = topicID
			slot.Flags = flags

			if pLen > 16 {
				pLen = 16
			}
			slot.PayloadLen = uint32(pLen)
			for i := 0; i < 16; i++ {
				if i < pLen {
					slot.Payload[i] = payload[i]
				} else {
					slot.Payload[i] = 0
				}
			}

			slot.State.Store(SlotReady)
			return head, nil
		}
	}
}

// Dequeue extrait le prochain message prêt et le passe en état SlotReserved.
func (r *Ring) Dequeue(nowNs uint64, out *Slot) (uint64, bool) {
	for {
		tail := r.consumer.tail.Load()
		head := r.producer.head.Load()

		if int64(head-tail) <= 0 {
			return 0, false
		}

		idx := tail & r.mask
		slot := &r.slots[idx]

		st := slot.State.Load()
		if st != SlotReady {
			if st == SlotFree {
				return 0, false
			}
			if r.consumer.tail.CompareAndSwap(tail, tail+1) {
				continue
			}
			continue
		}

		if slot.VisibleAtNs > nowNs {
			return 0, false
		}

		if r.consumer.tail.CompareAndSwap(tail, tail+1) {
			if slot.State.CompareAndSwap(SlotReady, SlotReserved) {
				slot.CopyTo(out)
				return tail, true
			}
		}
	}
}

// LeaseExtend prolonge la durée de validité du bail de traitement.
func (r *Ring) LeaseExtend(seq uint64, additionalNs uint64) error {
	idx := seq & r.mask
	slot := &r.slots[idx]

	if slot.State.Load() != SlotReserved {
		return ErrSlotNotReserved
	}

	slot.LeaseExpire.Add(additionalNs)
	return nil
}

// Ack confirme le traitement du créneau et le libère.
func (r *Ring) Ack(seq uint64) error {
	idx := seq & r.mask
	slot := &r.slots[idx]

	st := slot.State.Load()
	if st != SlotReserved {
		return ErrSlotNotReserved
	}

	slot.State.Store(SlotAcked)
	return nil
}

// Nack signale un échec de traitement, réincrémente les retries et bascule en DLQ si maxRetries est atteint.
func (r *Ring) Nack(seq uint64, maxRetries uint32, retryDelayNs uint64, nowNs uint64) error {
	idx := seq & r.mask
	slot := &r.slots[idx]

	if slot.State.Load() != SlotReserved {
		return ErrSlotNotReserved
	}

	retries := slot.Flags & 0xFF
	if retries+1 >= maxRetries {
		slot.State.Store(SlotDLQ)
		return nil
	}

	_, err := r.Enqueue(
		slot.IDLow,
		slot.IDHigh,
		nowNs+retryDelayNs,
		nowNs+retryDelayNs+5000000,
		slot.TopicID,
		(slot.Flags&0xFFFFFF00)|(retries+1),
		slot.Payload[:slot.PayloadLen],
	)

	slot.State.Store(SlotAcked)
	return err
}
