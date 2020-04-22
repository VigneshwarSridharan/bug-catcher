package main

import "fmt"

type Book struct {
	title  string
	author string
}

func main() {
	books := []Book{
		Book{author: "Jams", title: "Go Lang!"},
		Book{author: "Lusifer", title: "JAVA!"},
	}
	for _, value := range books {
		fmt.Printf("%+v \n", value)
	}
}

func calc(a int, b int) (int, int) {
	return a + b, a - b
}

func add(x int, y int) int {
	return x + y
}
