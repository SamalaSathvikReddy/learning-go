package main

import (
	"errors"
	"fmt"
	"strings"
)

// import "strings"

type Message struct {
	Recipient string
	Success   bool
}

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cus *customer, trs transaction) error {
	if trs.transactionType == transactionDeposit {
		cus.balance += trs.amount
		return nil
	} else if trs.transactionType == transactionWithdrawal {
		if cus.balance-trs.amount >= 0 {
			cus.balance -= trs.amount
			return nil
		}
		return errors.New("Insufficient balance")
	}
	return errors.New("No such transaction mode")
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`To: %v, Message: %v`, m.Recipient, m.Success)
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	msg := *message
	n := len(msg)
	result := strings.Repeat("*", n)
	*message = result
}

func analyzeMessage(analytics *Analytics, message Message) {
	if message.Success {
		(*analytics).MessagesFailed++
	} else {
		(*analytics).MessagesSucceeded++
	}
	(*analytics).MessagesTotal++

}

func main() {
	msg := Message{Recipient: "Sathvik", Success: true}
	res := getMessageText(msg)
	fmt.Printf("%s\n", res)
}
