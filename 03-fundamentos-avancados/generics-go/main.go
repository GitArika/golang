// Package main -> Generics em Go
package main

import (
	"fmt"
	"generics-go/examples"
)

func main() {
	foo("string")
	foo(20)
	foo([]int{1})

	var mt MyType = "meu tipo customizado"
	myGeneric(mt)
	myGeneric2("Hello, world")
	myGeneric2(100)
	myGeneric2(mt)

	var ms = MyStruct[string]{"Generic Struct"}
	fmt.Println(ms)

	var slice = []int{1,2,3,4,5}
	fmt.Println(examples.Contains(slice, 4))
}

// tipo é definido durante a compilação
func foo[T any](arg T) {
	fmt.Println(arg)
}
// tipo é definido durante o runtime
func bar(arg any) {
	fmt.Println(arg)
}

type MyConstraint interface {
	Foo()
}

type MyConstraint2 interface {
	~int | ~string
}

type MyType string
func (MyType) Foo() {}

func myGeneric[T MyConstraint](arg T) {
	fmt.Println(arg)
}

func myGeneric2[T MyConstraint2](arg T) {
	fmt.Println(arg)
}

type MyStruct[T any] struct{
	Foo T
}
