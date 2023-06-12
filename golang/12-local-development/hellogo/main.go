// go build && ./hellogo

package main

import (
	"fmt"

	"github.com/mdromi/mystrings"
)

func main() {
	fmt.Println(mystrings.Reverse("hello world"))
}