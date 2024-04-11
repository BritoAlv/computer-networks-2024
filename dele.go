package main

func (cs *CommandsStruct) DELE(input string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("DELE "+input+"\r\n"))
	if err != nil {
		return "", err
	}

	return ParseFTPCode(string(response)[0:3])
}
