package main

import "fmt"

func main() {
	fmt.Println("entered the main function")
	sonu()
	iter()
}
func sonu() {
	fmt.Println("entered into sonu finction ")
}
func iter() {
	for i := 0; i < 10; i++ {
		fmt.Println("sonu")
	}
}
