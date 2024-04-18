package core

func (cs *FtpSession) HELP(args string) (string, error) {
	return writeAndreadOnMemory(cs, "HELP ")
}
