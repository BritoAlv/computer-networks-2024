package core

func (cs *FtpSession) NOOP(args string) (string, error) {
	return writeAndreadOnMemory(cs, "NOOP")
}