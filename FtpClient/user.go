package main

func (cs *FtpSession) USER(username string) (string, error) {
	return writeAndreadOnMemory(cs, "USER " + username) 
}