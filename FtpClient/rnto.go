package main


func (cs *FtpSession) RNTO(newName string) (string, error) {
	return writeAndreadOnMemory(cs, "RNTO " + newName)
}