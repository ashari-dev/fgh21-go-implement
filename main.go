package main

import (
	"fazztrack/fazzfood/fazz"
	"fmt"
)

func main() {
	var price int 
	fmt.Print("Masukkan Harga : ")
	fmt.Scanln(&price)

	var voucer string
	fmt.Print("Masukkan Voucer : ")
	fmt.Scanln(&voucer)
	
	var distance int
	fmt.Print("Jarak antar : ")
	fmt.Scanln(&distance)
	const tax bool = true
	
	fazz.FazzFood(price,voucer,distance,tax) 

}