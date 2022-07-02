package main

import (
	"fmt"
	"reflect"
)

type Usb interface {
	start()
	stop()
}

type Phone struct {
	camera
	phoneNum int
}

type camera struct {
	Name   string
	pinpai string
}

// 正则表达式
func main() {

	var ss int = 1
	s1 := &ss
	cc := *s1
	fmt.Println(cc)

	p1 := Phone{camera: camera{Name: "zhangg"}, phoneNum: 123123}
	// typep := reflect.TypeOf(p1)
	valuep := reflect.ValueOf(p1)

	fmt.Println(valuep.Field(1))

	fmt.Println(valuep.Kind())
}
