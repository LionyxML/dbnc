package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Token struct {
	Type  string
	Value any
}

type LexerOut struct {
	Input  string
	Output []Token
}

type ok = bool

func Lexer(line string) (LexerOut, ok) {
	outputs := []Token{}
	re := regexp.MustCompile(`\s+`)
	tokens := re.Split(line, -1)

	if len(tokens) == 0 || len(line) == 0 {
		return LexerOut{}, false
	}

	for _, part := range tokens {
		if _, err := strconv.Atoi(part); err == nil {
			outputs = append(
				outputs,
				Token{
					Type:  "word",
					Value: part,
				})
		} else {
			outputs = append(
				outputs,
				Token{
					Type:  "number",
					Value: part,
				})
		}
	}

	lexedLine := LexerOut{
		Input:  line,
		Output: outputs,
	}

	return lexedLine, true
}

func check(err error, code int, msg string) {
	if err != nil {
		logError(fmt.Sprintf(">>> %+v\n", err))
		logError(fmt.Sprintf(">>> %s\n", msg))
		os.Exit(code)
	}
}

func checkArgs(args []string) {
	argsLenght := len(args)
	msg := "Error while reading program arguments"

	if (argsLenght) > 2 {
		check(fmt.Errorf("More than 1 argument provided"), 1, msg)
	}

	if (argsLenght) <= 1 {
		check(fmt.Errorf("No argument provided"), 2, msg)
	}

	if len(args) != 2 {
	}
}

func logInfo(msg string) {
	fmt.Printf(">>> LOG [INFO]: %s", msg)
}

func logError(msg string) {
	fmt.Printf(">>> LOG [ERROR]: %s", msg)
}

func main() {

	filename := os.Args
	checkArgs(filename)

	file, err := os.ReadFile(filename[1])
	check(err, 0, "Cant read the provided file... bye!")

	lines := strings.SplitSeq(string(file), "\n")

	for line := range lines {
		lexed, ok := Lexer(string(line))
		if ok {
			logInfo(fmt.Sprintf("%+v\n", lexed))
		}
	}

}
