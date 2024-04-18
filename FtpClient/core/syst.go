package core

func (cs *FtpSession) SYST(arg string) (string, error) {
	return writeAndreadOnMemory(cs, "SYST ") 
}