package core

import (
	"strings"
)

func (cs *FtpSession) NLST(input string) (string, error) {

	// Enter Passive Mode
	err := cs.check_connection()
	if err != nil {
		cs.release_connection()
		return "", err
	}
	_, lisErr := writeAndreadOnMemory(cs, "NLST " + strings.TrimSpace(input))
	if lisErr != nil {
		cs.release_connection()
		return "", lisErr
	}

	data, dataErr := readOnMemoryPassive(cs.connectionData)
	if dataErr != nil {
		cs.release_connection()
		return "", dataErr
	}
	cs.release_connection()
	_, doneErr := readOnMemoryDefault(cs)
	if doneErr != nil {
		cs.release_connection()
		return "", doneErr
	}
	return data, nil
}
