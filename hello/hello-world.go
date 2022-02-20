package main

import "fmt"

func main() {
	fmt.Println("welcome Gophers 😁 ")

	var x float32
	var y float32

	x = 12
	y = 13

	fmt.Printf("the value of x = %v of type %T \n", x, x)
	fmt.Printf("the value of y = %v of type %T \n", y, y)

	var mean = (x + y) / 2

	fmt.Printf("Mean of x and y = %v", mean)
}
