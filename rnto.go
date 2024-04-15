package main


func (cs *CommandsStruct) RNTO(newName string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "RNTO " + newName)
}