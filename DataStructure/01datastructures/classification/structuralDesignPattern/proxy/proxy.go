package main

import "fmt"

// IRealObject interface
type IRealObject interface {
	performAction()
}

// RealObject struct
type RealObject struct{}

// RealObject class method performAction
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

// VirtualProxy struct
type VirtualProxy struct {
	realObject *RealObject
}

// VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}

	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}

func main() {
	var object VirtualProxy
	object.performAction()
}
