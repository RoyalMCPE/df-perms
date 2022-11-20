package session

import (
	"sync"

	"github.com/df-mc/atomic"
	"github.com/df-mc/dragonfly/server/player"
	dfperms "github.com/royalmcpe/df-perms"
)

// Session holds the groups and tags of a player
// Groups are sorted based on permission level
type Session struct {
	p atomic.Value[*player.Player]

	// TODO: Implement a binary tree that goes through based on permission level
	groups  map[string]*dfperms.Group
	groupRW sync.RWMutex
}
