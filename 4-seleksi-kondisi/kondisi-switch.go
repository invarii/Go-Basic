package main

import "fmt"

func main() {
	var nilai string

	fmt.Print("Masukan Nilai : ")
	fmt.Scan(&nilai)

	switch nilai {
	case "a":
		fmt.Println("Sangat Bagus")
	case "b":
		fmt.Println("Bagus")
	case "c":
		fmt.Println("Buruk")
	default:
		fmt.Println("Coba Lagi")
	}
}
