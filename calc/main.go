package calc

func Discound(price int, voucher string) int {
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

func Taxed(price int, tax bool) int {
	var biayaTax int	=	0
	if tax {
	 biayaTax = 5 * price /100
	} 
	return biayaTax
}
func DeliferyFee(distance int) int {
	var biayaDistance int = 0
	if distance>2 {
		biayaDistance = (distance - 2) * 3000 + 5000
	}
	return biayaDistance
}