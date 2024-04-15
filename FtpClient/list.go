package main

import "strings"

func (cs *CommandsStruct) LS(path string) (string, error) {
	// first try yo establish a PASSIVE Connection Data.
	connData, err := cs.PASV()
	if err != nil {
		return "", err
	}
	_, err = writeAndreadOnMemory(cs.connectionConfig, "LIST "+ strings.TrimSpace(path))
	if err != nil {
		return "", err
	}
	data, err := readOnMemoryPassive(connData)
	if err != nil {
		return "", err
	}
	_, err = readOnMemoryDefault(cs.connectionConfig)
	if err != nil {
		return "", err
	}
	defer (*connData).Close()
	return data, nil
}