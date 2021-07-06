package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukan Nilai : ")
	fmt.Scan(&nilai)

	if nilai == 10 {
		fmt.Println("Sangat Bagus")
	} else if nilai < 10 && nilai >= 5 {
		fmt.Println("Bagus")
	} else if nilai < 5 && nilai > 1 {
		fmt.Println("Buruk")
	} else {
		fmt.Println("Coba Lagi")
	}
}
