package main

import (
	"fmt"
	"sync"
)

type MyClass struct {
	attrib string
}

func (c *MyClass) SetAttrib(val string) {
	c.attrib = val
}

func (c *MyClass) GetAttrib() string {
	return c.attrib
}

var (
	once     sync.Once
	instance *MyClass
)

func GetMyClass() *MyClass {
	// once.Do() ใช้สำหรับการดำเนินการเพียงครั้งเดียวในโปรแกรม
	once.Do(func() {
		instance = &MyClass{"first"}
	})
	return instance
}

func main() {
	a := GetMyClass()
	a.SetAttrib("second")
	fmt.Println(a.GetAttrib())
	b := GetMyClass()
	b.SetAttrib("third")
	fmt.Println(b.GetAttrib())
}
