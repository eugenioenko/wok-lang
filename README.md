# wok-lang (WIP)

WokLang is a minimalist programming language crafted in go.
This project is being used to explore some of the complexities of
interpretation of a functional programing language.

WokLang as language is initially broadly inspired by Lisp but eventually will
end up diverging from it

### > [Try it out in webassembly playground!](https://eugenioenko.github.io/wok-lang/)

## Building the project

> go build

## Building web assembly binary

> GOOS=js GOARCH=wasm go build cli/wasm/wok.go

or using tinygo (currently not working)

> tinygo build -o ./live/wok.wasm -target wasm ./cli/wasm/wasm.go

## Running the project

> wok exec [filename] will execute the script

> wok eval [code] will execute the code passed as argument

## Running playground locally

Because WebAssembly.instantiateStreaming requires CORS, a local server is requiredcd

> npm install http-server
> cd live
> http-server

## Updates

- Implemented tokenizer (scanner)
- Implemented parser (s-expression parser)
- Added rough runtime interpretation schema
- Added conditional expressions
- Added loop expressions
- Added math expressions
- Added func definition expression
