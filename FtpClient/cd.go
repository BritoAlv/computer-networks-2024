package main

import "strings"

func (cs *CommandsStruct) CD(args string) (string, error) {
	return writeAndreadOnMemory(cs, "CWD " + strings.TrimSpace(args))
}