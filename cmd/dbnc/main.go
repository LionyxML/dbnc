package main

import (
	"os"

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

	_, ok := lexer.Lexer(file, cliLogger)

	if ok {
		cliLogger.Debug("It worked!")
	} else {
		cliLogger.Error("Oh no, something went wrong")
	}

}
