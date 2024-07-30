package main

import "fmt"

func discound(price int, voucher string) int {
	var biayaVoucher int = 0
	if voucher == "FAZZFOOD50" ||voucher == "DITRAKTIR60" {
		if voucher == "FAZZFOOD50"  {
			if price >= 50000 {
				biayaVoucher = price *50/100
				if biayaVoucher > 50000 {
					biayaVoucher = 50000
				}
			}
		}
		if voucher == "DITRAKTIR60"  {
			if price >= 25000 {
				biayaVoucher = price *60/100
				if biayaVoucher > 30000 {
					biayaVoucher = 30000
				}
			}
		}
	}
	return biayaVoucher
}

func taxed(price int, tax bool) int {
	var biayaTax int	=	0
	if tax {
	 biayaTax = 5 * price /100
	} 
	return biayaTax
}
func jarak(distance int) int {
	var biayaDistance int = 0
	if distance>2 {
		biayaDistance = (distance - 2) * 3000 + 5000
	}
	return biayaDistance
}

func fazzFood(price int, voucher string, distance int, tax bool) {
	biayaVoucher:=discound(price,voucher)
	biayaTax:=taxed(price,tax)
	biayaDistance:= jarak(distance)

	var total int = price + biayaTax - biayaVoucher + biayaDistance

	fmt.Printf("Harga 			: %d\n", price)
	fmt.Printf("Potongan 		: %d\n", biayaVoucher)
	fmt.Printf("Biaya Antar 		: %d\n", biayaDistance)
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