package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//NG Challenge 3 : Looping 1
	people := []map[string]any{
		{"name": "Hank", "Age": 50, "Job": "Polisi"},
		{"name": "Heisenberg", "Age": 52, "Job": "Ilmuwan"},
		{"name": "Skyler", "Age": 48, "Job": "Akuntan"},
	}
	for _, person := range people {
		name, _ := person["name"].(string)
		age, _ := person["Age"].(int)
		job, _ := person["Job"].(string)

		fmt.Printf("Hi, Perkenalkan, Nama saya %s, umur saya %d,dan saya bekerja sebagai %s\n", name, age, job)

	}

	//NG Challenge 3 : Looping 2
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	sum1 := 0.0
	for _, value := range slice1 {
		sum1 += value
	}
	avg := float64(sum1) / float64(len(slice1))

	sum2 := 0.0
	for _, value := range slice2 {
		sum2 += value
	}
	avg1 := float64(sum2) / float64(len(slice2))

	sort.Float64s(slice1)
	var median1 float64
	if len(slice1)%2 == 0 {
		median1 = (slice1[len(slice1)/2-1] + slice1[len(slice1)/2]) / 2
	} else {
		median1 = slice1[len(slice1)/2]
	}

	sort.Float64s(slice2)
	var median2 float64
	if len(slice2)%2 == 0 {
		median2 = (slice2[len(slice2)/2-1] + slice2[len(slice2)/2]) / 2
	} else {
		median2 = slice2[len(slice2)/2]
	}

	fmt.Printf("Slice 1: %v\n", slice1)
	fmt.Printf("Rata-rata: %.2f\n", avg)
	fmt.Printf("Jumlah Slice 1: %.2f\n", sum1)
	fmt.Printf("Median Slice 1: %2f\n", median1)

	fmt.Printf("Slice 1: %v\n", slice2)
	fmt.Printf("Rata-rata: %.2f\n", avg1)
	fmt.Printf("Jumlah Slice 2: %.2f\n", sum2)
	fmt.Printf("Median Slice 1: %2f\n", median2)

	//NG Challenge 3 : Logic 1 - Palindrome

	var kata1 string

	fmt.Printf("Masukkan kata : ")
	fmt.Scanf("%s", &kata1)

	kataRune := []rune(kata1) // merubah type data "kata" menjadi rune
	reverseKata := []rune{}   // menampung nilai reverse kata

	// Proses reverse kata
	for i := len(kataRune) - 1; i >= 0; i-- {
		reverseKata = append(reverseKata, kataRune[i])
	}

	// Check apakah kata sama dengan reverse kata
	if kata1 == string(reverseKata) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//NG Challenge 3 : Logic 2 - XOXO
	var kata string
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&kata)

	countX := strings.Count(kata, "x")
	countO := strings.Count(kata, "o")

	fmt.Printf("Jumlah 'x': %d\n", countX)
	fmt.Printf("Jumlah 'o': %d\n", countO)

	if countX == countO {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	//NG Challenge 3 : Logic 3
	angka := []int{3, 2, 7, 4, 5, 10, 1, 8, 9, 6}
	fmt.Println("Angka sebelum :", angka)

	var temp int
	var swapped bool

	for i := 0; i < len(angka)-1; i++ {
		swapped = false
		for j := 0; j < len(angka)-i-1; j++ {
			if angka[j] > angka[j+1] {
				temp = angka[j]
				angka[j] = angka[j+1]
				angka[j+1] = temp
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}
	fmt.Println("Angka setelah di sorting :", angka)

	//NG Challenge 3 : Logic 4 - Asterisk Level 1
	rows1 := 5
	for i := 0; i < rows1; i++ {
		fmt.Println("* ")
	}
	fmt.Println()

	//NG Challenge 3 : Logic 5 - Asterisk Level 2
	rows2 := 5
	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	//NG Challenge 3 : Logic 6 - Asterisk Level 3
	rows3 := 5
	for i := 0; i < rows3; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

	//NG Challenge 3 : Logic 7 - Asterisk Level 4
	rows4 := 5
	for i := 0; i < rows4; i++ {
		for j := i; j < rows4; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
