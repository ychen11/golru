package lrucore

import (
    caches "github.com/ychen11/golru/caches"
)

type GoLRU struct {
	DataCache *caches.SyncMap
	IndexCache *
}

func NewLRU () {
	return &GoLRU{

	}
}

func (lru *GoLRU) insert(key interface{}, value interface{}) {

}