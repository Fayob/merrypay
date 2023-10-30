package coupon

import (
	"fmt"
	"math/rand"
)

func coupon() string {
	var coupon string
	alphab := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	num := []int{1, 2, 3, 4 ,5, 6, 7, 8, 9, 0}

	for i := 0; i < 6; i++ {
		rAlp, rNum := rand.Intn(26), rand.Intn(10)
		coupon += fmt.Sprint(num[rNum]) + alphab[rAlp]
	}
	return coupon
}

var db = map[string]string{}

func GenerateCoupon() string  {
	coupon := coupon()

	if val, ok := db[coupon]; ok {
		fmt.Println(val)
		GenerateCoupon()
	}
	db[coupon] = coupon
	fmt.Println(db)
	return coupon
}