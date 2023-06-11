package main

import "fmt"

func main() {
	var smsSendLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendLimit, costPerSMS, hasPermission, username)

	empty := ""
	numCar := 10
	temperature := 0.0
	isFunny := true
	fmt.Printf("%q %v %f %v\n", empty, numCar, temperature, isFunny)
	fmt.Printf("The type of isFunny: %T\n", isFunny)

	milage, company := 8080, "Tesla"
	fmt.Println(milage, company)

	accountAge := 2.4
	accountAgeInt := int(accountAge)
	fmt.Println("Your account existed for", accountAgeInt, "years")

	const pi = 3.1416

	hight := 3

	if hight > 6 {
		fmt.Println("You are supper tall")
	} else if hight > 4 {
		fmt.Println("You are supper enough")
	} else {
		fmt.Println("You are not tall enough")
	}

	fmt.Println(9)
}