package main

import "fmt"

func main() {
	var name string

	fmt.Println("Write your name:")
	fmt.Scanf("%s", &name)
	fmt.Printf("name - %s", name)
}
