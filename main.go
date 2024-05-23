package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, question := range questions {
		fmt.Printf("Question %s: Answer %s\n", question[0], question[1])
	}
}
