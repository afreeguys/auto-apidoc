package log

import "fmt"

func Info(format string, argv ...interface{}) {
	fmt.Printf("[INFO]"+format+"\n", argv...)
}

func Error(format string, argv ...interface{}) {
	fmt.Printf("[ERROR]"+format+"\n", argv...)
}
