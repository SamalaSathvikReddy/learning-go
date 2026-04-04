package main

import "errors"

const planFree = "free"
const planPro = "pro"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	message := [3]string{primary, secondary, tertiary}
	cost := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}
	return message, cost
}

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case "pro":
		return messages[:], nil
	case "free":
		return messages[:2], nil
	}
	return nil, errors.New("unsupported plan")
}

func getMessageCosts(messages []string) []float64 {
	cost := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost[i] = 0.01 * float64((len(messages[i])))
	}
	return cost
}

func sums(nums ...int) int {
	finalSum := 0
	for i := 0; i < len(nums); i++ {
		finalSum += nums[i]
	}
	return finalSum
}
