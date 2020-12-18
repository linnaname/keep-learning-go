package main

import (
	"reflect"
)

type A struct {
	s string
}

type AS []*A

func main() {
	a1 := A{s: "abc"}
	a2 := A{s: "abc"}
	as := make(AS, 2)
	as2 := make(AS, 2)

	as = append(as, &a1)
	as = append(as, &a2)

	b1 := A{s: "abc"}
	b2 := A{s: "abc"}

	as1 := make(AS, 2)
	as1 = append(as1, &b1)
	as1 = append(as1, &b2)

	as2 = append(as2, &a1)
	as2 = append(as2, &a2)

	println(reflect.DeepEqual(as, as2))
	println(reflect.DeepEqual(as, as1))

}
