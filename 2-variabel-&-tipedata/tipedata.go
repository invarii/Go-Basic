package main

import "fmt"

func main() {
	var angka uint8 = 123
	var min int = -123
	var desimal float32 = 1.28
	var boolean bool = true

	fmt.Printf("angka = %d %d %f \n", angka, min, desimal)
	fmt.Printf("angka = %.1f \n", desimal)
	fmt.Printf("bolean = %t", boolean)
}
