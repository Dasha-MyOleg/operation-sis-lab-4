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

	c.Unlink("file1")
	c.Ls()
}
