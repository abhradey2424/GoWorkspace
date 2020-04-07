package main

import "fmt"

func simpleMap() {
	hash := make(map[int]string)
	hash[0] = "abc"
	hash[1] = "def"

	for index := range hash {
		fmt.Println(hash[index])
	}
}

func main() {
	simpleMap()
}
