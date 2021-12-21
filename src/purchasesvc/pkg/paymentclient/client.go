package paymentclient

import (
	"errors"
	"math/rand"
	"time"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Pay(cardNumber, cvc, exp string) error {

	expTime, err := time.Parse("01/06", exp)
	if err != nil {
		return errors.New("wrong expiration date")
	}

	if time.Now().After(expTime.Add(time.Hour * 24 * 31)) { // add month
		return errors.New("card expired")
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	if num < 20 {
		return errors.New("not enough money") //20% chance of not enough money
	}

	return nil
}
