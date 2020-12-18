package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//contains
	fmt.Println("==========================")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	//????? true
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	//index
	fmt.Println("==========================")
	fmt.Println(strings.Contains("team", "t"))
	fmt.Printf("%d\n", strings.IndexFunc("studygolang", func(c rune) bool {
		if c > 'u' {
			return true
		}
		return false
	}))

	//count
	fmt.Println("==========================")
	fmt.Println(strings.Count("teamt", "t"))
	//????当 sep 为空时，Count 的返回值是：utf8.RuneCountInString(s) + 1
	fmt.Println(strings.Count("teamt", ""))

	//upper/lower
	fmt.Println("==========================")
	fmt.Println(strings.ToUpper("teamt"))
	fmt.Println(strings.ToLower("TEamt"))

	//compare
	fmt.Println("==========================")
	fmt.Println(strings.Compare("teamt", "teea"))
	fmt.Println(strings.EqualFold("teamt", "teamT"))
	fmt.Println(strings.EqualFold("teamt", "tamT"))

	//trim
	fmt.Println("==========================")
	var s = "aaasddfgaaaa"
	fun := func(c rune) bool {
		if c != 'a' {
			return false
		}
		return true
	}
	fmt.Println(strings.TrimFunc(s, fun)) //输出   sddfg

	//split
	fmt.Println("==========================")
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))

	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))

	//join
	fmt.Println("==========================")
	sl := []string{"aaa", "b"}
	fmt.Println(strings.Join(sl, ","))

	s1 := "Go编程"

	fmt.Println(len(s1))
	fmt.Println(len(string(rune('编'))))
	fmt.Println(len([]rune(s1)))
}

func StringPlus(p []string) string {
	var s string
	l := len(p)
	for i := 0; i < l; i++ {
		s += p[i]
	}
	return s
}

func StringFmt(p []interface{}) string {
	return fmt.Sprint(p...)
}

func StringJoin(p []string) string {
	return strings.Join(p, "")
}

func StringBuffer(p []string) string {
	var b bytes.Buffer
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func StringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

func StringBuilderGrowFirst(p []string, cap int) string {
	var b strings.Builder
	l := len(p)
	b.Grow(cap)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
