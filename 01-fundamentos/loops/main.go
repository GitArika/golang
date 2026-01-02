package main

import "fmt"

func main() {
	loopEstruturado()
	loopResumido()
	loopModerno()
	loopNovaVersao()
}

func loopEstruturado() {
	var res int
	for i:= 0; i < 10; i++ {
		res++
	}

	fmt.Println(res)
}

func loopResumido() {
	var res int
	var i int 
	for i < 10 {
		res++
		i++
	}
	fmt.Println(res)
}

func loopModerno() {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	var res int
	for range arr {
		res++
	}
	fmt.Println(res)
}

func loopNovaVersao() {
	var res int
	for range 10 {
		res++
	}
	fmt.Println(res)
}
