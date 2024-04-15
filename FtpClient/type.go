package main

import "strings"

func (cs *CommandsStruct) TYPE(typeC string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "TYPE " + strings.TrimSpace(typeC))
}