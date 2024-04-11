package main

func (cs *CommandsStruct) MKD(input string) (string, error) {
	response, err := writeAndreadOnMemory(cs.connection, []byte("MKD "+input+"\r\n"))
	if err != nil {
		return "", err
	}

	return ParseFTPCode(string(response)[0:3])
}
