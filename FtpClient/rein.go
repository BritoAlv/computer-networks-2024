package main

func (cs *CommandsStruct) REIN(args string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "REIN ")
}