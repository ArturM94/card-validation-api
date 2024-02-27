package validator

import (
	"errors"
	"strconv"
	"time"
)

const (
	ErrCodeInvalidCardNumber = 001
	ErrCodeInvalidMonth      = 002
	ErrCodeCardExpired       = 003
	ErrCodeInternal          = 999
)

var (
	ErrInvalidCardNumber = errors.New("invalid card number")
	ErrInvalidMonth      = errors.New("expiration month must be double digit numerical string from 01 to 12")
	ErrCardExpired       = errors.New("card expired")
)

type Card struct {
	Number          string `json:"number"`
	ExpirationMonth string `json:"expirationMonth"`
	ExpirationYear  string `json:"expirationYear"`
}

type CardValidator struct {
}

func (v *CardValidator) validateNumber(number string) error {
	var sum int
	parity := len(number) % 2

	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return ErrInvalidCardNumber
		}

		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	if sum%10 != 0 {
		return ErrInvalidCardNumber
	}

	return nil
}

func (v *CardValidator) validateMonth(month string) (time.Month, error) {
	m, err := strconv.ParseInt(month, 10, 64)
	if err != nil {
		return 0, err
	}
	if m < 1 || m > 12 {
		return 0, ErrInvalidMonth
	}
	return time.Month(m), nil
}

func (v *CardValidator) validateYear(year string) (int, error) {
	y, err := strconv.ParseInt(year, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(y), nil
}

func (v *CardValidator) Validate(c *Card) error {
	err := v.validateNumber(c.Number)
	if err != nil {
		return err
	}

	month, err := v.validateMonth(c.ExpirationMonth)
	if err != nil {
		return err
	}

	year, err := v.validateYear(c.ExpirationYear)
	if err != nil {
		return err
	}

	if month == time.December {
		year++
		month = time.January
	} else {
		month++
	}

	now := time.Now()
	firstDayNextMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())

	if now.After(firstDayNextMonth) {
		return ErrCardExpired
	}

	return nil
}

func NewCardValidator() *CardValidator {
	return &CardValidator{}
}
