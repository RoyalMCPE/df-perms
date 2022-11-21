package dfperms

import (
	"fmt"
	"sync"
)

var groups sync.Map

func Register(group *Group) error {
	if _, loaded := groups.LoadOrStore(group.Title(), group); loaded {
		return fmt.Errorf("Group '%v' is already registered", group.Title())
	}
	return nil
}

func ByTitle(title string) *Group {
	group, ok := groups.Load(title)
	if !ok {
		return nil
	}

	return group.(*Group)
}
