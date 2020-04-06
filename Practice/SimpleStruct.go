package main

import "fmt"

type Size struct {
	width  int
	height int
}

func setSize() Size {
	size := Size{2, 3}
	newSize := &size
	newSize.width = 10
	return size
}

func main() {
	fmt.Println(setSize())
}
