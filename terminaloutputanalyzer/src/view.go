package src

import "fmt"

func ViewFolder() {
	fmt.Println("================ START STRUCTURE ================")
	viewSubFolder(DirectoryPath["/"], 0)
	fmt.Println("================ END STRUCTURE ================")
}

func viewSubFolder(dict *Directory, level int) {
	printTab(level)
	fmt.Printf("%s %d", dict.Name, dict.GetSize())
	// if dict.previousDirectory != nil {
	// 	fmt.Printf("	prevDir:%s", dict.previousDirectory.name)
	// }
	// fmt.Printf("	fullPath:%s", dict.fullpath)
	fmt.Printf("\n")
	for _, v := range dict.Directories {
		viewSubFolder(v, level+1)
	}
	for _, v := range dict.Files {
		viewFile(v, level+1, dict.Fullpath)
	}
}
func viewFile(file *File, level int, fullpath string) {
	printTab(level)
	fmt.Printf("%s %d", file.Name, file.Size)
	// fmt.Printf("	fullPath:%s%s", fullpath, file.name)
	fmt.Printf("\n")
}

func printTab(level int) {
	for i := 0; i < level*4; i++ {
		fmt.Printf(" ")
	}
}

func AnswerNo1() {
	// fmt.Println("Find the sum of all directories whose size is at most 100000")

	var sum int
	for _, v := range DirectoryPath {
		size := v.GetSize()
		// fmt.Printf("size of %s = %d\n", key, size)
		if size <= 100000 {
			sum += size
		}
	}

	fmt.Printf("1. %d\n", sum)
}
func AnswerNo2() {
	// fmt.Println("The disk size of the drive described by the below input is 70000000. You need to delete a directory such that the free space available in the disk is at least 30000000. Find the size of the smallest possible directory that can be deleted to make at least 30000000 bytes available as free space.")

	wholeDiskSize := DirectoryPath["/"].GetSize()
	var chosen = -1

	for _, v := range DirectoryPath {
		if v.Fullpath == "/" {
			continue
		}
		dirSize := v.GetSize()
		if wholeDiskSize-dirSize <= 30000000 {
			if chosen == -1 || dirSize < chosen {
				chosen = dirSize
			}
		}
	}

	fmt.Printf("2. %d\n", chosen)
}
