package c2q55

import (
	"sync/atomic"
	"unsafe"
)

// États atomiques d'un créneau dans l'anneau
const (
	SlotFree      uint32 = 0
	SlotReady     uint32 = 1
	SlotReserved  uint32 = 2
	SlotCommitted uint32 = 3
	SlotAcked     uint32 = 4
	SlotPoison    uint32 = 5
	SlotDLQ       uint32 = 6
)

// Slot représente un créneau unitaire de 64 octets exacts (1 ligne de cache L1).
// Aucun pointeur sur le tas n'est présent (structure invisible pour le GC).
type Slot struct {
	IDLow       uint64        // 0..7   - UUIDv7 partie basse
	IDHigh      uint64        // 8..15  - UUIDv7 partie haute (horodatage monotone)
	VisibleAtNs uint64        // 16..23 - Horodatage de visibilité (TDMA)
	LeaseExpire atomic.Uint64 // 24..31 - Expiration du bail de traitement
	TopicID     uint32        // 32..35 - Identifiant du sujet
	State       atomic.Uint32 // 36..39 - État atomique du créneau
	PayloadLen  uint32        // 40..43 - Longueur effective de la charge utile (0..16)
	Flags       uint32        // 44..47 - Compteur de retries et options
	Payload     [16]byte      // 48..63 - Charge utile inlinée 0-allocation
}

// Invariants matériels ARCHTIME bidirectionnels : taille strictement égale à 64 octets.
type _slotSizeMin [64 - unsafe.Sizeof(Slot{})]byte
type _slotSizeMax [unsafe.Sizeof(Slot{}) - 64]byte

// Pack initialise les champs d'un créneau de façon déterministe.
func (s *Slot) Pack(idLow, idHigh, visibleAtNs, leaseExpireNs uint64, topicID, flags uint32, payload []byte, state uint32) {
	s.IDLow = idLow
	s.IDHigh = idHigh
	s.VisibleAtNs = visibleAtNs
	s.LeaseExpire.Store(leaseExpireNs)
	s.TopicID = topicID
	s.Flags = flags

	pLen := len(payload)
	if pLen > 16 {
		pLen = 16
	}
	s.PayloadLen = uint32(pLen)
	for i := 0; i < 16; i++ {
		if i < pLen {
			s.Payload[i] = payload[i]
		} else {
			s.Payload[i] = 0
		}
	}

	s.State.Store(state)
}

// CopyTo transfère les champs vers un autre slot sans copier de verrou noCopy.
func (s *Slot) CopyTo(dst *Slot) {
	dst.IDLow = s.IDLow
	dst.IDHigh = s.IDHigh
	dst.VisibleAtNs = s.VisibleAtNs
	dst.LeaseExpire.Store(s.LeaseExpire.Load())
	dst.TopicID = s.TopicID
	dst.Flags = s.Flags
	dst.PayloadLen = s.PayloadLen
	dst.Payload = s.Payload
	dst.State.Store(s.State.Load())
}
