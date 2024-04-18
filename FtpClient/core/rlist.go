package core

import (
	"errors"
	"strings"
)

func (cs *FtpSession) RNLS(path string) (string, error) {
	split := strings.Split(path, " ")
	if len(split) >= 2 || len(split) == 0 {
		return "", errors.New("wrong Argument Format: rnls (path)*")
	}
	return recLS(cs, split[0], 0, "")
}

func recLS(cs *FtpSession, path string, i int, prev string) (string, error) {
	spacing := ""
	j := 0
	for {
		if j == i {
			break
		}
		spacing += "   "
		j++
	}

	folders, files , err  := Get_files_folders_Server(cs, path)
	if err != nil {
		return "", err
	}
	for index, fold := range folders {
		var marker string = ""
		if index == len(folders)-1 {
			marker = "└──"
		} else {
			marker = "├──"
		}
		prev += spacing + marker + fold + "\n"
		prev, err = recLS(cs, path+"/"+fold, i+1, prev)
		if err != nil {
			return "", errors.New("something wrong when recursion: " + path + " " +  err.Error())
		}
	}
	for index, file := range files {
		var marker string = ""
		if index == len(files)-1 {
			marker = "└──"
		} else {
			marker = "├──"
		}
		prev += spacing + marker + file + "\n"
	}
	return prev, nil
}
