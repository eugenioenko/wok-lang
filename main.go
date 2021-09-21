package main

import (
	"fmt"
	"os"
	"strings"
	"wok/woklang"
)

func main() {

	args := os.Args
	if len(args) <= 2 || args[1] == "help" {
		help()
		return
	}
	if args[1] == "exec" {
		woklang.Exec(args[2])
		return
	}
	if args[1] == "eval" {
		woklang.Eval(strings.Join(args[2:], " "))
		return
	}

}

func help() {
	fmt.Print(`Wok is an interpreter of WokLang programming language.

Usage:

  go <command> [arguments]

  The commands are:

      exec [filename]     executes the script with filename
      eval [code]         executes the code passed as argument
      help                prints this message
`)
}
