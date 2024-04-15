package main

func (cs *CommandsStruct) HELP(args string) (string, error) {
	return writeAndreadOnMemory(cs.connectionConfig, "HELP ")
}
