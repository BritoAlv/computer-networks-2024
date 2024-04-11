package main

import "strings"

func (cs *CommandsStruct) LS(path string) (string, error) {
	// first try yo establish a PASSIVE Connection Data.
	connData, err := cs.PASV()
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connection, []byte("LIST "+ strings.TrimSpace(path) + "\r\n"))
	if err != nil {
		return "", err
	}
	data, err := readOnMemory(connData)
	if err != nil {
		return "", err
	}
	_, err = readOnMemory(cs.connection)
	if err != nil {
		return "", err
	}
	defer (*connData).Close()
	return string(data), nil
}