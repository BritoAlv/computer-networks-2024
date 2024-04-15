package main

import "strings"

func (cs *CommandsStruct) APPEND(arg string) (string, error) {
	return cs.GET(append_flag + " " + strings.TrimSpace(arg))
}