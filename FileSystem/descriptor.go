package FileSystem

type FileDescriptor struct {
	FileType string
	Nlink    int
	Size     int
	Blocks   []int
	IsOpen   bool
}
