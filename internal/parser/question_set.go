package parser

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
		return &qset, withCtx(err)
	}
	if err := json.Unmarshal(bytes, &qset); err != nil {
		return &qset, withCtx(err)
	}
	return &qset, nil
}

func withCtx(err error) error {
	return fmt.Errorf("%w: %w", errContext, err)
}
