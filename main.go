package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Read CSV file
	_, err := os.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Convert to MAP

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
