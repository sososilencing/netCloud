package controler

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"summer/model"
	"summer/service"
)


// 后期 可能以时间 来命名文件 或者 他上传什么文件 以什么名字命名
func Upload(c *gin.Context){

	defer func() {
		if recover()!=nil{

			c.JSON(200,"已上传")
		}
	}()

	session := sessions.Default(c)
	name := session.Get("user")
	private := c.PostForm("private")

	fmt.Println(private)
	file,header,err := c.Request.FormFile("file")

	if err != nil {
		//这里应该是打日志
		fmt.Println(err.Error())
		c.JSON(500,gin.H{
			"status" :"Fail",
			"message" : "上传失败",
		})
	}
	if header.Size > 1024*1024*1024*1024 {
		c.JSON(200,gin.H{
			"status" : "Fail",
			"message" : "文件过大",
		})
		return
	}

	filename := header.Filename

	file1 := &model.File{
		Filename: filename,
		Owner:    name.(string),
	}

	var out *os.File
	if private=="1"{
		file1.Private = 1
		path1:=".//"+"file//"+name.(string)


		bl,err := PathExists(path1)

		if err != nil {
			fmt.Println(err.Error())
		}

		if !bl {
			os.Mkdir(path1,os.ModePerm)
		}
		out ,err = os.Create(path1+"//"+filename)

	}else {
		file1.Private = 2
		out, err = os.Create(".//" + "file//" + filename)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	defer out.Close()
	if service.InsertFile(file1) {

		write, err := io.Copy(out, file)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(write)
		c.JSON(200, gin.H{
			"stauts":  "OK",
			"message": "上传成功",
		})
	}else {
		c.JSON(200, gin.H{
			"stauts":  "Fail",
			"message": "已上传",
		})
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}