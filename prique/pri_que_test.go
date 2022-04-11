package prique_test

import (
	"testing"

	"github.com/bcyxy/golibs/prique"
)

type pair struct {
	k int
	v int
}

var iCases = []pair{{9, 1}, {2, 2}, {8, 3}, {3, 4}, {7, 5}, {8, 6}}
var oCases = []pair{{9, 1}, {8, 3}, {8, 6}, {7, 5}, {3, 4}, {2, 2}}

func TestQue(t *testing.T) {
	q := prique.NewQue()
	for _, c := range iCases {
		q.Put(c.v, c.k)
		//fmt.Println(q)
	}

	for _, c := range oCases {
		val, pri, ok := q.Get()
		//fmt.Println(q)
		if !ok {
			t.Error("count_err")
		}
		if val != c.v || pri != c.k {
			t.Error("value_err")
		}
	}
}
