package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (m directMessage) importance() int {
	if m.isUrgent {
		return 50
	}
	return m.priorityLevel
}

func (m groupMessage) importance() int {
	return m.priorityLevel
}

func (m systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.priorityLevel
	case groupMessage:
		return v.groupName, v.priorityLevel
	case systemAlert:
		return v.alertCode, 100
	}
	return "", 0
}
