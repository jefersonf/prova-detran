package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var errContext = errors.New("parser error")

func ParseQuestionSet(reader io.Reader) (*QuestionSet, error) {
	var qset QuestionSet
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return &qset, withContext(err)
	}
	if err := json.Unmarshal(bytes, &qset); err != nil {
		return &qset, withContext(err)
	}
	return &qset, nil
}

func withContext(err error) error {
	return fmt.Errorf("%w: %w", errContext, err)
}
