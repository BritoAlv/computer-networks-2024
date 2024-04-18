package core

func (cs *FtpSession) PWD(input string) (string, error) {	
	response, err := writeAndreadOnMemory(cs, "PWD ")
	if err != nil {
		return "", err
	}
	result := string(response)
	start := 0
	end := len(result)
	for i := 0; i < len(result); i++ {
		if result[i] == '"'{
			start = i+1
			for j := i+1; j < len(result); j++ {
				if result[j] == '"' {
					end = j
					break
				}
			}
			break
		}		
	}
	result = result[start:end]
	return result, nil
}