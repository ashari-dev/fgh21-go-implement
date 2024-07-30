package main

import "fmt"

func fazzFood(price int, voucher string, distance int, tax bool) (int, int, int, int) {
	var biayaTax int	=	0
	if tax {
	 biayaTax = 5 * price /100
	} 

	var biayaDistance int = 0
	if distance>2 {
		biayaDistance = (distance - 2) * 3000 + 5000
	}

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

	var total int = price + biayaTax - biayaVoucher + biayaDistance

	return biayaVoucher, biayaDistance, biayaTax ,total
}

func main() {
	const price int = 25000
	// const voucer string = "FAZZFOOD50"
	const voucer string = "DITRAKTIR60"
	const distance int = 5
	const tax bool = true
	
	var potongan,biayaAntar,pajak,total = fazzFood(price,voucer,distance,tax) 

	fmt.Printf("Harga : %d\n", price)
	fmt.Printf("Potongan : %d\n", potongan)
	fmt.Printf("Biaya Antar : %d\n", biayaAntar)
	fmt.Printf("Pajak : %d\n", pajak)
	fmt.Printf("Total : %d\n", total)
}