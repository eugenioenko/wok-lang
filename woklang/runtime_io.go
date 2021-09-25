package woklang

import "fmt"

func RuntimePrint(interpreter *Interpreter, expressions []Expression) WokData {
	var result WokData
	for _, expr := range expressions {
		result = interpreter.Evaluate(expr)
		fmt.Println(result.ToString())
	}
	return result
}
