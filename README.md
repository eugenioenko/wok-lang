# wok-lang (WIP)

WokLang is a minimalist programming language crafted in go.
This project is being used to explore some of the complexities of
interpretation of a functional programing language.

WokLang as language is initially broadly inspired by Lisp but eventually will
end up diverging from it

## Building the project

> go build

## Building web assembly binary

> GOOS=js GOARCH=wasm go build cli/wasm/wok.go

## Running the project

> wok exec [filename] will execute the script

> wok eval [code] will execute the code passed as argument

## Updates

- Implemented tokenizer (scanner)
- Implemented parser (s-expression parser)
- Added rough runtime interpretation schema

## To-dos

- Function definition
- Scope creation
- tons of Predicates
