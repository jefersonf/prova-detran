package cmd

import (
	"errors"
	"strconv"
	"strings"
)

func validateNumberOfQuestions(input string) error {
	n, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	if n < 1 || n > 90 {
		return errors.New("número de questões deve ser maior que zero e no máximo 90")
	}
	numQuestions = n
	return nil
}

func validateDuration(input string) error {
	durationString := strings.Trim(input, " mM")
	var err error
	mocktestDurationInMinutes, err = strconv.Atoi(durationString)
	if err != nil {
		return err
	}
	if mocktestDurationInMinutes < 1 || mocktestDurationInMinutes > 60 {
		return errors.New("duração inválida")
	}
	return nil
}

func indexToOptionLabel(offset int) string {
	return string('A' + offset)
}
