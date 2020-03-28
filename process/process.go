package process

import (
	"errors"
	"fmt"

	"github.com/atyagi9006/qnaSolution/model"
)

var (
	errInvalidInputQues             = errors.New("Invalid questions")
	errInvalidQuestionBank          = errors.New("Invalid questionBank")
	errInvalidAnsByUser             = errors.New("Invalid Answers By User")
	errInvalidQuestionAttemptByUser = errors.New("Invalid Question attempted By User")
)

// QuestionsToMap convert Questions
func QuestionsToMap(inputQuestions []*model.Question) (map[int]*model.Question, error) {
	if len(inputQuestions) <= 0 {
		return nil, errInvalidInputQues
	}

	qnaMap := make(map[int]*model.Question)
	for _, question := range inputQuestions {
		qnaMap[question.ID] = question
	}
	
	return qnaMap, nil
}

// EvaluateResult result for user
func EvaluateResult(QnABank map[int]*model.Question, AnsByUser []*model.QueAttempt) (*model.Result, error) {
	if len(QnABank) <= 0 {
		return nil, errInvalidQuestionBank
	}

	if len(AnsByUser) <= 0 {
		return nil, errInvalidAnsByUser
	}
	resultMap := make(map[string]int)
	totalScore := 0
	for _, ans := range AnsByUser {
		if que, ok := QnABank[ans.QuestionID]; ok {
			if len(que.Answers) == len(ans.Answers) {
				if validateAns(que.Answers, ans.Answers) {
					totalScore++
					if val, ok := resultMap[que.Topic]; ok {
						resultMap[que.Topic] = val + 1
					} else {
						resultMap[que.Topic] = 1
					}
				}
			} else {
				return nil, errInvalidQuestionAttemptByUser
			}

		}
	}
	result := model.Result{
		TopicWiseScore: resultMap,
		TotalScore:     totalScore,
	}
	return &result, nil
}

func validateAns(actual []int, attempt []int) bool {
	for i, v := range actual {
		if v != attempt[i] {
			return false
		}
	}
	return true
}

// PrintResult print result for user
func PrintResult(res *model.Result) {
	fmt.Println("Total Score :", res.TotalScore)
	fmt.Println("Score topic Wise: ", res.TopicWiseScore)
}
