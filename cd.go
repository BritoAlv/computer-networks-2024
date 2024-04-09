package main

func (cs *CommandsStruct) CD(args string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("CWD " + args + "\r\n"))
	if err != nil{
		return "There was something wrong", err
	}
	if starts_with(string(response), "250") {
		return "OK " + string(response)[4:], nil
	} else {
		return "?", nil
	}
}