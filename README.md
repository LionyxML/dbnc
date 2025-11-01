# A minimal (only 3 keywords) DBN (Drawy By Numeral) compiler written in Go

A basic DBN compiler written in Go, as an exercise to learn the language.

This project is inspired by the great article written by Mariko Kaneko: [How to
be a
compiler](https://medium.com/@kosamari/how-to-be-a-compiler-make-a-compiler-with-javascript-4a8a13d473b4).

Goals:

- [ ] Read a .dbn file and generate an SVG image.
- [ ] Read a serve an web page with the SVG image.
- [ ] Alter the web page to contain a small editor from which the compiler will
      read.

For the compiler:

- [x] Lexer
- [ ] Parser
- [ ] Transformer
- [ ] Generator

For the webserver:

- [ ] Server for public folder
- [ ] Main page to host the editor and output
- [ ] Connect the editor to the output

For the project:

- [ ] Use only built-in go features
- [ ] Exception to make and Makefile to store all commands

## Run the compiler cli

### While developing

Use:

```go
go run ./cmd/dbnc/main.go ./docs/example_programs/prog01.dbn
```
