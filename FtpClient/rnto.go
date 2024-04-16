package main


func (cs *CommandsStruct) RNTO(newName string) (string, error) {
	return writeAndreadOnMemory(cs, "RNTO " + newName)
}