package main

import "fmt"

func main() {
	// constant adalah varible yang nilai tidak bisa diubah lagi setelah diberi nilai
	const firstNmae, lastName string = "muhammad", "abduh"

	fmt.Println(firstNmae + " " + lastName)
}
