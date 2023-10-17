package src

func CreateNewDir(directoryName string, currentDirectory *Directory) {
	var fullpath = GetFullpath(directoryName, currentDirectory)
	_, ok := DirectoryPath[fullpath]
	if !ok {

		var previousDir *Directory
		if currentDirectory != nil {
			previousDir = currentDirectory
		}
		// fmt.Println("curDir:", directoryName, "prevDir:", previousDir)
		// fmt.Println("fullpath:", fullpath)
		DirectoryPath[fullpath] = &Directory{
			Name:              directoryName,
			Fullpath:          fullpath,
			PreviousDirectory: previousDir,
			Directories:       []*Directory{},
			Files:             []*File{},
		}
	}
}

func GetFullpath(directoryName string, currentDirectory *Directory) string {
	var fullpath = directoryName
	if currentDirectory != nil {
		fullpath = currentDirectory.Fullpath + fullpath + "/"
	}

	return fullpath
}
