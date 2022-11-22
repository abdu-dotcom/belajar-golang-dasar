package main

import (
	"fmt"
)

func main() {
	// operasi perbandingan adalah operasi untuk membandingkan dua buah data
	// menghasilkan true dan false
	// operator { >, <, >=, <=, ==, !- }

	var name1 = "Eko"
	var name2 = "Ek"
	var result bool = name1 >= name2

	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
