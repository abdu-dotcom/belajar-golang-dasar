package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	// operasi matematika
	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// Augmanted Assignment
	var d = 30
	d += 10 // d = d + 10
	fmt.Println(d)

	//Unary Operator
	var e = 40
	e-- // e = e + 1
	fmt.Println(e)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
