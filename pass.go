package main

func (cs *CommandsStruct) PASS(args string) (string, error)  {
	response, err := writeAndreadOnMemory(cs.connection, []byte("PASS " + args + "\r\n"))
	if err != nil{
		return "", err
	}
	if starts_with(string(response), "230") {
		return string(response)[3:], nil
	} else {
		return "Wrong: " + string(response)[3:], nil
	}
}