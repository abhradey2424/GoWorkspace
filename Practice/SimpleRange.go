package main

import "fmt"

func simpleRange() {
	arr := []int{}
	index := 0
	for index < 5 {
		arr = append(arr, index)
		index++
	}
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

func main() {
	simpleRange()
}
