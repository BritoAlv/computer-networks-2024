package main

import (
	"errors"
)

func (cs *CommandsStruct) PWD(input string) (string, error) {
	if input != "" {
		return "", errors.New("invalid input")
	}

	response, err := writeAndreadOnMemory(cs.connection, []byte("PWD "+"\r\n"))
	if err != nil {
		return "", err
	}

	if starts_with(string(response), "257") {
		return string(response)[3:], nil
	} else {
		return "Wrong: " + string(response)[3:], nil
	}
}
