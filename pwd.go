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

	return ParseFTPCode(string(response)[0:3])
}
