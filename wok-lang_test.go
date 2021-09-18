package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	source := `
			(+ 10 1 20)
		`
	total := WokEvaluate(source).ToInteger()
	if total != 31 {
		t.Fail()
	}
}
