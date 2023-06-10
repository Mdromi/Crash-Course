package main

import (
	"errors"
	"fmt"
	"sort"
)

// 10.1
func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

// don't touch below this line

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}
func test(messages []string, formatter func(string) string) {
	defer fmt.Println("====================================")
	formattedMessages := getFormattedMessages(messages, formatter)
	if len(formattedMessages) != len(messages) {
		fmt.Println("The number of messages returned is incorrect.")
		return
	}
	for i, message := range messages {
		formatted := formattedMessages[i]
		fmt.Printf(" * %s -> %s\n", message, formatted)
	}
}

// 10.4

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a, b string) {
		fmt.Println(formatter(a, b))
	}
}

// don't touch below this line

func test4(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

// 10.5
const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

// don't touch below this line

type user struct {
	name   string
	number int
	admin  bool
}

func test5(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

func main() {
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addSignature)
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addGreeting)

	// 10.4
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test4("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test4("Error on mail server", mailErrors, commaDelimit)

	// 10.5
	users := map[string]user{
		"john": {
			name:   "john",
			number: 18965554631,
			admin:  true,
		},
		"elon": {
			name:   "elon",
			number: 19875556452,
			admin:  true,
		},
		"breanna": {
			name:   "breanna",
			number: 98575554231,
			admin:  false,
		},
		"kade": {
			name:   "kade",
			number: 10765557221,
			admin:  false,
		},
	}

	fmt.Println("Initial users:")
	usersSorted := []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	test5(users, "john")
	test5(users, "santa")
	test5(users, "kade")

	fmt.Println("Final users:")
	usersSorted = []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")
}
