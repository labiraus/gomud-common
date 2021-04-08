package utils

import (
	"strconv"
	"strings"

	"../core"
)

//Search finds an entity in a room list
func Search(m map[uint]core.Entity, search string) core.Entity {
	id, err := strconv.Atoi(search)
	if err != nil {
		return m[uint(id)]
	}
	var current core.Entity = nil
	for _, ent := range m {
		name := ent.Name()
		if !strings.HasPrefix(name, search) {
			continue
		}
		if name == search {
			return ent
		}
		if name < current.Name() {
			current = ent
		}
	}
	return current
}
