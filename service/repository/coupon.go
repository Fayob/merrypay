package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
)

func (m *Model) GenerateCoupon(ctx context.Context, username string) (string, error)  {
	user, err := m.Model.FindUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("unknown user")
		}
		return "", err
	}

	if user.Membership != "admin" {
		return "", fmt.Errorf("unauthorized route")
	}

	coupon := coupon()

	c, err := m.Model.GetCoupon(ctx, coupon)

	if err != sql.ErrNoRows {
		return "", err
	}

	if c.Digit == coupon {
		m.GenerateCoupon(ctx, username)
	}

	_, err = m.Model.SaveCoupon(ctx, coupon, username)
	if err != nil {
		return "", err
	}

	return coupon, nil
}

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
