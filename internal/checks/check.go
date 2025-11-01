package checks

import (
	"fmt"
	"os"

	"github.com/LionyxML/dbnc/internal/logx"
)

func Check(err error, code int, msg string) {
	if err != nil {
		logx.LogError(fmt.Sprintf(">>> %+v\n", err))
		logx.LogError(fmt.Sprintf(">>> %s\n", msg))
		os.Exit(code)
	}
}

func CheckArgs(args []string) {
	argsLenght := len(args)
	msg := "Error while reading program arguments"

	if (argsLenght) > 2 {
		Check(fmt.Errorf("More than 1 argument provided"), 1, msg)
	}

	if (argsLenght) <= 1 {
		Check(fmt.Errorf("No argument provided"), 2, msg)
	}

	if len(args) != 2 {
	}
}
