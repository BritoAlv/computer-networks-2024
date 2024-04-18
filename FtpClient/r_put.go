package main

import (
	"errors"
)

func (cs *FtpSession) RPUT(path string) (string, error) {
	parts := SplitString(path, '&')
	if len(parts) == 1 || parts[1] == "" {
		return "", errors.New("wrong arguments")
	}
	return cs.rPUT(parts[0], parts[1])
}

func (cs *FtpSession) rPUT(directorySource string, directoryDestination string) (string, error) {
	current_folder_path := directoryDestination + "/" + get_filename_path(directorySource) 
	_, err := cs.MKD(current_folder_path)
	if err != nil {
		return "", err
	}
	folders, files, err := get_files_folders_local(directorySource)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		_, err = cs.PUT(directorySource + "/" + file + "&" + current_folder_path + "/" + file)
		if err != nil {
			return "", err
		}
	}
	for _, fold := range folders {
		_, err = cs.rPUT(directorySource + "/" + fold, current_folder_path)
		if err != nil {
			return "", errors.New("something wrong when recursion: " + directoryDestination + " " + err.Error())
		}
	}
	return "Everything seems oky", nil
}