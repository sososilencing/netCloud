package controler

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"summer/model"
	"summer/service"
	"summer/utils"
)

//加密分享 首先 第一点这个文件 必须是你自己的,公私有 我倒是无所谓
//一个加密口令
// 加密过后的url 需要输入密码才能下载
//把秘钥和文件名字 存在redis 里面当做key  然后会得到一个 分享者的名字 和 文件名 再转发到Download
func EncryptedSharing(c *gin.Context)  {

	filename := c.Query("filename")

	secret := c.Query("secret")
	fmt.Println(secret+"1")
	if secret == ""{
		secret = utils.RandSeq(4)
	}
	fmt.Println(secret+"2")
	fmt.Println(filename)
	session := sessions.Default(c)
	name := session.Get("user")
	hname :=name.(string)
	file := &model.File{
		Filename: filename,
		Owner:    hname,
	}
	if service.FindFile(file){
		if file.Private ==1 {
			service.Set(secret+":"+filename, hname)
		}else {
			service.Set(secret+":"+filename,"")
		}
		c.JSON(200,gin.H{
			"status":"OK",
			"secret":secret,
			"Download": c.Request.Host+"/"+"get/"+filename,
		})
	}else {
		c.JSON(200,gin.H{
			"messgae" :"你加密的文件分享不存在",
		})
	}
}

