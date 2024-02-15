package repository

import (
	"context"
	"fmt"
	"math/rand"
	service "merrypay/service/model"
	"merrypay/types"
	"sync"
)

var count int

func (m *Model) Jackpot(ctx context.Context, guess [5]int, username string) ([5]int, error) {
	if username == "" {
		return [5]int{}, fmt.Errorf("username can't be empty")
	}

	if len(guess) != 5 {
		return [5]int{}, fmt.Errorf("kindly predict 5 numbers")
	}

	var wg sync.WaitGroup
	var user types.User
	var earning types.Earning
	var errs []error

	wg.Add(1)
	go func() {
		defer wg.Done()
		searchedUser, err := m.Model.FindUser(ctx, username)
		errs = append(errs, err)
		user = searchedUser
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		userEarning, err := m.Model.GetEarning(ctx, username)
		errs = append(errs, err)
		earning = userEarning
	}()
	wg.Wait()

	// fmt.Printf("user %+v\n", user)
	// fmt.Printf("earning %+v\n", earning)
	// fmt.Printf("errors %+v\n", errs)
	for _, err := range errs {
		if err != nil {
			fmt.Printf("error %+v", errs)
			return [5]int{}, err
		}
	}

	withdrawArg := types.WithdrawalParam{
		Kind: "media",
		Amount: 1,
		WithdrawBy: username,
	}

	withdrawal, err := m.WithdrawFund(ctx, withdrawArg)
	if err != nil {
		// fmt.Println("withdraw fund")
		return [5]int{}, err
	}

	completeWithdrawalArg := types.CompleteWithdrawalParams{
		ID: withdrawal.ID,
		Amount: withdrawal.Amount,
		Kind: withdrawal.Kind,
	}

	err = m.CompleteWithdrawal(ctx, completeWithdrawalArg.ID)
	if err != nil {
		fmt.Println("complete withdrawal")
		return [5]int{}, err
	}

	count++
	fmt.Println(count)

	if !user.WonJackpot && earning.Referrals < 2 && count%2 == 0 {
		// logic here
		arg := service.CreateTransaction{
			Kind:       "referral",
			Amount:     12,
			TransactBy: user.Username,
		}
		err = m.transactHelperFunc(ctx, arg, earning)
		if err != nil {
			return [5]int{}, err
		}

		user.WonJackpot = true
		updateUserArg := types.UpdateUserParams{
			Username: user.Username,
			Email: user.Email,
			WonJackpot: user.WonJackpot,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Membership: user.Membership,
		}
		_, err = m.Model.UpdateUser(ctx, updateUserArg)
		if err != nil {
			return [5]int{}, err
		}

		return guess, nil
	}

	return numbersPredictor(guess), nil
}

func numbersPredictor(guess [5]int) [5]int {
	num1, num2, num3, num4, num5 := rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)
	if num1 == guess[0] && num2 == guess[1] && num3 == guess[2] && num4 == guess[3] && num5 == guess[4] {
		numbersPredictor(guess)
	}
	return [5]int{num1, num2, num3, num4, num5}
}
