package main

import (
	"fmt"
	"time"
)

func main() {
	do(1)
	do(2)
	do(4)

	fmt.Println(isWeekend(time.Now()))
}

func do(x int) {
	switch x {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
			fallthrough
		case 3:
			fmt.Println(3)
		default:
			fmt.Println("outra coisa")
	}
}

func isWeekend(x time.Time) bool {
	switch x.Weekday() {
		case time.Sunday, time.Saturday:
			return true
		default:
			return false
	}
}
