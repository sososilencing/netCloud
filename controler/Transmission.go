package controler

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"summer/service"
	"summer/utils"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

//也是通过redis 来负责
func  Transmission(c *gin.Context)  {

	session := sessions.Default(c)
	name := session.Get("user")

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

	path4:=".//" + "tmp//"
	bl,err := PathExists(path4)

	if err != nil {
		fmt.Println(err.Error())
	}

	if !bl {
		os.Mkdir(path4,os.ModePerm)
	}
	path1 := path4 +filename
	out, err := os.Create(path1)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer out.Close()

	write, err := io.Copy(out, file)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(write)
	code :=name.(string)+":"+filename+utils.RandSeq(10)
	fmt.Println(code)

	decode := base64Encode([]byte(code))

	service.SetOne(decode,path1)

	c.JSON(200, gin.H{
		"stauts":  "OK",
		"message":c.Request.Host+"/"+"down?filename="+decode,
	})
}

func Down(c *gin.Context)  {
	filename := c.Query("filename")

	name ,err := base64Decode(filename)

	if err != nil {
		fmt.Println(err.Error())
	}

	path1:=service.GetOne(filename)

	if path1 == nil{
		c.JSON(200,gin.H{
			"message":"下载失败",
		})
	}else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
		c.Writer.Header().Add("Content-Type", "application/octet-stream")

		path2 := utils.ByteToString(path1.([]byte))
		c.File(path2)

		os.Remove(path2)
	}
}

func base64Encode(src []byte) string {
	return coder.EncodeToString(src)
}

func base64Decode(src string) ([]byte, error) {
	return coder.DecodeString(src)
}
