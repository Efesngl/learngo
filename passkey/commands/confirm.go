package commands

import "fmt"

func confirm(msg string) bool {
	var input string
	fmt.Printf("%s (y/n): ", msg)
	fmt.Scan(&input)
	return input == "y" || input == "Y"
}
