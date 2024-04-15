package main

func (cs *CommandsStruct) USER(username string) (string, error) {
	return writeAndreadOnMemory(cs.connection, "USER " + username) 
}