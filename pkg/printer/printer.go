package printer

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Colors functions for printing.
var (
	Yellow = color.New(color.FgYellow).SprintfFunc()
	Red    = color.New(color.FgRed).SprintfFunc()
	Green  = color.New(color.FgGreen).SprintfFunc()
)

// Info logs info text in console.
func Info(tag, text string) {
	fmt.Fprintf(color.Output, "%s: %s \n", Yellow(strings.ToUpper(tag)), Green(text))
}

// Fatal logs error text in console and quit the program.
func Fatal(tag string, err error, text ...string) {
	info := strings.Join(append(text, err.Error()), ":")
	fmt.Fprintf(color.Output, "%s: %s \n", Yellow(strings.ToUpper(tag)), Red(info))
	os.Exit(1)
}

// Error logs error text in console.
func Error(tag string, err error, text ...string) {
	info := strings.Join(append(text, err.Error()), ":")
	fmt.Fprintf(color.Output, "%s: %s \n", Yellow(strings.ToUpper(tag)), Red(info))
}
