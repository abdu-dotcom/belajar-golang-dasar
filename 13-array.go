package main

import "fmt"

func main() {
	// Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
	// perlu menentukan jumlah data yang akan ditampung, tidak dapat ditambah setelah dibuat

	var names [3]string // nama variable -> jumlah array -> tipe data array
	names[0] = "Eko "
	names[1] = "Kurniawan "
	names[2] = "Khannedy"

	fmt.Println(names[0] + names[1] + names[2])

	// membuat array secara langsung
	var values = [3]int{90, 80, 95}
	values[0] = 10

	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(values[0])
	// beberapa function untuk memanipulasi array
	// len(array) = nilai panjang array
	// array[index] = nilai di posisi index
	// array[index] =mengubah data di posisi index
}
