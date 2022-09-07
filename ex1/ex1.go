package ex1

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Quiz struct {
	question string
	answer   string
}

func Exec() {
	filePath := flag.String("filePath", "./ex1/problem.csv", "filePath of csv file")
	timeout := flag.Int("timeout", 5, "time for one quiz question")
	flag.Parse()

	file, err := os.Open(*filePath)
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

	fmt.Println("Game Start!!")

loopQuiz:
	for _, s := range allQuiz {
		fmt.Printf("Question %v, You Answer: ", s.question)

		timer := time.NewTimer(time.Duration(*timeout) * time.Second)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case answer := <-answerCh:
			if answer == s.answer {
				totalCorrect++
				continue
			}
		case <-timer.C:
			fmt.Println("\nTime Out!!")
			break loopQuiz
		}

	}

	fmt.Printf("\nCorrect %d in total %d\n", totalCorrect, len(allQuiz))
}
