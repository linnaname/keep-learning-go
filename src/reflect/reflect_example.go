package main

import (
	"fmt"
	"reflect"
)

type Child struct {
	Name     string
	handsome bool
}

type MyInt int
type YourInt int

type dog struct {
	Name string
	Age  int
}

func (d dog) Run(speed string) {
	fmt.Println(d.Name, "is running", speed)
}

func main() {
	qcrao := Child{Name: "qcrao", handsome: true}
	fmt.Println(reflect.TypeOf(qcrao))
	v := reflect.ValueOf(&qcrao)
	f := v.Elem().FieldByName("Name")
	fmt.Println(f.String())

	f.SetString("stefno")
	fmt.Println(f.String())

	f = v.Elem().FieldByName("handsome")
	// 这一句会导致 panic，因为 handsome 字段未导出
	//f.SetBool(true)
	fmt.Println(f.Bool())

	//DeepEqual
	m := MyInt(1)
	y := YourInt(1)
	fmt.Println(reflect.DeepEqual(m, y)) // false

	//reflect call method
	d := dog{"linana", 1}
	val := reflect.ValueOf(d)
	mv := val.MethodByName("Run")
	args := []reflect.Value{reflect.ValueOf("fastly")}
	mv.Call(args)
}
