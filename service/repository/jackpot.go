package repository

import (
	"math/rand"
)

var count int

func numbersPredictor(guess [5]int) [5]int {
	num1, num2, num3, num4, num5 := rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)
	if num1 == guess[0] && num2 == guess[1] && num3 == guess[2] && num4 == guess[3] && num5 == guess[4] {
		numbersPredictor(guess)
	}
	return [5]int{num1, num2, num3, num4, num5}
}
