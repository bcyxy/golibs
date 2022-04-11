package prique

import (
	"container/list"
	"fmt"
)

// Que priority queue
// Not concurrency safe
type Que struct {
	// TODO It's better to use skip list.
	priList list.List
	subQIdx map[int]*subQue
}

func NewQue() *Que {
	return &Que{
		subQIdx: make(map[int]*subQue),
	}
}

// subQue subqueue
type subQue struct {
	pri    int
	valQue list.List
}

// Put Add an element.
func (sf *Que) Put(val interface{}, pri int) {
	sq, ok := sf.subQIdx[pri]
	if !ok {
		sq = sf.addSubQue(pri)
	}
	sq.valQue.PushFront(val)
	return
}

// Get Get and delete an element.
func (sf *Que) Get() (val interface{}, pri int, ok bool) {
	ok = false
	for e := sf.priList.Front(); e != nil; e = e.Next() {
		sq := e.Value.(*subQue)
		if sq.valQue.Len() == 0 {
			continue
		}
		val = sq.valQue.Remove(sq.valQue.Back())
		pri = sq.pri
		ok = true

		// Delete empty subqueue
		if sq.valQue.Len() == 0 {
			sf.priList.Remove(e)
			delete(sf.subQIdx, sq.pri)
		}
		return
	}
	return
}

func (sf *Que) String() (s string) {
	for e := sf.priList.Front(); e != nil; e = e.Next() {
		sq := e.Value.(*subQue)
		s += fmt.Sprintf("%d:", sq.pri)
		for e2 := sq.valQue.Front(); e2 != nil; e2 = e2.Next() {
			s += fmt.Sprintf("%v,", e2.Value)
		}
		s += ";"
	}
	return
}

func (sf *Que) addSubQue(pri int) (sq *subQue) {
	defer func() {
		sf.subQIdx[pri] = sq
	}()

	var e *list.Element
	for e = sf.priList.Front(); e != nil; e = e.Next() {
		sq = e.Value.(*subQue)
		if sq.pri > pri {
			continue
		}
		if sq.pri < pri {
			sq = &subQue{pri: pri}
			e = sf.priList.InsertBefore(sq, e)
			return sq
		} else { //==
			return sq
		}
	}
	if e == nil {
		sq = &subQue{pri: pri}
		sf.priList.PushBack(sq)
		return sq
	}
	return nil
}
