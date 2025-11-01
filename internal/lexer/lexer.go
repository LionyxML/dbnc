package lexer

import (
	"regexp"
	"strconv"
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
