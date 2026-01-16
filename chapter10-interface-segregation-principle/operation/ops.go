package operation

import "fmt"

type U1Ops interface {
	Ops1()
}

type U2Ops interface {
	Ops2()
}

type U3Ops interface {
	Ops3()
}

type Operation struct{}

func (o Operation) Ops1() {
	fmt.Println("Operation Ops1")
}

func (o Operation) Ops2() {
	fmt.Println("Operation Ops2")
}

func (o Operation) Ops3() {
	fmt.Println("Operation Ops3")
}
