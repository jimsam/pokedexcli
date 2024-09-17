package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	var lastResponse any
	var err error
	for scanner.Scan() {
		input := scanner.Text()
		val, found := getCommands()[input]
		if found {
			err = val.callback(&lastResponse)
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
