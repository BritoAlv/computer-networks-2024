package main

import (
	"strings"
)

func (cs *CommandsStruct) NLST(input string) (string, error) {

	// Enter Passive Mode
	connData, passErr := cs.PASV()
	if passErr != nil {
		return "", passErr
	}

	_, lisErr := writeAndreadOnMemory(cs.connection, "NLST " + strings.TrimSpace(input))
	if lisErr != nil {
		return "", lisErr
	}

	data, dataErr := readOnMemoryPassive(connData)
	if dataErr != nil {
		return "", dataErr
	}

	_, doneErr := readOnMemoryDefault(cs.connection)
	if doneErr != nil {
		return "", doneErr
	}
	return data, nil
}
