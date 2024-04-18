package main

import (
	"strings"
)

func (cs *FtpSession) NLST(input string) (string, error) {

	// Enter Passive Mode
	err := cs.check_connection()
	if err != nil {
		return "", err
	}
	_, lisErr := writeAndreadOnMemory(cs, "NLST " + strings.TrimSpace(input))
	if lisErr != nil {
		return "", lisErr
	}

	data, dataErr := readOnMemoryPassive(cs.connectionData)
	if dataErr != nil {
		return "", dataErr
	}
	cs.release_connection()
	_, doneErr := readOnMemoryDefault(cs)
	if doneErr != nil {
		return "", doneErr
	}
	return data, nil
}
