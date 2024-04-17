package main

import "strings"

func (cs *FtpSession) ACCT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "ACCT " + strings.TrimSpace(args))
}