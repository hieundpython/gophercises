package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	file, err := os.Open("problem.csv")
	if err != nil {
		panic("Can not open the files")
	}

	cvsReader := csv.NewReader(file)
	allRecord, err := cvsReader.ReadAll()
	if err != nil {
		panic("Can read csv file")
	}

	var allQuiz []Quiz

	for _, s := range allRecord {
		allQuiz = append(allQuiz, Quiz{question: s[0], answer: s[1]})
	}

	totalCorrect := 0

	for _, s := range allQuiz {
		fmt.Printf("Question %v, You Answer: 10", s.question)
		readInput := bufio.NewReader(os.Stdin)
		answer, err := readInput.ReadString('\n')

		if err != nil {
			panic("Can read from keyboard")
		}

		if strings.Trim(answer, "\r\n") == s.answer {
			totalCorrect++
		}
	}

	fmt.Printf("Total correct: %d\n", totalCorrect)
}
