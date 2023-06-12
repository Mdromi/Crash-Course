package main

import (
	"fmt"
	"time"
)

// 13.4
func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	fmt.Println(emailsToSend)
	return emailsToSend
}

// TEST SUITE - Don't Touch Below This Line

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

// 13.5
func countReports(numSentCh chan int) int {
	total := 0

	for {
		numSent, ok := <- numSentCh
		if !ok {
			break
		}
		total += numSent
	}

	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func test2(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

// 13.6
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
func concurrrentFib(n int) {
	chInts := make(chan int)
	go func ()  {
		fibonacci(n, chInts)
	}()
	for v := range chInts {
		fmt.Println(v)
	}
}
func test3(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}


func main() {
	// 13.6
	test3(10)
	test3(5)
	test3(20)
	test3(13)
	// 13.5
	test2(3)
	test2(4)
	test2(5)
	test2(6)

	// 13.4
	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}
