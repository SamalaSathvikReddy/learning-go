package main

import (
	"errors"
	"fmt"
)

type divideError struct {
	dividendValue float64
}

// func (D divideError) Error() string {
// 	return fmt.Sprintf("cannot divide %v by zero", D.dividendValue)
// }

// func divide(dividend float64, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		return 0, divideError{dividendValue: dividend}
// 	}
// 	return dividend / divisor, nil
// }

func divide(x, y float64) (float64, error) {
	if y == 0 {
		var err error = errors.New("Cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}

func validateStatus(status string) error {
	if(status == ""){
		var err error = errors.New("status cannot be empty");
		return  err
	} 
	if(len(status) > 140) {
		var err error = errors.New("status exceeds 140 characters");
		return err
	}
	return nil
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costCustomer, errCustomer := sendSMS(msgToCustomer)
	if errCustomer != nil {
		return 0, errCustomer
	}
	costSpouse, errSpouse := sendSMS(msgToSpouse)
	if errSpouse != nil {
		return 0, errSpouse
	}
	return costCustomer + costSpouse, nil

}
func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that $%.2f to be sent to %s cannnot be sent", cost, recipient)
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
