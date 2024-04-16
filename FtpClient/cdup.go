package main


func (cs *CommandsStruct) CDUP(args string) (string, error) {
	return writeAndreadOnMemory(cs, "CDUP ")
}