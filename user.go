package main

func (cs *CommandsStruct) USER(username string) (string, error) {
	// first write to the server.
	resp, err1 := writeAndreadOnMemory(cs.connection, []byte("USER " + username + "\r\n"))
	if err1 != nil{
		return "", err1
	}
	return string(resp), nil 
}