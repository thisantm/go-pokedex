package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	cliCommandMap := cliStart()
	for {
		print("pokedex > ")
		scanner.Scan()
		input = scanner.Text()
		cliCommandMap[input].callback()
	}
}
