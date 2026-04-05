package main

import (
	"errors"
	"strings"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}

	res := make(map[string]user)
	for i := 0; i < len(names); i++ {
		res[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}
	return res, nil
}

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	val, ok := users[name]
	if !ok {
		return false, errors.New("User Not Found")
	}
	if val.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for i := 0; i < len(messagedUsers); i++ {
		_, ok := validUsers[messagedUsers[i]]
		if ok {
			validUsers[messagedUsers[i]] += 1
		}
	}
}

func getNameCounts(names []string) map[rune]map[string]int {
	res := make(map[rune]map[string]int)
	for i := 0; i < len(names); i++ {
		res[(rune)(names[i][0])][names[0]] += 1
	}
	return res
}

func countDistinctWords(messages []string) int {
	res := make(map[string]struct{})
	for i := 0; i < len(messages); i++ {
		v := strings.Fields(messages[i])
		for _, strs := range v {
			if _, ok := res[strings.ToLower(strs)]; !ok {
				res[strings.ToLower(strs)] = struct{}{}
			}
		}
	}
	return len(res)
}

type user struct {
	name        string
	phoneNumber int
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
