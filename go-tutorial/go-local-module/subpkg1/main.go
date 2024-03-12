package subpkg1

import (
	"errors"
	"fmt"
)

func Subpkg1() string {
	return "subpkg1"
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
