package utils

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func ReadInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)

	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	todoLabel := scanner.Text()

	if len(todoLabel) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return todoLabel, nil

}
