package models

import "fmt"

func toCopyright(copyright string) string {
	if copyright == "" {
		return copyright
	}

	return fmt.Sprintf("/*\n * %s\n */", copyright)
}
