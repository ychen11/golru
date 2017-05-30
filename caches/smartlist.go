package caches

import (
    "container/list"
)


type element struct {
	// time to destroy
	ttd int64
	value [] interface{}
}

type SmartList struct {
	li *list.List
}

func NewSmartList() *SmartList {
	return &SmartList{
		li: list.New(),
	}
}

func (sl *SmartList) SimpleInsert() {
	
}
