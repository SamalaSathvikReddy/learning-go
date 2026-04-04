package main

import "fmt"

func sendBulk(numMessages int) float64 {
	var totalCost float64 = 0.0
	for i := 0; i < numMessages; i++ {
		var j float64 = float64(i)
		var curCost float64 = 1 + 0.01*j
		totalCost = totalCost + curCost
	}
	return totalCost
}

func maxMessages(thresh int) int {
	msgs := 0
	cost := 0
	for i := 0; ; i++ {
		if cost+i+100 <= thresh {
			cost += (i + 100)
			msgs += 1
		} else {
			break
		}
	}
	return msgs
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%s\n", "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Printf("%s\n", "fizz")
		} else if i%5 == 0 {
			fmt.Printf("%s\n", "buzz")
		} else {
			fmt.Print(i)
		}
	}
}

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}


func countConnections(groupSize int) int {
	connections := 0
	for i := 1; i < groupSize; i++ {
		connections += (groupSize - i)
	}
	return connections
}