package model

type Question struct {
	ID        int      `json:"id"`
	Topic     string   `json:"topic"`
	Statement string   `json:"statement"`
	Answers   []int    `json:"answers"`
	Options   []string `json:"options"`
}

type Questions struct {
	Questions []*Question `json:"questions"`
}

type QueAttempt struct {
	QuestionID int   `json:"question_id"`
	Answers    []int `json:"answers"`
}

type QueAttempts struct {
	Attempts []*QueAttempt `json:"attempts"`
}

type Result struct {
	TopicWiseScore map[string]int
	TotalScore     int
}
