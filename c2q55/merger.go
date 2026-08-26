// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

// LoserTreeMerger implémente un véritable arbre des perdants (Tournament Tree) en O(log K).
// La structure d'arbre reste confinée en cache L1D (< 10 Ko pour K <= 64).
type LoserTreeMerger struct {
	k       int
	shards  []*Ring
	heads   []Slot
	hasHead []bool
	tree    []int // tree[0] contient le gagnant absolu, tree[1..k-1] stocke les perdants intermédiaires
}

// NewLoserTreeMerger initialise l'arbre pour K shards partitionnés.
func NewLoserTreeMerger(shards []*Ring) *LoserTreeMerger {
	k := len(shards)
	if k == 0 {
		return nil
	}

	m := &LoserTreeMerger{
		k:       k,
		shards:  shards,
		heads:   make([]Slot, k),
		hasHead: make([]bool, k),
		tree:    make([]int, k),
	}

	return m
}

// Init remplit les têtes de chaque shard et bâtit l'arbre initial en O(K).
func (m *LoserTreeMerger) Init(nowNs uint64) {
	for i := 0; i < m.k; i++ {
		var s Slot
		if seq, ok := m.shards[i].Dequeue(nowNs, &s); ok {
			_ = m.shards[i].Ack(seq)
			s.CopyTo(&m.heads[i])
			m.hasHead[i] = true
		} else {
			m.hasHead[i] = false
		}
	}

	for i := 0; i < m.k; i++ {
		m.tree[i] = -1
	}

	for i := m.k - 1; i >= 0; i-- {
		m.replay(i)
	}
}

// less compare deux têtes selon l'ordre strict monotone de l'UUIDv7 (IDHigh puis IDLow).
func (m *LoserTreeMerger) less(i, j int) bool {
	if !m.hasHead[i] {
		return false
	}
	if !m.hasHead[j] {
		return true
	}
	if m.heads[i].IDHigh != m.heads[j].IDHigh {
		return m.heads[i].IDHigh < m.heads[j].IDHigh
	}
	return m.heads[i].IDLow < m.heads[j].IDLow
}

// replay fait remonter un candidat le long de la branche de l'arbre en O(log K).
func (m *LoserTreeMerger) replay(player int) {
	parent := (player + m.k) / 2

	for parent > 0 {
		opponent := m.tree[parent]
		if opponent == -1 {
			m.tree[parent] = player
			return
		}

		if m.less(opponent, player) {
			m.tree[parent] = player
			player = opponent
		}
		parent /= 2
	}

	m.tree[0] = player
}

// Next extrait le message minimum global en O(log K) et réapprovisionne le shard gagnant.
func (m *LoserTreeMerger) Next(nowNs uint64, out *Slot) bool {
	winner := m.tree[0]
	if winner == -1 || !m.hasHead[winner] {
		refilled := false
		for i := 0; i < m.k; i++ {
			if !m.hasHead[i] {
				var s Slot
				if seq, ok := m.shards[i].Dequeue(nowNs, &s); ok {
					_ = m.shards[i].Ack(seq)
					s.CopyTo(&m.heads[i])
					m.hasHead[i] = true
					m.replay(i)
					refilled = true
				}
			}
		}
		if !refilled {
			return false
		}
		winner = m.tree[0]
		if winner == -1 || !m.hasHead[winner] {
			return false
		}
	}

	m.heads[winner].CopyTo(out)

	// Recharger le shard gagnant
	var s Slot
	if seq, ok := m.shards[winner].Dequeue(nowNs, &s); ok {
		_ = m.shards[winner].Ack(seq)
		s.CopyTo(&m.heads[winner])
		m.hasHead[winner] = true
	} else {
		m.hasHead[winner] = false
	}

	// Remontée en O(log K)
	m.replay(winner)
	return true
}
