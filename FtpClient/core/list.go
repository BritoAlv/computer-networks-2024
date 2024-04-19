package core

import (
	"strings"
)

func (cs *FtpSession) LS(path string) (string, error) {
	// first try yo establish a PASSIVE Connection Data.
	err := cs.check_connection()
	if err != nil {
		cs.release_connection()
		return "", err
	}
	_, err = writeAndreadOnMemory(cs, "LIST "+ strings.TrimSpace(path))
	if err != nil {
		cs.release_connection()
		return "", err
	}
	data, err := readOnMemoryPassive(cs.connectionData)
	if err != nil {
		cs.release_connection()
		return "", err
	}
	cs.release_connection()
	_, err = readOnMemoryDefault(cs)
	if err != nil {
		return "", err
	}
	return data, nil
}

