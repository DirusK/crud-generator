package generator

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// executeTemplate open template in <templatePath> and executes <data> structure in <resultPath> file.
func (g Generator) executeTemplate(templatePath string, resultPath string, withGoImports bool, data interface{}) error {
	tmpl, err := readFile(templatePath)
	if err != nil {
		return err
	}

	if _, err = g.Template.Parse(tmpl); err != nil {
		return errors.Wrap(err, "can't parse template")
	}

	file, err := openFile(resultPath)
	if err != nil {
		return err
	}

	var sb strings.Builder

	if err = g.Template.Execute(&sb, data); err != nil {
		return errors.Wrap(err, "can't execute template")
	}

	return writeToFile(file, sb.String(), withGoImports)
}

func currentTimeForMigration() string {
	return time.Now().Format("20060102150105")
}

func readFile(filepath string) (string, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", errors.Wrap(err, "cannot read file")
	}

	return string(file), nil
}

func openFile(filepath string) (*os.File, error) {
	var file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open file")
	}

	return file, nil
}

func runGoImports(generatedCode string) (string, error) {
	formatterCmd := exec.Command("goimports")
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

func writeToFile(file *os.File, text string, withGoImports bool) error {
	var err error

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
