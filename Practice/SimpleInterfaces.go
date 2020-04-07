package main

import "fmt"

type Display interface {
	display()
	displayNew()
}

type Data int
type DataF float64

func (data Data) display() {
	fmt.Println("Display Method: ", data)
}

func (data Data) displayNew() {
	fmt.Println("DisplayNew Method: ", data)
}

func (dataF DataF) display() {
	fmt.Println("Display Float Method: ", dataF)
}

func (dataF DataF) displayNew() {
	fmt.Println("DisplayNew Float Method: ", dataF)
}

func main() {
	var displayData Display
	data := (Data)(10)
	displayData = data
	displayData.display()
	displayData.displayNew()
	dataF := (DataF)(20.33)
	displayData = dataF
	displayData.display()
	displayData.displayNew()
}
