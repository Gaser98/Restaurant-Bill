package main

import (
	"fmt"
	
)

func main() {
	mybill := newBill("Mark's bill")
	fmt.Println(mybill.format())

	
}