package internal

import "fmt"

// QuestionSet ...
type QuestionSet struct {
	Questions []Question `json:"questoes"`
}

// Question ...
type Question struct {
	Topic        string        `json:"tema"`
	Statement    string        `json:"enunciado"`
	Alternatives []Alternative `json:"alternativas"`
	RightAnswer  string        `json:"correta"`
}

/ Alternative represents a prompt question option.
type Alternative struct {
	Label     string `json:"opcao"`
	Statement string `json:"alternativa"`
}

// PromptQuestion represents a prompt question statement.
type PromptQuestion struct {
	Text       string
	QuestionID string
}

// FormattedAlternatives returns a list of pre-formatted question statements.
func (q *Question) FormattedAlternatives(questionID int) []PromptQuestion {
	formatted := make([]PromptQuestion, len(q.Alternatives))
	for i, a := range q.Alternatives {
		formatted[i].Text = fmt.Sprintf("[%s] %s", a.Label, a.Statement)
		formatted[i].QuestionID = fmt.Sprintf("Questão %0d", questionID)
	}
	return formatted
}
