package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
) 

func main() {

	// soal 1 
	var1 := "Bootcamp"
	var2 := "Digital"
	var3 := "Skill"
	var4 := "Sanbercode"
	var5 := "Golang"
	fmt.Println(var1,var2, var3, var4, var5)

	// soal 2

	var halo string = "Halo Dunia"
	var ganti string = "Golang"
	hasil := strings.Replace(halo, "Dunia", ganti, 1)
	fmt.Println(hasil)

	// 	soal 3

	var kataPertama = "saya"
	var kataKedua = "senang"
	var  kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.ReplaceAll(kataKedua, "s", "S")   
	kataKetiga = strings.ReplaceAll(kataKetiga, "r", "R") 
	kataKeempat = strings.ToUpper(kataKeempat)           

	hasil = strings.Join([]string{kataPertama, kataKedua, kataKetiga, kataKeempat}, " ")

	fmt.Println(hasil)

	// soal 4 

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angka1, _ := strconv.Atoi(angkaPertama)
	angka2, _ := strconv.Atoi(angkaKedua)
	angka3, _ := strconv.Atoi(angkaKetiga)
	angka4, _ := strconv.Atoi(angkaKeempat)

	result  := angka1 + angka2 + angka3 + angka4

	fmt.Println(result) 
	fmt.Println("Tipe Data result =>",reflect.TypeOf(result))   

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi",2 )
	fmt.Println("Tipe Data kalimat =>",reflect.TypeOf(kalimat))  
	fmt.Println(kalimat, "-", angka)
	
	
	

}