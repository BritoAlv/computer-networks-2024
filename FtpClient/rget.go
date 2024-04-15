package main

import (
	"errors"
	"os"
)

func (cs *CommandsStruct) RGET(path string) (string, error) {
	return cs.rGET("RGET", path)
}
func (cs *CommandsStruct) rGET(foldername string, path string) (string, error) {
	err := os.MkdirAll(foldername, 0777)
	if err != nil {
		return "", err
	}
	os.Chdir(foldername)
	defer os.Chdir("../")
	folders, files , err  := get_files_folders_current(cs, path)
	if err != nil {
		return "", err
	}
	for _, fold := range folders {
		_, err = cs.rGET(fold, path+"/"+fold)
		if err != nil {
			return "", errors.New("something wrong when recursion: " + path + " " +  err.Error())
		}
	}
	for _, file := range files {
		_, err = cs.GET(path + "/" + file + " B")
				if err != nil {
					return "", err
			}
	}
	return "Everything seems oky", nil
}