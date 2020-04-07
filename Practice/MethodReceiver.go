package main

import "fmt"

type Data int

func (data Data) printData() {
	fmt.Println(data)
}

func (data *Data) scaleData() {
	*data = *data * 10
}

func main() {
	var data Data = 10
	data.printData()
	data.scaleData()
	data.printData()
}
