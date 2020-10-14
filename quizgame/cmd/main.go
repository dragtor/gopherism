package main

// The program should output the total number of questions correct and how many questions there were in total.
// Questions given invalid answers are considered incorrect.

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type timer struct {
	timeLimit int
	event     chan string
}


func NewTimer(timeLimit int) *timer {
	return &timer{
		timeLimit: timeLimit,
        event : make(chan string,0),
	}
}

func (t *timer) stopwatch() {
	time.Sleep(time.Second * time.Duration(t.timeLimit))
	t.event <- "Done"
}

type Quiz struct {
	questionCount       int
	correctAnswersCount int
	enableTimer         bool
	timer               *timer
}

func New(enableTimer bool, timer *timer) *Quiz {
	return &Quiz{
		questionCount:       0,
		correctAnswersCount: 0,
		enableTimer:         enableTimer,
		timer:               timer,
	}
}

func (q *Quiz) IncrementQuestionCount() {
	q.questionCount++
}

func (q *Quiz) IncrementCorrectAnsCount() {
	q.correctAnswersCount++
}

func evaluateExpression(input string) (int, error) {
	f := func(c rune) bool {
		return c == '+'
	}
	result := strings.FieldsFunc(input, f)
	op1, _ := strconv.ParseInt(result[0], 10, 64)
	op2, _ := strconv.ParseInt(result[1], 10, 64)
	return int(op1 + op2), nil
}

var (
	path        *string
	enableTimer *bool
	timeLimit   *int
)

func init() {
	path = flag.String("p", "../samples/problems.csv", "path to csv file")
	enableTimer = flag.Bool("et", false, "Flag to enable time")
	timeLimit = flag.Int("ts", 30, "Flag to set quiz time limit")
	flag.Parse()
	validateFlag()
}

func validateFlag() {
	if strings.TrimSpace(*path) == "" {
		panic("Invalid file path")
	}
}

func validateCSV(record []string) bool {
	if len(record) != 2 {
		return false
	}
	return true
}

func (quiz *Quiz) performQuiz() {
	f, err := os.Open(*path)
	check(err)
	csvReader := csv.NewReader(f)
		for {
			// Read csv line by line
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			check(err)
			// validate csv record
			isvalid := validateCSV(record)
			if !isvalid {
				log.Println("Err: invalid record")
				continue
			}
			quesStr := record[0]
			expectedAnsInt, err := strconv.ParseInt(record[1], 10, 64)
			if err != nil {
				log.Println("Err: invalid expected result")
				continue
			}
			quiz.IncrementQuestionCount()
			fmt.Printf("%s : ", quesStr)
			// take input from user
			ioreader := bufio.NewReader(os.Stdin)
			data, err := ioreader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			trimmedInput := strings.TrimSpace(data)
			if trimmedInput == "" {
				continue
			}
			userInput, err := strconv.Atoi(trimmedInput)
			if err != nil {
				log.Println("Err: invalid user input ")
				continue
			}
			// equate user input to exepected answer
			if userInput == int(expectedAnsInt) {
				//log.Println("incrementting correct count")
				quiz.IncrementCorrectAnsCount()
			}
		}

}

func main() {

	timer := NewTimer(*timeLimit)
	quiz := New(*enableTimer, timer)
	if *enableTimer {
		for {
			fmt.Printf("would you like to start quiz [y/n]: ")
			reader := bufio.NewReader(os.Stdin)
			userResponse, err := reader.ReadString('\n')
			if err != nil {
                panic(err)
				continue
			}
			if strings.ToLower(strings.TrimSpace(userResponse)) == "y"{
				go timer.stopwatch()
				break
			}
		}
	}
    go quiz.performQuiz()
	<-quiz.timer.event
	fmt.Printf("\ntotal number of questions : %d \ntotal invalid answers : %d \n", quiz.questionCount, quiz.questionCount-quiz.correctAnswersCount)
}
