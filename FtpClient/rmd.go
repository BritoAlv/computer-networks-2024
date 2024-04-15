package main

func (cs *CommandsStruct) RMD(input string) (string, error) {
	return RMD_Recursive(cs, input)
}

func RMD_Recursive(cs *CommandsStruct, directory string) (string, error) {

	
	folders, files , err  := get_files_folders_current(cs, directory)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(files); i++ {
		_, err := writeAndreadOnMemory(cs.connectionConfig, "DELE "+directory+"/"+files[i]+"\r\n")
		if err != nil {
			return "", err
		}
	}

	for i := 0; i < len(folders); i++ {
		_, err := RMD_Recursive(cs, directory+"/"+folders[i])
		if err != nil {
			return "", err
		}
	}

	return writeAndreadOnMemory(cs.connectionConfig, "RMD "+ directory)
}