package checks

import (
	"fmt"
	"os"

	"github.com/LionyxML/dbnc/internal/logx"
)

func CheckErr(err error, code int, msg string, logger logx.Logger) {
	if err != nil {
		logger.Error(">>> %+v", err)
		logger.Error(">>> %s", msg)
		os.Exit(code)
	}
}

func CheckArgs(args []string, logger logx.Logger) {
	argsLenght := len(args)
	msg := "Error while reading program arguments"

	if (argsLenght) > 2 {
		CheckErr(fmt.Errorf("More than 1 argument provided"), 1, msg, logger)
	}

	if (argsLenght) <= 1 {
		CheckErr(fmt.Errorf("No argument provided"), 2, msg, logger)
	}

	if len(args) != 2 {
	}
}
