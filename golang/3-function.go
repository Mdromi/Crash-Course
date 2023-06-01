package main

import "fmt"

func sub(x, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func increment(x int) int {
	x++
	return x
}

func getNames() (string, string) {
	return "Md", "Romi"
}

func getPoints() (x int, y int) {
	return 3, 4
}

func main() {
	fmt.Println(concat("Romi, ", "Happy Golang"))
	fmt.Println(sub(20, 10))

	x := 5
	x = increment(x)
	fmt.Println(x)
	fmt.Println(getNames())

	y, _:= getPoints()
	fmt.Println(y)
}

