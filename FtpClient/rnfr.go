package main

func (cs *FtpSession) RNFR(oldName string) (string, error) {
	return writeAndreadOnMemory(cs, "RNFR " + oldName)
}