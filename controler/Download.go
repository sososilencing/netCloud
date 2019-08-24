package controler

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
	"summer/model"
	"summer/service"
	"time"
)

const path = "./file/"

func Download(c *gin.Context) {

	session := sessions.Default(c)
	name1 := session.Get("user")

	name, _ := c.Params.Get("name")

	fmt.Println(name)
	fmt.Println(path + name)

	file := &model.File{
		Filename: name,
		Owner:    name1.(string),
	}

	if !service.FindFile(file) {
		c.JSON(200, gin.H{
			"status":  "Fail",
			"message": "文件不存在",
		})
		return
	} else {
		var pathA string

		if file.Private == 1 {
			pathA = path + file.Owner + "/" + name
		} else {
			pathA = path + name
		}

		fi, _ := os.Stat(pathA)

		size := int(fi.Size()) //文件长度
		fmt.Println(size)

		c.Writer.Header().Add("cache-control","public")
		//c.Writer.Header().Add("Range",fmt.Sprintf("bytes %d-%d","0","0"))
		c.Writer.Header().Add("Content-Range", "bytes 0-"+strconv.Itoa(size-1)+"/"+strconv.Itoa(size))
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
		c.Writer.Header().Add("Content-Type", "application/octet-stream")

		fmt.Println(c.Writer.Header().Get("StatusCode"))
		fp, err := os.Open(pathA)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer fp.Close()

		//start := 0

		buf := make([]byte, 64)

		//c.Writer.Header().Set("Accept-Ranges","bytes="+strconv.Itoa(size-1)+"-")
		//fmt.Println(c.Writer.Header().Get("Accept-Ranges"))
		//之前这里是一下子读完 现在我们分片读 分片 write 出去 这样就可以记录 大概 7
		for {
			_ , err := fp.Read(buf)

			//start+=10
			//end = size-1-start

			if err != nil {
				fmt.Println(err.Error())
			}

			c.Writer.Write(buf)
			time.Sleep(time.Millisecond)
			if err != nil {
				if err == io.EOF {
					err = nil
					break
				} else {
					return
				}
			}
		}
	}
}