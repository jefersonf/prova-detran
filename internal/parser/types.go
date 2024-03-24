package parser

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

func (q *Question) FormattedAlternatives() []string {
	formatted := make([]string, len(q.Alternatives))
	for i, a := range q.Alternatives {
		formatted[i] = fmt.Sprintf("[%s] %s", a.Label, a.Statement)
	}
	return formatted
}
