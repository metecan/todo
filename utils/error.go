package utils

import (
	"fmt"
	"os"
)

func Error(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
