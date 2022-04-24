package logger

import "fmt"

func Log(a ...interface{}) {
	fmt.Fprintln(logger.out, a...)
}

func Err(a ...interface{}) {
	fmt.Fprint(logger.out, "[Error] ")
	fmt.Fprintln(logger.out, a...)
}

func Info(a ...interface{}) {
	fmt.Fprint(logger.out, "[Info] ")
	fmt.Fprintln(logger.out, a...)
}
