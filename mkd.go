package main

import (
	"errors"
	"strings"
)

func (cs *CommandsStruct) MKD(input string) (string, error) {
	split := strings.Split(input, " ")

	if len(split) != 1 {
		return "", errors.New("invalid input")
	}

	response, err := writeAndreadOnMemory(cs.connection, []byte("MKD "+split[0]+"\r\n"))
	if err != nil {
		return "", err
	}

	return ParseFTPCode(string(response)[0:3])
}
