package main

import (
	"fmt"
	"strings"
)


func cleanInput(text string) []string {
	// return []string{}
	clean := strings.Trim(text, " ")
	clean = strings.Trim(clean, "\t")
	return strings.Fields(clean)
}


func main() {
	fmt.Println("Hello, World!")
}

