package main

import (
	core "operation-sis-lab-4/Core"
)

func main() {
	c := &core.Core{}
	c.Mkfs(10)

	c.Create("file1")
	c.Create("file2")
	c.Ls()

	fd1 := c.Open("file1")
	c.Write(fd1, []byte("Hello World"))
	c.Seek(fd1, 5)
	c.Read(fd1, 5)
	c.Close(fd1)

	c.Truncate("file1", 50)
	c.Ls()

	c.Unlink("file1")
	c.Ls()
}
