package main


func (cs *FtpSession) CDUP(args string) (string, error) {
	return writeAndreadOnMemory(cs, "CDUP ")
}