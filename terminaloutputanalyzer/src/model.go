package src

type Directory struct {
	Name              string // ex: "documents"
	Fullpath          string // ex: "users/global/documents"
	PreviousDirectory *Directory
	Directories       []*Directory
	Files             []*File
}
type File struct {
	Name string
	Size int
}

func (r *Directory) GetSize() int {
	size, ok := SizeMap[r.Fullpath]
	if !ok {
		var sum int
		for _, v := range r.Directories {
			sum += v.GetSize()
		}
		for _, v := range r.Files {
			sum += v.Size
		}

		SizeMap[r.Fullpath] = sum
		size = sum
	}
	return size
}
