package SafeMap

import "sync"

type Storage interface {
	Delete(key string)
	Get(key string) (string, bool)
	Put(key, value string)
}

type SafeMap struct {
	_map map[string]string
	sync.Mutex
}

func (sm *SafeMap) Delete(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm._map, key)
}

func (sm *SafeMap) Get(key string) (string, bool) {
	sm.Lock()
	defer sm.Unlock()
	value, ok := sm._map[key]
	return value, ok
}

func (sm *SafeMap) Put(key, value string) {
	sm.Lock()
	defer sm.Unlock()
	if sm == nil {
		sm = &SafeMap{
			_map: map[string]string{},
		}
	}
	sm._map[key] = value
}
