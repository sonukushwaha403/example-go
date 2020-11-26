package main

import "fmt"

func main() {
	p, _ := fmt.Println("entered the main function")
	sonu()
	iter()
	fmt.Println(p)
	//fmt.Println(e)
}
func sonu() {
	fmt.Println("entered into sonu finction ")
}
func iter() {
	for i := 0; i < 10; i++ {
		fmt.Println("sonu")
	}
}
