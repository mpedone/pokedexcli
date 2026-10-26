package main

import (
	"errors"
	"fmt"
)

func catch(links *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("What are you trying to catch, a cold?")
	}

	fmt.Printf("Throwing a ball at %s...\n", args[0])
	return nil
}
