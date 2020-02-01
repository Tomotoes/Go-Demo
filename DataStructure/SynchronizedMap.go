package DataStructure

import "sync"

type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

func New() *SynchronizedMap {
	return &SynchronizedMap{rw: &sync.RWMutex{}, data: map[interface{}]interface{}{}}
}

func (sm *SynchronizedMap) New() *SynchronizedMap {
	return &SynchronizedMap{rw: &sync.RWMutex{}, data: map[interface{}]interface{}{}}
}

func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	return sm.data[k]
}

func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	delete(sm.data, k)
}

func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()

	for key, value := range sm.data {
		cb(key, value)
	}
}

func (sm *SynchronizedMap) Map(handle func(interface{}, interface{}) interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	for key, value := range sm.data {
		sm.data[key] = handle(key, value)
	}
}

func (sm *SynchronizedMap) Filter(filter func(interface{}, interface{}) bool) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	for key, value := range sm.data {
		if !filter(key, value) {
			delete(sm.data, key)
		}
	}
}

func (sm *SynchronizedMap) Size() int {
	return len(sm.data)
}

func (sm *SynchronizedMap) Keys() []interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	result := []interface{}{}
	for key := range sm.data {
		result = append(result, key)
	}
	return result
}

func (sm *SynchronizedMap) Values() []interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	result := []interface{}{}
	for _, value := range sm.data {
		result = append(result, value)
	}
	return result
}
