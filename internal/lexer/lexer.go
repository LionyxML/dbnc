package lexer

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/LionyxML/dbnc/internal/logx"
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

func Lexer(code []byte, cliLogger logx.Logger) ([]LexerOut, ok) {
	lexedLines := []LexerOut{}
	ok_flag := true

	lines := strings.SplitSeq(string(code), "\n")

	for line := range lines {

		cliLogger.Debug("LEXER: LINE: %+v", line)

		outputs := []Token{}
		re := regexp.MustCompile(`\s+`)
		tokens := re.Split(line, -1)

		if len(tokens) == 0 || len(line) == 0 {
			continue
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

		lexedLines = append(
			lexedLines,
			LexerOut{
				Input:  line,
				Output: outputs,
			})
	}

	cliLogger.Debug("LEXER: OUT: %+v", lexedLines)

	return lexedLines, ok_flag
}
