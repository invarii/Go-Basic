package main

import "fmt"

func main() {
	var buah []string
	buah = append(buah, "apel")
	buah = append(buah, "semangka")
	buah = append(buah, "jeruk")
	buah = append(buah, "mangga")

	fmt.Println("Sebelum")
	fmt.Println("Jumlah Data :", len(buah))
	fmt.Println("Data Array :", buah)

	buah = append(buah, "durian")
	buah = append(buah, "jambu")

	fmt.Println("sesudah")
	for i := 0; i < len(buah); i++ {
		fmt.Println("buah [", i, "]:", buah[i])
	}
}
