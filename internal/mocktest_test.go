package internal

import (
	"testing"
)

func TestNewMocktest(t *testing.T) {
	_, err := NewMocktest(1)
	if err != nil {
		t.Error(err)
	}
}
