package main

func (cs *CommandsStruct) SYST(arg string) (string, error) {
	return writeAndreadOnMemory(cs, "SYST ") 
}