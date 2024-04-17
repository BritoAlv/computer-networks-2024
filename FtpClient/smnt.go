package main

import "strings"

func (cs *CommandsStruct) SMNT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "SMNT " + strings.TrimSpace(args))
}