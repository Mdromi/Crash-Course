package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", &m.Recipient)
	fmt.Printf("Message: %v\n", &m.Text)
}

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}

// 11.2
func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}
func test2(messages []string) {
	for _, message := range messages {
		if message == "" {
			removeProfanity(nil)
			fmt.Println("nil message detected")
		} else {
			removeProfanity(&message)
			fmt.Println(message)
		}
	}
}

// 11.7
func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}
func (e email) print() {
	fmt.Println(e.message)
}
type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test7(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}


func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")

	// 11.2
	messages := []string{
		"well shoot, this is awful",
		"",
		"dang robots",
		"dang them to heck",
		"",
	}

	messages2 := []string{
		"well shoot",
		"",
		"Allan is going straight to heck",
		"dang... that's a tough break",
		"",
	}

	test2(messages)
	test2(messages2)

	// 11.7
	test7(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test7(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")
}
