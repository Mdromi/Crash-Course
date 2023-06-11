package main

import (
	"errors"
	"fmt"
)

const (
	retry1 = "click here to sign up"
	retry2 = "pretty please click here"
	retry3 = "we beg you to sign up"
)

func getMessageWithRetries() [3]string {
	return [3]string {
		retry1,
		retry2,
		retry3,
	}
}

// 8.2
const (
	planFree = "free"
	planPro  = "pro"
)
func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

// 8.6
func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++{
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

// 8.10
func sum(nums ...float64) float64 {
	sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum
}

// 8.11
type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costByDay := []float64{}
	return costByDay
	for i := 0; i < len(costs); i++{
		cost := costs[i]
		for cost.day >= len(costByDay){
			costByDay = append(costByDay, 0.0)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}

// 8.12
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// 8.16
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

// Test Code
func testSend(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("complete failure")
		}
	}
}

func test2(name string, doneAt int, plan string) {
	defer fmt.Println("=====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

func test3(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func test4(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func test5(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func test6(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func test7(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}


func main() {
	testSend("Bob", 0)
	testSend("Alice", 1)
	testSend("Mangalam", 2)
	testSend("Ozgur", 3)

	// 8.2
	test2("Ozgur", 3, planFree)
	test2("Jeff", 3, planPro)
	test2("Sally", 2, planPro)
	test2("Sally", 3, "no plan")

	// 8.6
	test3([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test3([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test3([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})

	// 8.10
	test4(1.0, 2.0, 3.0)
	test4(1.0, 2.0, 3.0, 4.0, 5.0)
	test4(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test4(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)

	// 8.11
	test5([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test5([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})

	// 8.12
	test6(3, 3)
	test6(5, 5)
	test6(10, 10)
	test6(15, 15)

	// 8.16
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test7(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test7(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test7(message, badWords)
}