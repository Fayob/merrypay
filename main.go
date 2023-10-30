package main

import (
	"fmt"
	"merrypay/api/coupon"
)

func main()  {
	output := coupon.GenerateCoupon()
	fmt.Println(output)
}