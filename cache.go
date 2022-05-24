package cache

import (
	"time"
)

type containment struct {
	value string
	temporary bool
	before time.Time
}

type Cache struct {
	container map[string]containment
}

func NewCache() Cache {
	return Cache{map[string]containment{}}
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.container[key]; ok {
		if !val.temporary || time.Now().Before(val.before) {
			return val.value, true
		}
	}
	return "", false
}

func (c Cache) Put(key, value string) {
	c.container[key] = containment{value, false, time.Time{}}
}

func (c Cache) Keys() []string {
	keys := make([]string, len(c.container))
	counter := 0
	for key := range c.container {
		keys[counter] = key
		counter++
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.container[key] = containment{value, true, deadline}
}
