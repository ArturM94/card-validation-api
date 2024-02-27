package validator_test

import (
	"main/validator"
	"testing"
)

func TestInvalidCardNumber(t *testing.T) {
	cases := []struct {
		Name string
		Card *validator.Card
	}{
		{
			Name: "card number 1111111111111",
			Card: &validator.Card{
				Number:          "1111111111111",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "card number 23123413123123",
			Card: &validator.Card{
				Number:          "23123413123123",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "card number 333357664814443",
			Card: &validator.Card{
				Number:          "333357664814443",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "card number 8858857776625461",
			Card: &validator.Card{
				Number:          "8858857776625461",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "card number 96685747578214568",
			Card: &validator.Card{
				Number:          "96685747578214568",
				ExpirationMonth: "10",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "card number 999999999999999999",
			Card: &validator.Card{
				Number:          "999999999999999999",
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

			if err != validator.ErrInvalidCardNumber {
				t.Errorf("got %q want %q", err, validator.ErrInvalidCardNumber)
			}
		})
	}
}

func TestCardExpirationMonth(t *testing.T) {
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

func TestExpiredCard(t *testing.T) {
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

func TestUnexpiredCard(t *testing.T) {
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

func TestValidCard(t *testing.T) {
	cases := []struct {
		Name string
		Card *validator.Card
	}{
		{
			Name: "Visa 1",
			Card: &validator.Card{
				Number:          "4716316840400268",
				ExpirationMonth: "04",
				ExpirationYear:  "2029",
			},
		},
		{
			Name: "Visa 2",
			Card: &validator.Card{
				Number:          "4532472853421881",
				ExpirationMonth: "06",
				ExpirationYear:  "2030",
			},
		},
		{
			Name: "Mastercard 1",
			Card: &validator.Card{
				Number:          "5220224772895685",
				ExpirationMonth: "08",
				ExpirationYear:  "2030",
			},
		},
		{
			Name: "Mastercard 2",
			Card: &validator.Card{
				Number:          "5523256208900428",
				ExpirationMonth: "11",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "Amex 1",
			Card: &validator.Card{
				Number:          "344217768864121",
				ExpirationMonth: "08",
				ExpirationYear:  "2031",
			},
		},
		{
			Name: "Amex 2",
			Card: &validator.Card{
				Number:          "342247912056730",
				ExpirationMonth: "07",
				ExpirationYear:  "2029",
			},
		},
		{
			Name: "Unionpay 1",
			Card: &validator.Card{
				Number:          "6226986923714765",
				ExpirationMonth: "03",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "Unionpay 2",
			Card: &validator.Card{
				Number:          "6214839445487077",
				ExpirationMonth: "06",
				ExpirationYear:  "2028",
			},
		},
		{
			Name: "Diners 1",
			Card: &validator.Card{
				Number:          "3033717418112862",
				ExpirationMonth: "05",
				ExpirationYear:  "2030",
			},
		},
		{
			Name: "Diners 2",
			Card: &validator.Card{
				Number:          "3004958606181054",
				ExpirationMonth: "1",
				ExpirationYear:  "2030",
			},
		},
		{
			Name: "Discover 1",
			Card: &validator.Card{
				Number:          "6011562629233753",
				ExpirationMonth: "01",
				ExpirationYear:  "2031",
			},
		},
		{
			Name: "Discover 2",
			Card: &validator.Card{
				Number:          "6011456903879722",
				ExpirationMonth: "1",
				ExpirationYear:  "2030",
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
