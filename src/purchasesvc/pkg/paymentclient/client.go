package paymentclient

import (
	"errors"
	"math/rand"
	"time"
)

var (
	ErrWrongEXP       = errors.New("wrong expiration date")
	ErrCardExpired    = errors.New("card expired")
	ErrNotEnoughMoney = errors.New("not enough money")
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Pay(cardNumber, cvc, exp string) error {

	expTime, err := time.Parse("0106", exp)
	if err != nil {
		return ErrWrongEXP
	}

	if time.Now().After(expTime.Add(time.Hour * 24 * 31)) { // add month
		return ErrCardExpired
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	if num < 20 {
		return ErrNotEnoughMoney //20% chance of not enough money
	}

	return nil
}
