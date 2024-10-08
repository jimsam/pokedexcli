package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	var lastResponse any
	var err error
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if len(input) == 0 {
			fmt.Println("You must enter a command!")
		}
		val, found := getCommands()[input[0]]
		if found {
			err = val.callback(&lastResponse, input)
			if err != nil && err.Error() == "exit" {
				fmt.Println(err)
				break
			} else if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("This command was not found!")
		}
		fmt.Printf("pokedex > ")
	}
}
