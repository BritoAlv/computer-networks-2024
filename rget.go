package main

import (
	"errors"
	"os"
	"strings"
)

func (cs *CommandsStruct) RGET(path string) (string, error) {
	actual_path, err := cs.PWD("")
	if err != nil {
		return "", err
	}
	actual_path = strings.TrimSpace(actual_path)
	actual_path = actual_path[2:]
	actual_path = strings.Split(actual_path, "\"")[0]
	defer cs.CD(actual_path)
	if len(path) > 0 {
		_, err = cs.CD(actual_path + "/" + path)
		if err != nil {
			return "", err
		}
	}
	return cs.rGET("RGET")
}
func (cs *CommandsStruct) rGET(foldername string) (string, error) {
	err := os.MkdirAll(foldername, 0777)
	if err != nil {
		return "", err
	}
	os.Chdir(foldername)
	defer os.Chdir("../")
	response1, err := cs.LS("")
	if err != nil {
		return "", errors.New("something is wrong with LS")
	}
	response2, err := cs.NLST("")
	if err != nil {
		return "",  errors.New("something is wrong with NLS")
	}
	archives1 := strings.Split(response1, "\n")
	archives2 := strings.Split(response2, "\n")
	archivesFiltered1 := []string{}
	archivesFiltered2 := []string{}
	for index, arch := range archives1 {
		if len(arch) > 0 {
			archivesFiltered1 = append(archivesFiltered1, arch)
			archivesFiltered2 = append(archivesFiltered2, archives2[index])
		}
	}
	archives1 = archivesFiltered1
	if len(archives1) > 0 {
		for index, arch := range archives1 {
			parts := strings.Split(archivesFiltered2[index], "/")
			filename := parts[len(parts)-1]
			filename = filename[:len(filename)-1]
			if arch[0] == 'd' {
				_, err = cs.CD(filename)
				if err != nil {
					return "", err
				}
				_, err := cs.rGET(filename)
				if err != nil {
					return "", err
				}
				_, err = cs.CDUP("")
				if err != nil {
					return "", err
				}
			} else {
				_, err = cs.GET(filename + " B")
				if err != nil {
					return "", err
				}
			}
		}
	}
	return "Done", nil
}