package main

import "fmt"

func main() {
	// ketika membuat variable wajib nentukan tipe datanya
	// namu jika langsung menginisialisasikan value maka idak perlu menentukan tipe datanya
	var name string
	name = "Muhammad Abduh"

	// bisa mendeklrasai variable tanpa harus menggunakan var, dengan cara menggunakan :=
	age := 20

	// multiple variable
	var (
		firstCall = "Abduh "
		LastCall  = "Batubara"
	)
	// multiple variable another way
	firstName, lastName := "Muhammad Abduh ", "Husaini Batubara"

	fmt.Println(name)
	fmt.Println(firstCall + LastCall)
	fmt.Println(firstName + lastName)
	fmt.Println(age)
}
