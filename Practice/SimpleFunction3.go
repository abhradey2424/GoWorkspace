package main

import "fmt"

func swapNew(s1, s2 string) (s4, s3 string) {
	s3 = s1
	s4 = s2
	return
}

func main() {
	s1 := "Hello"
	s2 := "World"
	fmt.Println(s1, s2)
	s1, s2 = swapNew(s1, s2)
	fmt.Println(s1, s2)
	fmt.Printf("%T, %v", s1, s1)
}
