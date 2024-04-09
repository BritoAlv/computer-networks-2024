package main

import (
	"errors"
	"strings"
)

func (cs *CommandsStruct) NLST(input string) (string, error) {
	split := strings.Split(input, " ")

	if len(split) != 1 {
		return "", errors.New("invalid input")
	}

	// Enter Passive Mode
	connData, passErr := cs.PASV()
	if passErr != nil {
		return "", passErr
	}

	_, lisErr := writeAndreadOnMemory(cs.connection, []byte("NLST "+"\r\n"))
	if lisErr != nil {
		return "", lisErr
	}

	data, dataErr := readOnMemory(connData)
	if dataErr != nil {
		return "", dataErr
	}

	response, doneErr := readOnMemory(cs.connection)
	if doneErr != nil {
		return "", doneErr
	}

	if starts_with(string(response), "226") {
		return string(data), nil
	} else {
		return "Wrong: " + string(response)[3:], nil
	}
}
