package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"summer/model"
)

var db *gorm.DB
func Init()  {
	db , _ = gorm.Open("mysql","root:iiroxi@tcp(127.0.0.1)/netdisc?charset=utf8")
}

func Insert(user *model.User)  bool{
	db.Where("name=?",user.Name).Find(&user)

	if user.Id==0{
		db.Create(user)
		return true
	}
	fmt.Println(user)
	return false
}

func Login(user *model.User)  bool{


	db.Where(&model.User{
		Name:      user.Name,
		Passwd:    user.Passwd,
	}).Find(&user)

	if user.Id==0{
		return false
	}

	return true
}

func InsertFile(file *model.File)  bool{

	db.Where("filename = ? AND private = ? AND owner=?",file.Filename,file.Private,file.Owner).Find(&file)
	f
	if file.Id==0{
		db.Create(file)
		return true
	}
	return false
}

func FindFile(file *model.File)  bool{

	db.Where("filename = ? AND private = ?", file.Filename, 2).Find(&file)

	db.Where("filename = ? AND owner = ?",file.Filename,file.Owner).Find(&file)

	if file.Id==0 {
		return false
	}
	return true
}