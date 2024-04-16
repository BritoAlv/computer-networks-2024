package main

func (cs *CommandsStruct) ABORT(args string) (string, error) {
	return writeAndreadOnMemory(cs, "ABOR ")
}