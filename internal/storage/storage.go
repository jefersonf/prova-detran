package storage

import (
	"errors"
	"time"

	"github.com/jefersonf/prova-detran/internal"
)

type Repository interface {
	SaveMocktest(internal.Mocktest) error
	GetMocktests(...FilterOption) ([]internal.Mocktest, error)
}

type FilterOption func(*Filter) error

type Filter struct {
	fromDate              time.Time
	orderByCorrectAnswers bool
}

func FromDate(date time.Time) FilterOption {
	return func(f *Filter) error {
		if date.After(time.Now()) {
			return errors.New("from date must be a date before the mocktest realization")
		}
		f.fromDate = date
		return nil
	}
}

func OrderByCorrectAnswers() FilterOption {
	return func(f *Filter) error {
		f.orderByCorrectAnswers = true
		return nil
	}
}
