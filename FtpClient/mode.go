package main

import "strings"

func (cs *CommandsStruct) MODE(typeC string) (string, error) {
	return writeAndreadOnMemory(cs, "MODE " + strings.TrimSpace(typeC))
}