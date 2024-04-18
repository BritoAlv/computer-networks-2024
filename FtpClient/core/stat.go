package core


func (cs *FtpSession) STAT(filename string) (string, error) {
	return writeAndreadOnMemory(cs, "STAT " + filename)
}