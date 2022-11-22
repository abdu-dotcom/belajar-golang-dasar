package main

import "fmt"

func main() {
	// Type declarations adalah kemampuan membuat ulang tipe data baru dari
	// tipe data yang sudah ada

	type NoKTP string

	var noKtpEko NoKTP = "13123124123231"
	fmt.Println(noKtpEko)
}
