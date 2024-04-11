package main

func (cs *CommandsStruct) NOOP(args string) (string, error) {
	return ParseFTPCode("200")
}