package util

import (
	"fmt"
	"github.com/fatih/color"
)

func Info(message string) {
	fmt.Fprintf(color.Output, "%v %v\n", color.GreenString("!"), message)
}

func Warn(message string) {
	fmt.Fprintf(color.Output, "%v %v\n", color.YellowString("!"), message)
}

func Fatal(message string) {
	fmt.Fprintf(color.Output, "%v %v\n", color.RedString("!"), color.HiRedString(message))
}