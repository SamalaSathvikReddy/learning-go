package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i := 0; i < len(messages); i++ {
		messages[i].tags = tagger(messages[i])
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lowerCase := strings.ToLower(msg.content)
	if strings.Contains(lowerCase, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowerCase, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
