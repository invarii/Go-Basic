package main

import "fmt"

func main() {
	var kiri = true
	var kanan = false

	kirinkanan := kiri && kanan
	kiriorkanan := kiri || kanan
	notkiri := !kiri

	fmt.Printf("kiri and kanan = %t \n", kirinkanan)
	fmt.Printf("kiri or kanan = %t \n", kiriorkanan)
	fmt.Printf("not kiri = %t \n", notkiri)
}
