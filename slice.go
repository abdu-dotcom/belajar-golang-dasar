package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"February",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice = months[3:12]
	var slice2 = months[5:]
	slice3 := append(slice2, "Senini")

	slice3[6] = "Des"

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	//========================== || ======================

	// make slice in at the beginning
	//make([] tipe kembalian, lenghtSlice ,capasitasSlice)
	newSlice := make([]string, 2, 5)

	//newSlice = append(newSlice, "hello world,", "world,", "hello,")
	newSlice[0] = "abduh"
	newSlice[1] = "batubara"

	fmt.Println("make slice ", newSlice)
	//========================== || ======================
	// 	//make([] tipe kembalian, lenghtSlice(firstSlice) ,capasitasSlice(firstSlice))
	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copy slice", copySlice)

	// Hati hati saat membuat array
	// saat membuat array, kita harus berhati-hati, jika salah, maka yang kita
	//	buat bukanlah array, melainkan slice

	// perbedaan slice dengan array
	iniArray := [...]int{1, 2, 3, 4, 5} // mendefinisikan jumlah array
	iniSlice := []int{1, 2, 3, 4, 5}    // tidak mendifiniskan jumlah

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
