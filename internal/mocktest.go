package internal

import (
	"math/rand"
	"os"
)

var (
	questionSetFile = "./internal/data/questions.json"

	testSet *QuestionSet
)

type LabeledQuestion struct {
	Label int
	Question
}

type Mocktest []LabeledQuestion

func NewMocktest(numberOfQuestions int) (Mocktest, error) {

	jsonFile, err := os.Open(questionSetFile)
	if err != nil {
		return Mocktest{}, err
	}
	defer jsonFile.Close()

	testSet, err = ParseQuestionSet(jsonFile)
	if err != nil {
		return Mocktest{}, err
	}

	ids := make([]int, len(testSet.Questions))
	for i := range ids {
		ids[i] = i + 1
	}

	for i := range ids {
		j := rand.Intn(i + 1)
		ids[i], ids[j] = ids[j], ids[i]
	}

	questionSet := make(Mocktest, min(numberOfQuestions, len(testSet.Questions)))
	for i := range questionSet {
		qID := ids[i] - 1
		questionSet[i].Label = i + 1
		questionSet[i].Topic = testSet.Questions[qID].Topic
		questionSet[i].Statement = testSet.Questions[qID].Statement
		questionSet[i].Alternatives = testSet.Questions[qID].Alternatives
		questionSet[i].RightAnswer = testSet.Questions[qID].RightAnswer
	}

	return questionSet, nil
}
