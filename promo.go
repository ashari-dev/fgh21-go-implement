package main

import (
	"fazztrack/fazzfood/calc"
	"fmt"
)

func fazzFood(price int, voucher string, distance int, tax bool) {
	discound:= calc.Discound(price,voucher)
	biayaTax:= calc.Taxed(price,tax)
	deliferyFee:= calc.DeliferyFee(distance)

	var total int = price + biayaTax - discound + deliferyFee

	fmt.Printf("Harga 			: %d\n", price)
	fmt.Printf("Potongan 		: %d\n", discound)
	fmt.Printf("Biaya Antar 		: %d\n", deliferyFee)
	fmt.Printf("Pajak 			: %d\n", biayaTax)
	fmt.Printf("Total 			: %d\n", total)
}

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
	
	fazzFood(price,voucer,distance,tax) 

}