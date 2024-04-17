package main

import "strings"

func (cs *FtpSession) TYPE(typeC string) (string, error) {
	return writeAndreadOnMemory(cs, "TYPE " + strings.TrimSpace(typeC))
}