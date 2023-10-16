package logic

import "fmt"


func LogOutput(message string) {
	fmt.Println(message)
}

type LoggerAdapter func(message string)

func(lg LoggerAdapter) Log(message string) {
	lg(message)
}