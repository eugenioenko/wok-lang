package main

import (
	"syscall/js"
	"wok/woklang"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func wokJsEval(this js.Value, args []js.Value) interface{} {
	return woklang.Eval(args[0].String()).ToString()
}

func main() {
	js.Global().Set("wok", js.FuncOf(wokJsEval))
	<-c
}
