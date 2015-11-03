package gotslist

import (
	"container/list"
	"sync"
)

type GoTSList struct {
	list      *list.List
	lockMutex sync.Mutex
}

func NewGoTSList() *GoTSList {
	return &GoTSList{list.New(), sync.Mutex{}}
}

func (this *GoTSList) PushBack(elem interface{}) {
	this.Lock()
	defer this.Unlock()

	if elem != nil {
		this.list.PushBack(elem)
	}
}

func (this *GoTSList) Remove(elem *list.Element) {
	this.Lock()
	defer this.Unlock()

	if elem != nil {
		this.list.Remove(elem)
	}
}

func (this *GoTSList) Pop() (interface{}, bool) {
	this.Lock()
	defer this.Unlock()

	if this.list.Len() == 0 {
		return nil, false
	}

	element := this.list.Back()
	this.list.Remove(element)

	return element.Value, true
}

func (this *GoTSList) Len() int {
	this.Lock()
	defer this.Unlock()

	return this.list.Len()
}

func (this *GoTSList) IsEmpty() bool {
	this.Lock()
	defer this.Unlock()

	return (this.list.Len() == 0)
}

func (this *GoTSList) Front() *list.Element {
	this.Lock()
	defer this.Unlock()

	return this.list.Front()
}

func (this *GoTSList) Back() *list.Element {
	this.Lock()
	defer this.Unlock()

	return this.list.Back()
}

func (this *GoTSList) Lock() {
	this.lockMutex.Lock()
}

func (this *GoTSList) Unlock() {
	this.lockMutex.Unlock()
}
