package main

import (
	"os"
	"strings"

	"github.com/LionyxML/dbnc/internal/checks"
	"github.com/LionyxML/dbnc/internal/lexer"
	"github.com/LionyxML/dbnc/internal/logx"
)

func main() {

	cliLogger := logx.New("[dbnc]", "debug", true)

	filename := os.Args
	checks.CheckArgs(filename, cliLogger)

	file, err := os.ReadFile(filename[1])
	checks.CheckErr(err, 0, "Cannot read the provided file... bye!", cliLogger)

	lines := strings.SplitSeq(string(file), "\n")

	for line := range lines {
		lexed, ok := lexer.Lexer(line)

		if ok {
			cliLogger.Debug("%+v", lexed)
		}
	}

}
