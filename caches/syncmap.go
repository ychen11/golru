package caches

import (
  "sync"
)

type SyncMap struct {
	lock *sync.RWMutex
	mp map[interface{}]interface{}
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		lock new(sync.RWMutex)
		mp make(map[interface{}]interface{})
	}
}

func (sm *SyncMap) Insert(key interface{}, value interface{}) {

}

func (sm *SyncMap) Get(key interface{}) interface{}{

}

func (sm *SyncMap) Delete(key interface{}) bool {

}

func (sm *SyncMap) Update(key interface{}, value interface{}) bool {
	
}