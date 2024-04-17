package main

import "strings"

func (cs *FtpSession) STRU(args string) (string, error) {
	return writeAndreadOnMemory(cs, "STRU " + strings.TrimSpace(args))
}