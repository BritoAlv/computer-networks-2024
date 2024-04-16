package main

import "strings"

func (cs *CommandsStruct) TYPE(typeC string) (string, error) {
	return writeAndreadOnMemory(cs, "TYPE " + strings.TrimSpace(typeC))
}