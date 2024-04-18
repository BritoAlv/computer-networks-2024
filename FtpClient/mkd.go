package main

func (cs *FtpSession) MKD(input string) (string, error) {
	return writeAndreadOnMemory(cs, "MKD "+input)
}