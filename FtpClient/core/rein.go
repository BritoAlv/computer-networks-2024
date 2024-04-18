package core

func (cs *FtpSession) REIN(args string) (string, error) {
	return writeAndreadOnMemory(cs, "REIN ")
}