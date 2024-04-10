package main

import (
	"errors"
	"os"
	"strings"
)

func (cs *CommandsStruct) RGET(path string) (string, error) {
	// create a folder with the name of the path.
	return cs.rGET(path, "RGET")
}
func (cs *CommandsStruct) rGET(serverpath string, foldername string) (string, error) {
	err := os.MkdirAll(foldername, 0777)
	if err != nil {
		return "", err
	}
	os.Chdir(foldername)
	defer os.Chdir("../")
	response, err := cs.LS(serverpath)
	if err != nil {
		return "", errors.New("something is wrong with LS")
	}
	archives := strings.Split(response, "\n")
	archivesFIltered := []string{}
	for _, arch := range archives {
		if len(arch) > 0 {
			archivesFIltered = append(archivesFIltered, arch)
		}
	}
	archives = archivesFIltered
	actual_path, err := cs.PWD("")
	if err != nil {
		return "", err
	}
	actual_path = strings.TrimSpace(actual_path)
	actual_path = actual_path[2:]
	actual_path = strings.Split(actual_path, "\"")[0]
	if len(archives) > 0 {
		for _, arch := range archives {
			arch = arch[:len(arch)-1]
			arch = strings.TrimSpace(arch)
			parts := strings.Split(arch, " ")
			filename := parts[len(parts)-1]
			if arch[0] == 'd' {
				cs.rGET(serverpath+"/"+filename, filename)
			} else {
				_, err := cs.CD(serverpath + "/")
				if err != nil {
					return "", err
				}
				_, err = cs.GET(filename + " B")
				if err != nil {
					return "", err
				}
				_, err = cs.CD(actual_path + "/")
				if err != nil {
					return "", err
				}
			}
		}
	}
	return "Done", nil
}
