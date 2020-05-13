package main

import (
	"fmt"
)

type Interface interface {
	GetStart()
}

type Implements struct {
}

func (impl *Implements) GetStart() {
	fmt.Println("======")
}

func Test(intf Interface) {

	// _, ok := intf.(Implements)
	// if ok {
	// 	fmt.Println("++++++")
	// } else {
	// 	fmt.Println("------")
	// }
}

func main() {

	var intf Interface
	//intf = new(Implements)
	intf = &Implements{}
	Test(intf)

}
