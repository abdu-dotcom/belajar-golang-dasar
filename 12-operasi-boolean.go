package main

import "fmt"

func main() {
	// true && true = true
	// true && false = false
	// false && true = false
	// false && false = false
	// ! true = false | !false = true (kebalikan dari nilai tersebut)

	var nilaAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaAkhir > 80 // false
	var lulusAbsensi bool = absensi > 80      // false

	var lulus bool = lulusNilaiAkhir && lulusAbsensi // false && false = false

	fmt.Println(lulus) // false
}
