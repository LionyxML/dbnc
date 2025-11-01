package logx

import "fmt"

func LogInfo(msg string) {
	fmt.Printf(">>> LOG [INFO]: %s", msg)
}

func LogError(msg string) {
	fmt.Printf(">>> LOG [ERROR]: %s", msg)
}
