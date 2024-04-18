package core

import "strings"

func (cs *FtpSession) SITE(args string) (string, error) {
	return writeAndreadOnMemory(cs, "SITE " + strings.TrimSpace(args))
}