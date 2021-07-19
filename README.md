# Lexeme and Token Translator
The idea of this project is to show some steps that the compiler does, for example create a symbol table with the occurrences and a token flow.

## Dependencies and Technologies used

- __[Go 1.11](https://golang.org/doc/go1.11)__ - This project uses [go modules](https://tip.golang.org/doc/go1.11#modules) as a package manager, so you need to use at least [Go 1.11](https://golang.org/doc/go1.11) for development mode.

## Running locally

**Development mode**
```bash
mkdir results
go run main.go [path/to/the/file]
```

**Teacher Mode** - _As it is an University project, and to avoid the need of installing GoLang locally, I am sending the binaries to the most common OS_

Just run in your command line:
```bash
mkdir results
executable/lexicalAnalysis-[your-OS] [path/to/the/file]
```

Example:
```bash
executable/lexicalAnalysis-Windows-Intel.exe examples/math/simple.txt
```

It will generate 2 new files that you can find at `results`:

    symbolTable_[youFileName]
    tokenFlow_[youFileName]
  
PS: We are facing some problems with LF and CRLF files with Notepad for breaking lines, so we higlhy encourage the usage of another text file editor (Eg. Notepad++).

![gopher-working-test](go-wip.gif)
