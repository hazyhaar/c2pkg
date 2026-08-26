// SPDX-License-Identifier: Apache-2.0 OR MIT

package c2q55

import (
	"errors"
	"strings"
	"sync"
)

var (
	ErrPermissionDenied = errors.New("c2q55/acl: permission denied for principal on topic")
)

// Privilege définit les droits atomiques d'accès sur le bus.
type Privilege uint8

const (
	PrivilegeNone      Privilege = 0
	PrivilegePublish   Privilege = 1 << 0 // Droit de publier (Publish)
	PrivilegeConsume   Privilege = 1 << 1 // Droit de lire (Consume)
	PrivilegeReplicate Privilege = 1 << 2 // Droit de synchroniser/fetcher (Replicate)
	PrivilegeAdmin     Privilege = 1<<3 | PrivilegePublish | PrivilegeConsume | PrivilegeReplicate
)

// ACLRule définit une règle d'autorisation liant un principal, un motif de topic et des privilèges.
type ACLRule struct {
	Principal    string    // NodeID ou CommonName du certificat mTLS (* pour tout le monde)
	TopicPattern string    // Nom de topic exact ou préfixe avec wildcard (ex: "telemetry.*", "*")
	Privileges   Privilege // Masque binaire des privilèges accordés
}

// ACLManager gère l'ensemble des règles d'accès du cluster avec évaluation à 0 allocation.
type ACLManager struct {
	mu    sync.RWMutex
	rules []ACLRule
}

// NewACLManager initialise un gestionnaire d'ACLs vierge.
func NewACLManager() *ACLManager {
	return &ACLManager{
		rules: make([]ACLRule, 0, 16),
	}
}

// AddRule ajoute une règle d'autorisation dans la table.
func (am *ACLManager) AddRule(principal, topicPattern string, privs Privilege) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.rules = append(am.rules, ACLRule{
		Principal:    principal,
		TopicPattern: topicPattern,
		Privileges:   privs,
	})
}

// CheckAccess vérifie si un principal possède le privilège requis sur un topic donné.
func (am *ACLManager) CheckAccess(principal, topic string, required Privilege) bool {
	if am == nil {
		return false
	}
	am.mu.RLock()
	defer am.mu.RUnlock()

	// Si aucune règle n'est définie, par défaut tout est refusé (fail-closed)
	if len(am.rules) == 0 {
		return false
	}

	for i := range am.rules {
		r := &am.rules[i]
		// Vérification du principal
		if r.Principal != "*" && r.Principal != principal {
			continue
		}

		// Vérification du topic pattern
		if !matchTopicPattern(r.TopicPattern, topic) {
			continue
		}

		// Vérification du privilège requis
		if (r.Privileges & required) == required {
			return true
		}
	}

	return false
}

// matchTopicPattern compare un motif de topic (supportant '*' en suffixe) avec un nom de topic.
func matchTopicPattern(pattern, topic string) bool {
	if pattern == "*" || pattern == topic {
		return true
	}
	if strings.HasSuffix(pattern, ".*") {
		prefix := pattern[:len(pattern)-2]
		return strings.HasPrefix(topic, prefix)
	}
	return false
}
