package main

import (
	"os"
)

func (cs *CommandsStruct) QUIT(args string) (string, error) {
	defer os.Exit(0)
	return writeAndreadOnMemory(cs, "QUIT ")
}