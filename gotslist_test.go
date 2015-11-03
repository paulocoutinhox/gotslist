package gotslist

import (
	"testing"
	"sync"
)

func TestNewList(t *testing.T) {
	tslist := NewGoTSList()

	if tslist == nil {
		t.Error("List is not created")
	}
}

func TestConcurrency(t *testing.T) {
	tslist := NewGoTSList()
	wg := sync.WaitGroup{}
	total := 100000

	wg.Add(total)

	for i := 0; i < total; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tslist.Push("New element")
		}(&wg)
	}

	wg.Wait()
}

func TestConcurrencyAddCount(t *testing.T) {
	tslist := NewGoTSList()
	wg := sync.WaitGroup{}
	total := 100000

	wg.Add(total)

	for i := 0; i < total; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tslist.Push("New element")
			_ = tslist.Len()
		}(&wg)
	}

	wg.Wait()
}

func TestConcurrencyAddRemove(t *testing.T) {
	tslist := NewGoTSList()
	wgAdd := sync.WaitGroup{}
	wgRemove := sync.WaitGroup{}
	totalAdd := 100000
	totalRemove := 100000
	element := "Unique Element"

	wgAdd.Add(totalAdd)
	wgRemove.Add(totalRemove)

	for i := 0; i < totalAdd; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tslist.Push(element)
		}(&wgAdd)
	}

	for i := 0; i < totalRemove; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			for e := tslist.list.Front(); e != nil; e = e.Next() {
				tslist.Remove(e)
			}
		}(&wgRemove)
	}

	wgAdd.Wait()
	wgRemove.Wait()
}

func TestConcurrencyAddPop(t *testing.T) {
	tslist := NewGoTSList()
	wgAdd := sync.WaitGroup{}
	wgPop := sync.WaitGroup{}
	totalAdd := 100000
	totalPop := 100000

	wgAdd.Add(totalAdd)
	wgPop.Add(totalPop)

	for i := 0; i < totalAdd; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tslist.Push("New element")
		}(&wgAdd)
	}

	for i := 0; i < totalPop; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			tslist.Pop()
		}(&wgPop)
	}

	wgAdd.Wait()
	wgPop.Wait()
}