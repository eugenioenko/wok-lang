package testing

import (
	"testing"
	"wok/woklang"
)

func TestDefineAFunction(t *testing.T) {
	source := `
		(func function (a b c)
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
		(func function (a b c)
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
		(func function (a b c)
			(return c)
		)
		(debug (function 1 2))
	`
	result := woklang.Eval(source)
	if result.GetType() != woklang.WokTypeNull {
		t.Fail()
	}
}

func TestFunctionShouldReturnFromInner(t *testing.T) {
	source := `
		(func function (a b c)
			(func inner (x y z)
				(return-from function 777)
			)
			(inner 1 2 3)
			(return c)
		)
		(debug (function 1 2 3))
	`
	result := woklang.Eval(source)
	if result.ToInteger() != 777 {
		t.Fail()
	}
}

func TestFunctionShouldThrow(t *testing.T) {
	source := `
		(func function (a b c)
			(func inner (x y z)
				(return-from function_does_not_exist 777)
			)
			(inner 1 2 3)
			(return c)
		)
		(debug (function 1 2 3))
	`
	result := woklang.Eval(source)
	if result.GetType() != woklang.WokTypeException {
		t.Fail()
	}
}
