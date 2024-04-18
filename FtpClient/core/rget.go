package core

import (
	"errors"
	"os"
	"strings"
)

func (cs *FtpSession) RGET(path string) (string, error) {
	parts := SplitString(path, '&')
	if len(parts) == 1 || parts[1] == "" {
		return cs.rGET("RGET", path)
	}
	part := strings.Split(parts[0], "/")

	return cs.rGET(parts[1] + "/" + part[len(part)-1], parts[0])
}
func (cs *FtpSession) rGET(foldername string, path string) (string, error) {
	os.Mkdir(foldername, 0777)
	err := os.Chdir(foldername)
	if err != nil {
		return "", err
	}
	folders, files , err  := Get_files_folders_Server(cs, path)
	if err != nil {
		return "", err
	}
	for _, fold := range folders {
		_, err = cs.rGET(foldername + "/" + fold, path+"/"+fold)
		if err != nil {
			return "", errors.New("something wrong when recursion: " + path + " " +  err.Error())
		}
	}
	for _, file := range files {
		_, err = cs.GET(path + "/" + file + "&" + foldername + "/")
				if err != nil {
					return "", err
			}
	}
	return "Everything seems oky", nil
}