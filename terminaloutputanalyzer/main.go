package main

import (
	"strconv"
	"strings"
	"unicode"

	"terminaloutputanalyzer/src"
)

func main() {
	textFile, err := src.ReadLines("terminaloutputanalyzer.txt")
	// textFile, err := src.ReadLines("terminaloutputanalyzer2.txt")
	if err != nil {
		panic(err)
	}

	src.DirectoryPath = make(map[string]*src.Directory)
	src.SizeMap = make(map[string]int)
	src.CreateNewDir("/", nil)

	var currentDirectory *src.Directory
	currentDirectory = nil
	for _, v := range textFile {
		// fmt.Println("command: ", v)
		if strings.HasPrefix(v, "$ cd") {
			splitted := strings.Split(v, " ")
			if len(splitted) < 3 {
				panic("wrong amount of split")
			}
			directoryName := strings.Join(splitted[2:], " ")
			fullpath := src.GetFullpath(directoryName, currentDirectory)

			if directoryName == ".." {
				if currentDirectory.PreviousDirectory == nil {
					panic("cannot return to previous directory: " + currentDirectory.Name)
				} else {
					currentDirectory = currentDirectory.PreviousDirectory
					// fmt.Println("current directory: " + currentDirectory.fullpath)
					continue
				}
			}

			_, ok := src.DirectoryPath[fullpath]
			if !ok {
				panic("dir does not exist: " + fullpath)
			}
			currentDirectory = src.DirectoryPath[fullpath]
			// fmt.Println("current directory: " + currentDirectory.fullpath)
		} else if v == "$ ls" {
			// fmt.Println("ls")
		} else if unicode.IsDigit([]rune(v)[0]) {
			splitted := strings.Split(v, " ")
			if len(splitted) < 2 {
				panic("wrong format")
			}
			sizeString := splitted[0]
			size, err := strconv.Atoi(sizeString)
			if err != nil {
				panic(err)
			}
			name := strings.Join(splitted[1:], " ")
			// fmt.Println("get file", size, name)

			newFile := &src.File{
				Name: name,
				Size: size,
			}
			currentDirectory.Files = append(currentDirectory.Files, newFile)
		} else if strings.HasPrefix(v, "dir ") {
			splitted := strings.Split(v, " ")
			if len(splitted) < 2 {
				panic("wrong format")
			}
			name := strings.Join(splitted[1:], " ")
			fullpath := src.GetFullpath(name, currentDirectory)
			// fmt.Println("get dir", size, name)

			src.CreateNewDir(name, currentDirectory)

			currentDirectory.Directories = append(currentDirectory.Directories, src.DirectoryPath[fullpath])
		} else {
			panic("wrong format")
		}
	}

	src.ViewFolder()

	src.AnswerNo1()
	src.AnswerNo2()
}
