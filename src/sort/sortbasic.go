package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type personSlice []Person

func (p personSlice) Len() int           { return len(p) }
func (p personSlice) Less(i, j int) bool { return p[i].Age > p[j].Age }
func (p personSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println(a)

	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)

	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)

	sl := personSlice{{Name: "test", Age: 10}, {Name: "A", Age: 12}, {Name: "A", Age: 9}}
	sort.Stable(sl)
	fmt.Println(sl)
}
