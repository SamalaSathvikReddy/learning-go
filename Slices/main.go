package main

import "errors"

const planFree = "free"
const planPro = "pro"

type cost struct {
	day   int
	value float64
}

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

func getDayCosts(costs []cost, day int) []float64 {
	var finalCost []float64
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			finalCost = append(finalCost, costs[i].value)
		}
	}
	return finalCost
}

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for idx, str := range msg {
		chk := false
		for _, v := range badWords {
			if v == str {
				chk = true
				break
			}
		}
		if chk {
			return idx
		}
	}
	return -1
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		curRow := make([]int, cols)
		for j := 0; j < cols; j++ {
			curRow[j] = i * j
		}
		matrix[i] = curRow
	}
	return matrix
}
