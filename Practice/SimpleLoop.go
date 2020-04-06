package main

import "fmt"

func loops() {
	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i = i + 1
	}
}

func main() {
	loops()
}
