package main

import (
	"errors"
	"strings"
)

func (cs *CommandsStruct) RNLS(path string) (string, error) {
	split := strings.Split(path, " ")
	if len(split) >= 2 || len(split) == 0 {
		return "", errors.New("wrong Argument Format: rnls (path)*")
	}
	return recLS(cs, split[0], 0, "")
}

func recLS(cs *CommandsStruct, path string, i int, prev string) (string, error) {

	spacing := ""
	j := 0
	for {
		if j == i {
			break
		}
		spacing += "   "
		j++
	}

	response1, err := cs.LS(path)
	if err != nil {
		return "", errors.New("something is wrong with LS")
	}
	response2, err := cs.NLST(path)
	if err != nil {
		return "", errors.New("something is wrong with NLS")
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
			var marker string = ""
			if index == len(archives1)-1 {
				marker = "└──"
			} else {
				marker = "├──"
			}
			if arch[0] == 'd' {
				prev += spacing + marker + filename + "\n"
				prev, err = recLS(cs, path+"/"+filename, i+1, prev)
				if err != nil {
					return "", errors.New("something wrong with recursion")
				}
			} else {
				prev += spacing + marker + filename + "\n"
			}
		}
	}
	return prev, nil
}
