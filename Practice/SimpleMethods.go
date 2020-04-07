package main

import "fmt"

type Area struct {
	length, width int
}

func (area Area) showArea() {
	fmt.Println(area.length, area.width)
}

func main() {
	area := Area{10, 20}
	area.showArea()
}
