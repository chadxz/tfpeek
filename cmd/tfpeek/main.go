package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chadxz/tfpeek/internal/terraform"
)

func run() error {
	path := os.Args[1]
	fmt.Printf("Searching '%s' for Terraform modules...\n", path)

	modules, err := terraform.CollectModules(path)
	if err != nil {
		return err
	}

	fmt.Printf("Discovered %d valid Terraform modules.\n", len(modules))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
