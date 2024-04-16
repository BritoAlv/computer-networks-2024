package main

import "strings"

func (cs *CommandsStruct) LS(path string) (string, error) {
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
	_, err = readOnMemoryDefault(cs)
	if err != nil {
		return "", err
	}
	defer cs.release_connection()
	return data, nil
}