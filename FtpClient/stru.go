package main

import "strings"

func (cs *CommandsStruct) STRU(args string) (string, error) {
	return writeAndreadOnMemory(cs, "STRU " + strings.TrimSpace(args))
}