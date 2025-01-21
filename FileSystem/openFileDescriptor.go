package FileSystem

type OpenFileDescriptor struct {
	Id     int
	Desc   *FileDescriptor
	Offset int
}
