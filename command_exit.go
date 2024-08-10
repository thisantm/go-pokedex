package main

import (
	"fmt"
	"os"
)

func commandExit(apiState *config) error {
	fmt.Println("bye bye")
	os.Exit(0)
	return nil
}
