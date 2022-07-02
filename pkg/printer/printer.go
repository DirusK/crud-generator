package printer

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/martian/log"
)

// Colors functions for printing.
var (
	Yellow = color.New(color.FgYellow).SprintfFunc()
	Red    = color.New(color.FgRed).SprintfFunc()
	Green  = color.New(color.FgGreen).SprintfFunc()
)

// Info logs info text in console.
func Info(tag, text string) {
	_, err := fmt.Fprintf(color.Output, "%s: %s \n", Yellow(strings.ToUpper(tag)), Green(text))
	if err != nil {
		log.Errorf("printer error: %s", err)
	}
}

// Fatal logs error text in console and quit the program.
func Fatal(tag string, err error, text ...string) {
	info := strings.Join(append(text, err.Error()), ":")

	_, err = fmt.Fprintf(color.Output, "%s: %s \n", Yellow(strings.ToUpper(tag)), Red(info))
	if err != nil {
		log.Errorf("printer error: %s", err)
	}

	os.Exit(1)
}

// Error logs error text in console.
func Error(tag string, err error, text ...string) {
	info := strings.Join(append(text, err.Error()), ":")

	_, err = fmt.Fprintf(color.Output, "%s: %s \n", Yellow(strings.ToUpper(tag)), Red(info))
	if err != nil {
		log.Errorf("printer error: %s", err)
	}
}
