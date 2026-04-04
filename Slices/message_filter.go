package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	var fi []Message
	for i := 0; i < len(messages); i++ {
		if messages[i].Type() == filterType {
			fi = append(fi, messages[i])
		}
	}
	return fi
}

func isValidPassword(password string) bool {
	dig := false
	len := (len(password) > 5 && len(password) <= 12)
	cas := false
	for _, ch := range password {
		if ch >= 'A' && ch <= 'Z' {
			cas = true
		}
		if ch >= '0' && ch <= '9' {
			dig = true
		}
	}
	return (dig && len && cas)
}
