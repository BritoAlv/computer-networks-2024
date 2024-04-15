package main

import "strings"

func (cs *CommandsStruct) CD(args string) (string, error) {
	return writeAndreadOnMemory(cs.connection, "CWD " + strings.TrimSpace(args))
}