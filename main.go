package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/atyagi9006/qnaSolution/model"
	"github.com/atyagi9006/qnaSolution/process"
)

func main() {
	//intput questions
	quefile, err := ioutil.ReadFile("inputQue.json")
	if err != nil {
		fmt.Print(err)
	}

	var questions model.Questions
	err = json.Unmarshal([]byte(quefile), &questions)
	if err != nil {
		fmt.Print(err)
	}

	quebank, err := process.QuestionsToMap(questions.Questions)
	if err != nil {
		fmt.Print(err)
	}
	//input Answer by user
	//intput questions
	ansfile, err := ioutil.ReadFile("inputAnsByUser.json")
	if err != nil {
		fmt.Print(err)
	}

	var queAttempts model.QueAttempts
	err = json.Unmarshal([]byte(ansfile), &queAttempts)
	if err != nil {
		fmt.Print(err)
	}

	result, err := process.EvaluateResult(quebank, queAttempts.Attempts)
	if err != nil {
		fmt.Print(err)
	}
	process.PrintResult(result)
}
