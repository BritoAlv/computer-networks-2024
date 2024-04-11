package main

func (cs *CommandsStruct) ANONYMOUS(args string) (string, error) {
	_, err := cs.USER("anonymous")
	if err != nil {
		return "", err
	}
	return cs.PASS("password")
}