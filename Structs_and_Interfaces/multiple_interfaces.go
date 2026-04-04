package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	a, isEmail := e.(email)
	if isEmail {
		return a.toAddress, e.cost()
	}
	c, isSMS := e.(sms)
	if isSMS {
		return c.toPhoneNumber, e.cost()
	}
	return "", 0
}

func getExpenseReportUsingSwitch(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	}
	return "", 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type formatter interface {
	format()
}

type plainText struct {
	message string
}

type bold struct {
	message string
}

type code struct {
	message string
}

func (text plainText) format() string {
	return text.message
}

func (text bold) format() string {
	return fmt.Sprintf("**%s**", text.message)
}

func (text code) format() string {
	return fmt.Sprintf("``%s``", text.message)
}
