package fazz

import (
	"fazztrack/fazzfood/calc"
	"fmt"
)

func FazzFood(price int, voucher string, distance int, tax bool) {
	discound := calc.Discound(price, voucher)
	biayaTax := calc.Taxed(price, tax)
	deliferyFee := calc.DeliferyFee(distance)

	var total int = price + biayaTax - discound + deliferyFee

	fmt.Printf("Harga 			: %d\n", price)
	fmt.Printf("Potongan 		: %d\n", discound)
	fmt.Printf("Biaya Antar 		: %d\n", deliferyFee)
	fmt.Printf("Pajak 			: %d\n", biayaTax)
	fmt.Printf("Total 			: %d\n", total)
}