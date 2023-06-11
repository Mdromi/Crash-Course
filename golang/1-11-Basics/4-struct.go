package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message string
}

type car struct {
	Make string
	Model string
	Height int
	Width int 
	FrontWheel Wheel
	BackWheel Wheel
}
type Wheel struct {
	Radius int
	Material string
}

type messageToSends struct {
	message string
	sender user
	recipient user
}
type user struct {
	name string
	number int
}

// Embedded Struct
type sender struct {
	user
	rateLimit int
}

// myCars := struct {
// 	Make string
// 	Model string
// } {
// 	Make: "Tesla",
// 	Model: "Model 2"
// }

// Anonymous Struct
type cars struct {
	Make string
	Model string
	Height int
	Width int 
	Wheel struct {
		Radius int
		Material string
	}
}

func test(m messageToSend) {
	fmt.Printf("Sending Message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("=======================================")
}

func canSendMessage(mToSend messageToSends) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}
func tests(m messageToSends) {
	if canSendMessage(m) {
		fmt.Printf("{\n Message: '%s'\n Sender: %s,\n Receiver: %s,\n Sender Number: %v,\n Receiver Number: %v\n}	\n", m.message, m.sender.name, m.recipient.name, m.sender.number, m.recipient.number)
		fmt.Println("=======================================")
	} else {
		fmt.Println("...Can't send message")
	}
}

func test3(s sender) {
	fmt.Println("Sender Name:", s.name)
	fmt.Println("Sender Number:", s.number)
	fmt.Println("RateLimit:", s.rateLimit)
}

// Struct Method
type rect struct {
	Width int
	Height int
}
func(r rect) area() int{
	return r.Width * r.Height
}

type authenticationInfo struct {
	username string
	password string
}

func ( authI authenticationInfo) getBasicsAuth() string {
	return fmt.Sprintf("Authentication: Basics %s:%s", authI.username, authI.password)
}

func test4(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicsAuth())
	fmt.Println("======================================")
}

func main() {
	test(messageToSend{ phoneNumber: 123434234, message: "Thanks for sign up"})
	test(messageToSend{ phoneNumber: 693902313, message: "Love to have you aboard"})
	test(messageToSend{ phoneNumber: 34989289, message: "We're so exited to have you"})

	myCar := car{}
	myCar.FrontWheel.Radius = 5

	// Create user instances
	senders := user{
		name:   "John",
		number: 1234567890,
	}

	recipient := user{
		name:   "Jane",
		number: 9876543210,
	}
	fmt.Println(senders, recipient)

	tests(messageToSends{message: "Hello, how are you?", sender: senders, recipient: recipient})

	test3(sender{user: senders, rateLimit: 12})

	r := rect{
		Width: 5,
		Height: 10,
	}
	fmt.Println(r.area())

	test4(authenticationInfo{username: "mdromi", password: "aweiroaw"})
}