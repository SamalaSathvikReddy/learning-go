package main

import (
	"fmt"
	"time"
)

// learning structs
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	Membership
	name   string
	number int
}

type sender struct {
	user
	rateLimit int
}

type authenticationInfo struct {
	username string
	password string
}

type contact struct {
	sendingLimit int32
	age          int32
	userId       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

type Membership struct {
	Type             string
	messageCharLimit int
}

func (s authenticationInfo) getBasicAuth() string {
	var msg string = fmt.Sprintf("Authorization : Basic %s:%s", s.username, s.password)
	return msg
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.recipient.name == "" || mToSend.sender.name == "" || mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func (u user) SendMessage(message string, messageLength int) (bool, string) {
	if messageLength <= u.messageCharLimit {
		return true, message
	}
	return false, ""
}

func newUser(name string, membershipType string, number int) user {
	if membershipType == "premium" {
		return user{
			Membership{
				"premium",
				1000,
			},
			name,
			number,
		}
	}
	return user{
		Membership{
			"basic",
			100,
		},
		name,
		number,
	}
}

// Interfaces
func sendMessage(msg message) (string, int) {
	return msg.getMessage(), 3 * len(msg.getMessage())
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}


