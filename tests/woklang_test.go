package testing

import (
	"testing"
	"wok/woklang"
)

func TestAddition(t *testing.T) {
	source := `
			(+ 10 1 20)
		`
	total := woklang.Eval(source).ToInteger()
	if total != 31 {
		t.Fail()
	}
}

func TestAddition2(t *testing.T) {
	source := `
			(+ 10 1 20)
		`
	total := woklang.Eval(source).ToInteger()
	if total != 31 {
		t.Fail()
	}
}
