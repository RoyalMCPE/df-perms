package session

import (
	"sync"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	dfperms "github.com/royalmcpe/df-perms"
)

var man *manager

func init() {
	man = &manager{
		sessions: map[uuid.UUID]*Session{},
	}
}

type manager struct {
	sessions map[uuid.UUID]*Session
	sessMu   sync.RWMutex
}

// Accept assigns a session to a player and applies default groups.
func Accept(p *player.Player, groups ...*dfperms.Group) (*Session, error) {
	man.sessMu.Lock()
	defer man.sessMu.Unlock()

	session := &Session{}
	session.p.Store(p)

	for _, group := range groups {
		session.groups[group.Title()] = group
	}

	man.sessions[p.UUID()] = session

	return session, nil
}

// Search for a session from a player
func FromPlayer(p *player.Player) (*Session, bool) {
	return FromUUID(p.UUID())
}

// Search for a session from player uuid
func FromUUID(uuid uuid.UUID) (*Session, bool) {
	man.sessMu.RLock()
	s, ok := man.sessions[uuid]
	man.sessMu.RUnlock()
	return s, ok
}
