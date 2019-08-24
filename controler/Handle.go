package controler

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"summer/model"
	"summer/service"
)

func Authentication(c *gin.Context){
	session := sessions.Default(c)
	name := session.Get("user")
	if name ==nil{
		c.JSON(200,gin.H{
			"message": "请登录",
		})
		c.Abort()
		return
	}else {
		c.Next()
	}
}

func Login(c *gin.Context){
	session := sessions.Default(c)
	name1 := session.Get("user")
	if name1 !=nil{
		c.JSON(200,"已登录")
		return
	}else {
		name := c.Query("name")
		password := c.Query("passwd")
		if name == "" || password == "" {
			c.JSON(200, gin.H{
				"message": "请输入账号或密码",
			})
			return
		}
		user := &model.User{
			Name:   name,
			Passwd: password,
		}
		if !service.Login(user) {
			c.JSON(200, gin.H{
				"message": "请输入正确的账号或密码",
			})
			return
		}
		session.Set("user", user.Name)
		session.Save()
		c.JSON(200, gin.H{
			"message": "登录成功",
		})
	}
}

func Register(c *gin.Context)  {
	name := c.Query("name")
	password := c.Query("passwd")
	if name =="" || password == ""{
		c.JSON(200,gin.H{
			"message":"请输入合适的参数",
		})
		return
	}
	user := &model.User{
		Level:  1,
		Name:   name,
		Passwd: password,
		UploadNum: 0,
	}
	if !service.Insert(user){
		c.JSON(500,gin.H{
			"status":"Fail",
		})
		return
	}
	c.JSON(200,user)
}
