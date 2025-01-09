package main

import (
	"fmt"
	"strings"
)

func main() {

	// Soal 1

    for i := 1; i <= 20; i++ {
        if i%2 == 1 && i%3 == 0 {
            fmt.Println(i, "- I Love Coding")
        } else if i%2 == 0 {
            fmt.Println(i, "- Berkualitas")
        } else {
            fmt.Println(i, "- Santai")
        }
    }

	// Soal 2

	for i := 1; i <= 7; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println() 
    }

	// Soal 3

	// Mendefinisikan array
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	// Mengambil elemen mulai dari indeks 2 sampai akhir
	result := strings.Join(kalimat[2:], " ")

	// Menampilkan hasil
	fmt.Println("[" + result + "]")

	// Soal 4

	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}

    for i, sayur := range sayuran {
        fmt.Print(i+1)  // Mencetak angka
        fmt.Print(". ")
        fmt.Println(sayur)  // Mencetak nama sayuran
    }

	// Soal 5
	var satuan = map[string]int{
        "panjang": 7,
        "lebar":   4,
        "tinggi":  6,
    }

    // Looping untuk mencetak panjang, lebar, dan tinggi
    for key, value := range satuan {
        fmt.Print(key, " = ", value, "\n")
    }

    // Menghitung volume balok
    volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
    fmt.Print("volume balok = ", volume, "\n")

}
