package testing

import (
	"fmt"
	"testing"
	"wok/woklang"
)

func TestDefineAFunction(t *testing.T) {
	source := `
		(defun function (a b c)
			(return c)
		)
		(debug function)
	`
	result := woklang.Eval(source)
	if result.GetType() != woklang.WokTypeFunction {
		t.Fail()
	}
}

func TestFunctionShouldReturnValue(t *testing.T) {
	source := `
		(defun function (a b c)
			(return c)
		)
		(debug (function 1 2 777))
	`
	result := woklang.Eval(source)
	if result.ToInteger() != 777 {
		t.Fail()
	}
}

func TestFunctionShouldNullUndefinedParams(t *testing.T) {
	source := `
		(defun function (a b c)
			(return c)
		)
		(debug (function 1 2))
	`
	result := woklang.Eval(source)
	if result.GetType() != woklang.WokTypeNull {
		t.Fail()
	}
}

func TestFunctionShouldFromInner(t *testing.T) {
	source := `
		(defun function (a b c)
			(defun inner (x y z)
				(return-from function 777)
			)
			(inner 1 2 3)
			(return c)
		)
		(debug (function 1 2 3))
	`
	result := woklang.Eval(source)
	fmt.Print(result.ToInteger())
	if result.ToInteger() != 777 {
		t.Fail()
	}
}
