package core

import (
	"strings"
)

func (cs *FtpSession) LS(path string) (string, error) {
	// first try yo establish a PASSIVE Connection Data.
	err := cs.check_connection()
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs, "LIST "+ strings.TrimSpace(path))
	if err != nil {
		return "", err
	}
	data, err := readOnMemoryPassive(cs.connectionData)
	if err != nil {
		return "", err
	}
	cs.release_connection()
	_, err = readOnMemoryDefault(cs)
	if err != nil {
		return "", err
	}
	return data, nil
}

