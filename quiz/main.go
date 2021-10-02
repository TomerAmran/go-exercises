package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	answer string 
}


func main(){
	testFilePath := flag.String("path", "quiz.csv","the path to the test")
	flag.Parse()
	// fmt.Println(*testFilePath)
	timeLimit := flag.Int("time", 3, "time limit to test in seconds")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0
	commandline_reader := bufio.NewScanner(os.Stdin)
	file, err := os.Open(*testFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
quizeloop:
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) < 2 {
			log.Fatal("line shold have at least one comma")	
		}
		question:= Question {
			question: strings.Join(record[:len(record)-1],","),
			answer :record[len(record)-1],
		}
		answer := make(chan string)
		fmt.Print(question.question, " ")
		go func(){
			commandline_reader.Scan()
			answer <- commandline_reader.Text()
		}()
		select{
		case <-timer.C:
			fmt.Println()
			fmt.Println("time is up")
			break quizeloop
		case user_answer:=<-answer:
			if user_answer == question.answer {
				score +=1
			}
		}
	}
	fmt.Println("your score is: ", score)
	





}