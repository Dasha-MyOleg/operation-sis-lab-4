package FileSystem

import (
	"fmt"
)

type FileSystem struct {
	Directory map[string]*FileDescriptor
}

func (fs *FileSystem) Mkfs() {
	fmt.Println("Creation of file system...")
	fs.Directory = make(map[string]*FileDescriptor)
	fmt.Println("File system created")
}

func (fs *FileSystem) Create(fileName string) {
	descriptor := &FileDescriptor{
		FileType: "reg",
		Nlink:    1,
		Size:     0,
		IsOpen:   true,
		Blocks:   []int{},
	}
	fs.Directory[fileName] = descriptor
	fmt.Println("Create file:", fileName, "| Size:", descriptor.Size)
}

func (fs *FileSystem) Ls() {
	fmt.Println("Hard links of current directory:")
	for f, d := range fs.Directory {
		fmt.Println("Name:", f, "\t Size:", d.Size)
	}
}

func (fs *FileSystem) Stat(fileName string) {
	descriptor := fs.Directory[fileName]
	fmt.Println("Type:", descriptor.FileType,
		"\tHard links count:", descriptor.Nlink,
		"\tSize:", descriptor.Size,
		"\tBlocks count:", len(descriptor.Blocks))
}

func (fs *FileSystem) Link(linkWith, toLink string) {
	descriptor := fs.Directory[linkWith]
	descriptor.Nlink++
	fs.Directory[toLink] = descriptor
	fmt.Println("Create hard link", toLink, "with", linkWith)
}

func (fs *FileSystem) Unlink(fileName string) {
	fmt.Println("Delete file:", fileName)
	descriptor := fs.Directory[fileName]
	descriptor.Nlink--
	if descriptor.Nlink == 0 {
		delete(fs.Directory, fileName)
	}
}

func (fs *FileSystem) Find(fileName string) bool {
	_, exists := fs.Directory[fileName]
	return exists
}

func (fs *FileSystem) GetDescriptor(fileName string) *FileDescriptor {
	return fs.Directory[fileName]
}
