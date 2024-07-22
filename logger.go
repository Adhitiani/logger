package logger

import "fmt"

func Info(data any) {
    fmt.Println("Log info: ", data)
}

func Error(data any) {
    fmt.Println("Log error: ", data)
}

func Warning(data any) {
    fmt.Println("Log warning: ", data)
}

func Fatal(data any) {
    fmt.Println("Log fatal: message", data)
}