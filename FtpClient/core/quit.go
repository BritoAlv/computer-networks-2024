package core

import (
)

func (cs *FtpSession) QUIT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "QUIT ")
}