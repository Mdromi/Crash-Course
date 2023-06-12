package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

// 13.2
func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	sendIsOld(isOldChan, emails)

	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

// TEST SUITE -- Don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails []email) {
	go func ()  {
		for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
		}
	}()
}

type email struct {
	body string
	date time.Time
}

func test2(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==========================================")
}

// 13.3
func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func test3(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}

func main() {
	// 13.3
	test3(3)
	test3(4)
	test3(5)

	// 13.2
	test2([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test2([]email{
		{
			body: "Yo are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test2([]email{
		{
			body: "Today is the day!",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What do you want for lunch?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Why are you the way that you are?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test2([]email{
		{
			body: "Did we do it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Letsa Go!",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Okay...?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})

	// 13.1
	// test("Hello there Stacy!")
	// test("Hi there John!")
	// test("Hey there Jane!")
}
