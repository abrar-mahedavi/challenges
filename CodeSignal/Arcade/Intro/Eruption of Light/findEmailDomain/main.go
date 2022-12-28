package main

import (
	"fmt"
	"strings"
)

func findEmailDomain(address string) string {
	emailAddress := strings.Split(address, "@")
	return emailAddress[len(emailAddress)-1]
}

func main() {
	fmt.Printf("%v", findEmailDomain("prettyandsimple@example.com"))
}
