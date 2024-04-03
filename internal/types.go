package internal

import "fmt"

type QuestionSet struct {
	Questions []Question `json:"questoes"`
}

type Question struct {
	Topic        string        `json:"tema"`
	Statement    string        `json:"enunciado"`
	Alternatives []Alternative `json:"alternativas"`
	RightAnswer  string        `json:"correta"`
}

type Alternative struct {
	Label     string `json:"opcao"`
	Statement string `json:"alternativa"`
}

type PromptQuestion struct {
	Text       string
	QuestionID string
}

func (q *Question) FormattedAlternatives(questionID int) []PromptQuestion {
	formatted := make([]PromptQuestion, len(q.Alternatives))
	for i, a := range q.Alternatives {
		formatted[i].Text = fmt.Sprintf("[%s] %s", a.Label, a.Statement)
		formatted[i].QuestionID = fmt.Sprintf("Questão %0d", questionID)
	}
	return formatted
}
