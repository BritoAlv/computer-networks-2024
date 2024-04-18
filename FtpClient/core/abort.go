package core

func (cs *FtpSession) ABORT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "ABOR ")
}