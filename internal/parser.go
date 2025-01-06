package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var errContext = errors.New("parser error")

func ParseQuestionSet(reader io.Reader) (*QuestionSet, error) {
	var questionSet QuestionSet
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return &questionSet, withContext(err)
	}
	if err := json.Unmarshal(bytes, &questionSet); err != nil {
		return &questionSet, withContext(err)
	}
	return &questionSet, nil
}

func withContext(err error) error {
	return fmt.Errorf("%w: %w", errContext, err)
}
