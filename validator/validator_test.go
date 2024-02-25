package validator_test

import (
	"main/validator"
	"testing"
)

func TestCardNumberValidation(t *testing.T) {
	cases := []struct {
		Name string
		Card *validator.Card
	}{
		{
			Name: "card number length is less than 16 digits",
			Card: &validator.Card{
				Number:          "1111111111111",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "card number length is greater than 16 digits",
			Card: &validator.Card{
				Number:          "111111111111111111",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			cv := validator.NewCardValidator()
			err := cv.Validate(test.Card)
			if err == nil {
				t.Fatal("expected error but didn't get one")
			}

			if err != validator.ErrInvalidCardNumberLength {
				t.Errorf("got %q want %q", err, validator.ErrInvalidCardNumberLength)
			}
		})
	}
}

func TestCardExpirationMonthValidation(t *testing.T) {
	cases := []struct {
		Name string
		Card *validator.Card
	}{
		{
			Name: "expiration month is 00",
			Card: &validator.Card{
				Number:          "4111111111111111",
				ExpirationMonth: "00",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "expiration month is greater than 12",
			Card: &validator.Card{
				Number:          "4111111111111111",
				ExpirationMonth: "13",
				ExpirationYear:  "2028",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			cv := validator.NewCardValidator()
			err := cv.Validate(test.Card)
			if err == nil {
				t.Fatal("expected error but didn't get one")
			}

			if err != validator.ErrInvalidMonth {
				t.Errorf("got %q want %q", err, validator.ErrInvalidMonth)
			}
		})
	}
}

func TestExpiredCardValidation(t *testing.T) {
	cases := []struct {
		Name string
		Card *validator.Card
	}{
		{
			Name: "expired card at 01 2021",
			Card: &validator.Card{
				Number:          "4111111111111111",
				ExpirationMonth: "01",
				ExpirationYear:  "2021",
			},
		},
		{
			Name: "expired card at 01 2024",
			Card: &validator.Card{
				Number:          "4111111111111111",
				ExpirationMonth: "01",
				ExpirationYear:  "2024",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			cv := validator.NewCardValidator()
			err := cv.Validate(test.Card)
			if err == nil {
				t.Fatal("expected error but didn't get one")
			}

			if err != validator.ErrCardExpired {
				t.Errorf("got %q want %q", err, validator.ErrCardExpired)
			}
		})
	}
}

func TestUnexpiredCardValidation(t *testing.T) {
	cases := []struct {
		Name string
		Card *validator.Card
	}{
		{
			Name: "card with expiration date 01 2028",
			Card: &validator.Card{
				Number:          "4111111111111111",
				ExpirationMonth: "01",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "expiration date is a last month of a year",
			Card: &validator.Card{
				Number:          "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2028",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			cv := validator.NewCardValidator()
			err := cv.Validate(test.Card)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
