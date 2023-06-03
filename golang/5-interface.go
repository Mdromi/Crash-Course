package main

import (
	"fmt"
	"time"
)

// 1
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * r.width * r.height
}

// 2
func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendReport struct {
	reportName string
	numberOfSends int
}

func (sr sendReport) getMessage() string {
	return fmt.Sprintf("Your %s, report is ready, You've sent %v", sr.reportName, sr.numberOfSends)
}

// 3
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// 4
func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

// Test Code
func test(m message) {
	sendMessage(m)
	fmt.Println("==================================")
}

func test3(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func test4(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	test(sendReport{reportName: "First Report", numberOfSends: 10})	
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})

	// 3
	test3(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test3(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test3(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})

	// 4
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test4(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test4(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test4(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test4(e, e)
}