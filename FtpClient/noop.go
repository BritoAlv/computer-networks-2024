package main

func (cs *CommandsStruct) NOOP(args string) (string, error) {
	return writeAndreadOnMemory(cs, "NOOP")
}