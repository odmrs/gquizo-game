package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type quiz struct {
	Question string
	Answer   string
	Line     string
}

func mapCsvData(data [][]string) (map[string]string, int) {
	questionsAndAnswers := make(map[string]string)
	var linesOfQuestions int
	for _, line := range data {
		var rec quiz
		for j, field := range line {
			if j == 0 {
				rec.Question = field
			} else if j == 1 {
				rec.Answer = field
			}
		}
		linesOfQuestions++
		questionsAndAnswers[rec.Question] = rec.Answer
	}
	return questionsAndAnswers, linesOfQuestions
}

func quiz_game(data [][]string) {
	var counter int = 1
	var answer string
	var right int
	questionsAnswers, line := mapCsvData(data)

	for k, v := range questionsAnswers {
		fmt.Printf("Problem #%v: %s = ", counter, k)
		fmt.Scanln(&answer)

		counter++
		if answer == v {
			right += 1
			continue
		}
	}

	fmt.Printf("You scored %v of counter %v\n", right, line)
}

func main() {
	// Read CSV file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Convert to MAP
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	quiz_game(data)
}
