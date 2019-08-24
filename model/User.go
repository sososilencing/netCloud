package model

type User struct {
	Id int
	Level int
	Name string
	Passwd string
	UploadNum int
}

type File struct {
	Id int
	Filename string
	Owner string
	Private int8
}
