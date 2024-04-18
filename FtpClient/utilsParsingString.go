package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func get_files_folders_local(path string ) ([]string, []string, error) {
	entries, err := os.ReadDir(path)
	var folders = make([]string, 0)
	var files = make([]string, 0)
	if err != nil {
		return folders, files, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		} else {
			files = append(files, entry.Name())
		}
	}
	return folders, files, nil
}

func get_filename_path(path string) (string) {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func SplitString(s string, c byte) ([]string){
	res := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			start := i
			j := i+1
			for ; j < len(s); j++ {
				if s[j] == c {
					break
				}
			}
			res = append(res, s[start : j])
			i = j
		}
	}
	return res
}

func parse_get_connection_ftp(input string) (string, string) {
	// Parse the input string to get the IP and Port
	// Example: "227 Entering Passive Mode (127,0,0,1,195,149)"
	// The port is calculated as (195*256+149)
	// Output: "127.0.0.1", "195*256+149".
	split1 := strings.Split(input, "(")
	split2 := strings.Split(split1[1], ",")
	split3 := strings.Split(split2[5], ")")
	ip := split2[0] + "." + split2[1] + "." + split2[2] + "." + split2[3]
	first_part_port, _ := strconv.ParseInt(split2[4], 10, 32)
	second_part_port, _ := strconv.ParseInt(split3[0], 10, 32)

	port := strconv.FormatInt((first_part_port*256 + second_part_port), 10)
	return ip, port
}

func get_files_folders_current(cs *FtpSession, path string) ([]string, []string, error) {
	response1, err := cs.LS(path)
	if err != nil {
		return []string{}, []string{}, errors.New("something is wrong with LS " +  path + " " + err.Error())
	}
	response2, err := cs.NLST(path)
	if err != nil {
		return []string{}, []string{}, errors.New("something is wrong with NLST " +  path + " " + err.Error())
	}
	archives1 := strings.Split(response1, "\n")
	archives2 := strings.Split(response2, "\n")
	if len(archives1) != len(archives2) {
		return []string{}, []string{}, errors.New("LS and NLST didn't return the same " + path)
	}
	
	archivesFiltered1 := []string{}
	archivesFiltered2 := []string{}
	for index, arch := range archives1 {
		if len(arch) > 0 {
			archivesFiltered1 = append(archivesFiltered1, arch)
			archivesFiltered2 = append(archivesFiltered2, archives2[index])
		}
	}
	folders := make([]string, 0)
	files := make([]string, 0)
	archives1 = archivesFiltered1
	archives2 = archivesFiltered2
	if len(archives1) > 0 {
		for index, arch := range archives1 {
			parts := strings.Split(archivesFiltered2[index], "/")
			filename := parts[len(parts)-1]
			filename = filename[:len(filename)-1]
			if arch[0] == 'd' {
				folders = append(folders, filename)				
			} else {
				files = append(files, filename)
			}
		}
	}
	return folders, files, nil
}