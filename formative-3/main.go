package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1

	// Variabel input dalam string
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// Mengubah string ke integer
	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	// Variabel hasil perhitungan
	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	// Perhitungan
	luasPersegiPanjang = panjang * lebar
	kelilingPersegiPanjang = 2 * (panjang + lebar)
	luasSegitiga = (alas * tinggi) / 2

	// Menampilkan hasil
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)


	// Soal 2

	// Variabel nilai
	var nilaiJohn = 80
	var nilaiDoe = 50

	// Menentukan indeks nilai untuk John
	var indeksJohn string
	if nilaiJohn >= 80 {
		indeksJohn = "A"
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		indeksJohn = "B"
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		indeksJohn = "C"
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		indeksJohn = "D"
	} else {
		indeksJohn = "E"
	}

	// Menentukan indeks nilai untuk Doe
	var indeksDoe string
	if nilaiDoe >= 80 {
		indeksDoe = "A"
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		indeksDoe = "B"
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		indeksDoe = "C"
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		indeksDoe = "D"
	} else {
		indeksDoe = "E"
	}

	// Menampilkan hasil
	fmt.Println("Indeks nilai John:", indeksJohn)
	fmt.Println("Indeks nilai Doe:", indeksDoe)

	// Soal 3

	// Variabel tanggal lahir
	var tanggal = 12
	var bulan = 5
	var tahun = 1997
	var namaBulan string

	// Switch case untuk menentukan nama bulan
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	// Menampilkan output
	fmt.Println(tanggal, namaBulan, tahun)


	// soal 4

	// Variabel tahun kelahiran
	var tahunLahir = 1997
	var generasi string

	// Conditional untuk menentukan istilah generasi
	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		generasi = "Baby Boomer"
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		generasi = "Generasi X"
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		generasi = "Generasi Y (Millenials)"
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Tidak termasuk dalam istilah generasi yang disebutkan"
	}

	// Menampilkan output
	fmt.Println("Anda termasuk dalam:", generasi)

}
