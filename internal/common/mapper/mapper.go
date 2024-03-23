// Package mapper provides a mappers.
package mapper

import (
	"os"
	"strings"
)

// EnvMapper is a mapper for environment variables.
func EnvyMapper(placeholderName string) string {
	split := strings.Split(placeholderName, ":")
	defValue := ""
	if len(split) == 2 {
		placeholderName = split[0]
		defValue = split[1]
	}

	val, ok := os.LookupEnv(placeholderName)
	if !ok {
		return defValue
	}

	return val
}
