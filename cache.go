package cache

import (
	"time"
)

type containment struct {
	value string
	before *time.Time
}

type Cache struct {
	container map[string]containment
}

func NewCache() Cache {
	return Cache{map[string]containment{}}
}

func (c Cache) cleanup() {
	for key, val := range c.container {
		if val.before != nil && !time.Now().Before(*val.before) {
			delete(c.container, key)
		}
	}
}

func (c Cache) Get(key string) (string, bool) {
	c.cleanup()
	if val, ok := c.container[key]; ok {
		// I know that after cleanup I shouldn't check containment expiration, but still
		if val.before == nil || time.Now().Before(*val.before) {
			return val.value, true
		}
	}
	return "", false
}

func (c Cache) Put(key, value string) {
	c.cleanup()
	c.container[key] = containment{value, nil}
}

func (c Cache) Keys() []string {
	c.cleanup()
	keys := make([]string, len(c.container))
	counter := 0
	for key := range c.container {
		keys[counter] = key
		counter++
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.cleanup()
	c.container[key] = containment{value, &deadline}
}
