package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LionyxML/dbnc/internal/checks"
	"github.com/LionyxML/dbnc/internal/lexer"
	"github.com/LionyxML/dbnc/internal/logx"
)

func main() {

	filename := os.Args
	checks.CheckArgs(filename)

	file, err := os.ReadFile(filename[1])
	checks.Check(err, 0, "Cant read the provided file... bye!")

	lines := strings.SplitSeq(string(file), "\n")

	for line := range lines {
		lexed, ok := lexer.Lexer(string(line))
		if ok {
			logx.LogInfo(fmt.Sprintf("%+v\n", lexed))
		}
	}

}
