package main

import (
	"fmt"
)

func main() {
y :=[]int{5,6,5,43,2,4,65,7,8677,44}
	fmt.Println(y)//printing the slice
	//now sliceing the slice
	fmt.Println(y[3:6])
	fmt.Println(y[2:])
	fmt.Println(y[:8])
	//now appendig to slice
	x:=[]int{6,4,3,5,5,6}
	x=append(x,3,4,5,56)
	fmt.Println(x)
	x=append(y,x...)
	fmt.Println(x)
	//now deleting from slice
	x=append(x[:13],x[15:]...)
	fmt.Println(x)
}

