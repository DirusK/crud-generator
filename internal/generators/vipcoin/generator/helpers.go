package generator

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// executeTemplate executes <data> structure in <resultPath> file.
func (g Generator) executeTemplate(templateFile, resultPath string, withGoImports bool, data interface{}) error {
	var sb strings.Builder
	if err := g.Template.ExecuteTemplate(&sb, templateFile, data); err != nil {
		return errors.Wrap(err, "can't execute template")
	}

	return writeToFile(resultPath, sb.String(), withGoImports)
}

func currentTimeForMigration() string {
	return time.Now().Format("20060102150105")
}

func runGoImports(generatedCode string) (string, error) {
	formatterCmd := exec.Command("goimports")
	// formatterCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // TODO: required for Windows

	stdinPipe, _ := formatterCmd.StdinPipe()

	var out, errout bytes.Buffer
	formatterCmd.Stdout = &out
	formatterCmd.Stderr = &errout

	err := formatterCmd.Start()
	if err != nil {
		return "", errors.Wrap(err, "can't start goimports command")
	}

	_, err = io.WriteString(stdinPipe, generatedCode)
	if err != nil {
		return "", err
	}

	err = stdinPipe.Close()
	if err != nil {
		return "", err
	}

	err = formatterCmd.Wait()
	if err != nil {
		return "", errors.Wrap(err, "invalid code in goimports command")
	}

	return out.String(), err
}

func writeToFile(path string, text string, withGoImports bool) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "cannot open file")
	}

	if withGoImports {
		text, err = runGoImports(text)
		if err != nil {
			return err
		}
	}

	_, err = file.WriteString(text)
	if err != nil {
		return errors.Wrap(err, "can't write generated code to file")
	}

	return nil
}
