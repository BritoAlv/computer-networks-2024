package main


func (cs *CommandsStruct) STAT(filename string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "STAT " + filename)
}