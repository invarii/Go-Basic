package main

import "fmt"

func main() {
	const konstanta = 100
	var operator = ((2 + 4) * 2) / 4
	var uji = (operator == 3)
	var uji2 = (operator < 3)
	var uji3 = (operator != 3)

	fmt.Println("nilai tinggi =", konstanta)
	fmt.Println("operator =", operator)
	fmt.Printf("hasil jumlah = %d (%t) \n", operator, uji)
	fmt.Printf("hasil jumlah = %d (%t) \n", operator, uji2)
	fmt.Printf("hasil jumlah = %d (%t) \n", operator, uji3)
}
