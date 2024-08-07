package logger

import (
	"fmt"

	c "github.com/chrisbradleydev/go-dependency-injection/pkg/color"
)

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

func LogOutput(message string) {
	fmt.Println(message)
}

func LogOutputBlack(message string) {
	fmt.Printf("%s%s%s\n", c.Black, message, c.Reset)
}

func LogOutputRed(message string) {
	fmt.Printf("%s%s%s\n", c.Red, message, c.Reset)
}

func LogOutputGreen(message string) {
	fmt.Printf("%s%s%s\n", c.Green, message, c.Reset)
}

func LogOutputYellow(message string) {
	fmt.Printf("%s%s%s\n", c.Yellow, message, c.Reset)
}

func LogOutputBlue(message string) {
	fmt.Printf("%s%s%s\n", c.Blue, message, c.Reset)
}

func LogOutputMagenta(message string) {
	fmt.Printf("%s%s%s\n", c.Magenta, message, c.Reset)
}

func LogOutputCyan(message string) {
	fmt.Printf("%s%s%s\n", c.Cyan, message, c.Reset)
}

func LogOutputWhite(message string) {
	fmt.Printf("%s%s%s\n", c.White, message, c.Reset)
}
