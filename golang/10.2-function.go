package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

// 10.9
func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(msg string) int {
			return len(msg) * 2
		}, message)
	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func test9(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func main() {
	test([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	test([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	test([]emailBill{
		{743},
		{13},
		{8},
	})

	// 10.9
	test9([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	test9([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}
