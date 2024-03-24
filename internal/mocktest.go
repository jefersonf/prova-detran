package internal

import (
	"math/rand"
	"os"

	"github.com/jefersonf/prova-detran/internal/parser"
)

var (
	questionSetFile = "./internal/data/questions.json"

	testSet *parser.QuestionSet
)

type LabeledQuestion struct {
	Label int
	parser.Question
}

func NewMocktest(numQuestions int) ([]LabeledQuestion, error) {

	jsonFile, err := os.Open(questionSetFile)
	if err != nil {
		return []LabeledQuestion{}, err
	}
	defer jsonFile.Close()

	testSet, err = parser.ParseQuestionSet(jsonFile)
	if err != nil {
		return []LabeledQuestion{}, err
	}

	ids := make([]int, len(testSet.Questions))
	for i := range ids {
		ids[i] = i + 1
	}

	for i := range ids {
		j := rand.Intn(i + 1)
		ids[i], ids[j] = ids[j], ids[i]
	}

	qset := make([]LabeledQuestion, min(numQuestions, len(testSet.Questions)))
	for i := range qset {
		qID := ids[i] - 1
		qset[i].Label = i + 1
		qset[i].Topic = testSet.Questions[qID].Topic
		qset[i].Statement = testSet.Questions[qID].Statement
		qset[i].Alternatives = testSet.Questions[qID].Alternatives
		qset[i].RightAnswer = testSet.Questions[qID].RightAnswer
	}

	return qset, nil
}
