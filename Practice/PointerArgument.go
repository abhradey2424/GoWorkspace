package main

import "fmt"

type Data int

func passByPointer(data *int) {
	fmt.Println("Pass by Pointer: ", *data)
}

func passByValue(data int) {
	fmt.Println("Pass by Value: ", data)
}

func (data Data) valueReceiver() {
	fmt.Println("Value Receiver Method: ", data)
}

func (data *Data) pointerReceiver() {
	fmt.Println("Pointer Receiver Method: ", *data)
	*data = *data * 100
}

func main() {
	var data int = 10
	passByPointer(&data)
	passByValue(data)
	var data1 Data = 20
	data1.valueReceiver()
	data1.pointerReceiver()
	data1.valueReceiver()
}
