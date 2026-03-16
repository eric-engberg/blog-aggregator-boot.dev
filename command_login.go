package main

import (
	"fmt"
)

func commandLogin(name string) error {
	fmt.Printf("Logging in as %s...\n", name)
	return nil
}
