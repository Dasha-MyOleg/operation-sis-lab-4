package core

import (
	"fmt"
	"operation-sis-lab-4/FileSystem"
)

type Core struct {
	fs                  *FileSystem.FileSystem
	openFileDescriptors []*FileSystem.OpenFileDescriptor
	blockSize           int
}

// ????????????? ???????? ???????
func (c *Core) Mkfs(descriptorsCount int) {
	fmt.Println("System initialization...")
	c.fs = &FileSystem.FileSystem{}
	c.fs.Mkfs()
	c.openFileDescriptors = make([]*FileSystem.OpenFileDescriptor, descriptorsCount)
	c.blockSize = FileSystem.BlockSize
	fmt.Println("System is ready to work!")
}

// ????????? ?????? ?????
func (c *Core) Create(fileName string) {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return
	}
	if c.fs.Find(fileName) {
		fmt.Println("Error: File", fileName, "already exists")
		return
	}
	c.fs.Create(fileName)
}

// ????????? ?????? ??????
func (c *Core) Ls() {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return
	}
	c.fs.Ls()
}

// ????????? ?????????? ??? ????
func (c *Core) Stat(fileName string) {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return
	}
	if c.fs.Find(fileName) {
		c.fs.Stat(fileName)
	} else {
		fmt.Println("Error: File", fileName, "does not exist")
	}
}

// ????????? ?????
func (c *Core) Unlink(fileName string) {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return
	}
	if !c.fs.Find(fileName) {
		fmt.Println("Error: File", fileName, "does not exist")
		return
	}
	c.fs.Unlink(fileName)
	fmt.Println("File", fileName, "deleted successfully")
}

// ????????? ????????? ?????????
func (c *Core) Link(linkWith, toLink string) {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return
	}
	if c.fs.Find(toLink) {
		fmt.Println("Error: File", toLink, "already exists")
		return
	}
	if !c.fs.Find(linkWith) {
		fmt.Println("Error: File", linkWith, "does not exist")
		return
	}
	c.fs.Link(linkWith, toLink)
	fmt.Println("Link created:", toLink, "->", linkWith)
}

// ????????? ????? (open name)
func (c *Core) Open(fileName string) *FileSystem.OpenFileDescriptor {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return nil
	}
	if !c.fs.Find(fileName) {
		fmt.Println("Error: File", fileName, "does not exist")
		return nil
	}
	index := c.findFreeIndex()
	if index == -1 {
		fmt.Println("No free file descriptor available")
		return nil
	}

	descriptor := &FileSystem.OpenFileDescriptor{
		Id:     index,
		Desc:   c.fs.GetDescriptor(fileName),
		Offset: 0,
	}
	c.openFileDescriptors[index] = descriptor
	descriptor.Desc.IsOpen = true
	fmt.Println("File", fileName, "opened with descriptor ID:", index)
	return descriptor
}

// ???????? ????? (close fd)
func (c *Core) Close(fd *FileSystem.OpenFileDescriptor) {
	if fd == nil {
		fmt.Println("Error: Closing non-existing file")
		return
	}
	fmt.Println("Closing file descriptor", fd.Id)
	c.openFileDescriptors[fd.Id] = nil
	fd.Desc.IsOpen = false
}

// ????? ???????? ??????? ??? ????????? ?????
func (c *Core) findFreeIndex() int {
	for i, v := range c.openFileDescriptors {
		if v == nil {
			return i
		}
	}
	return -1
}

// ??????? ? ????? (read fd size)
func (c *Core) Read(fd *FileSystem.OpenFileDescriptor, size int) {
	if fd == nil {
		fmt.Println("Error: Invalid file descriptor")
		return
	}
	if size <= 0 {
		fmt.Println("Error: Read size must be greater than zero")
		return
	}
	if fd.Offset+size > fd.Desc.Size {
		fmt.Println("Error: Read size exceeds file size")
		return
	}

	fmt.Println("Reading", size, "bytes from file descriptor", fd.Id)
	fd.Offset += size // ???????? ???????? ????? ???????
}

// ????? ? ???? (write fd size)
func (c *Core) Write(fd *FileSystem.OpenFileDescriptor, data []byte) {
	if fd == nil {
		fmt.Println("Error: Invalid file descriptor")
		return
	}
	totalSize := len(data)

	if fd.Offset+totalSize > fd.Desc.Size {
		fmt.Println("Error: Write exceeds file size, consider truncating the file first")
		return
	}

	fmt.Println("Writing data to file descriptor", fd.Id)
	fd.Offset += totalSize // ????????? ???????? ????? ??????
}

// ???????????? ???????? ? ????? (seek fd offset)
func (c *Core) Seek(fd *FileSystem.OpenFileDescriptor, offset int) {
	if fd == nil {
		fmt.Println("Error: File descriptor is not valid")
		return
	}
	if offset < 0 {
		fmt.Println("Error: Offset cannot be negative")
		return
	}
	if offset > fd.Desc.Size {
		fmt.Println("Error: Offset exceeds file size")
		return
	}
	fd.Offset = offset
	fmt.Println("Seek operation set to offset", offset)
}

// ????? ??????? ????? (truncate name size)
func (c *Core) Truncate(fileName string, size int) {
	if c.fs == nil {
		fmt.Println("Error: File system is not initialized")
		return
	}
	if !c.fs.Find(fileName) {
		fmt.Println("Error: File", fileName, "does not exist")
		return
	}
	descriptor := c.fs.GetDescriptor(fileName)
	if size < descriptor.Size {
		fmt.Println("Truncating file:", fileName, "to new size:", size)
	} else {
		fmt.Println("Extending file:", fileName, "to new size:", size)
	}
	descriptor.Size = size
}
