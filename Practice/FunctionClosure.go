package main

import "fmt"

func closure() func() {
	index := 0
	return func() {
		fmt.Print(index, " ")
		index++
	}
}

func main() {
	inc := closure()
	count := 0
	for count < 10 {
		inc()
		count++
	}
}
