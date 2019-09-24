package main

import (
	"reflect"
	"testing"
)

var l = []struct {
	a []interface{}
	r []int
}{
	{[]interface{}{[]interface{}{1, 2, []int{3}, 4}}, []int{1, 2, 3, 4}},
	{[]interface{}{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{nil, nil},
	{[]interface{}{"x", nil, []int{3}, -1, 0}, []int{3, -1, 0}},
	{[]interface{}{[]interface{}{1, []interface{}{1, 9, 4, 4}, []int{3}, []interface{}{1, 2, "x", 4}}}, []int{1, 1, 9, 4, 4, 3, 1, 2, 4}},
}

func TestFlatten(t *testing.T) {
	for _, i := range l {
		f := Flatten(i.a)
		if !reflect.DeepEqual(f, i.r) {
			t.Errorf("Flatten(%+v) was incorrect, got: %+v, want: %+v.", i.a, f, i.r)
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Flatten(l[0].a)
	}
}
