package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"summer/controler"
	Log "summer/log"
	"summer/service"
	"summer/utils"

)

func main(){
	service.Init()
	service.Ini()

	app:=gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession",store))
	app.Use(Log.Log)
	//注册一个中间件函数 执行 其他请求时 会进入中间件 就是一个拦截操作
	user:=app.Group("/user",gin.Logger(),gin.Recovery())
	user.Use(controler.Authentication)

	user.POST("/upload",controler.Upload)

	user.GET("/download/:name",controler.Download)

	app.GET("scan",controler.Scan)

	app.GET("/share", func(c *gin.Context) {
		path := c.Query("name")
		url:="http://qr.topscan.com/api.php?text="+path
		c.Redirect(http.StatusMovedPermanently,url)
	})

	app.POST("/register", controler.Register)

	app.POST("/login",controler.Login)

	user.GET("/share",controler.EncryptedSharing)

	app.GET("/get/:name", func(context *gin.Context) {
		name:=context.Param("name")
		secret := context.Query("secret")

		owner := service.Get(secret+":"+name)

		if owner!=nil {
			hname := utils.ByteToString(owner.([]byte))
			context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
			context.Writer.Header().Add("Content-Type", "application/octet-stream")
			fmt.Println(name + hname)
			path:="./file/"+hname+"/"+name
			context.File(path)
		} else {
			context.JSON(200,gin.H{
				"message": name+"没有被分享",
			})
		}
	})

	user.POST("/transmission",controler.Transmission)

	app.GET("/down",controler.Down)

	app.Run(":8080")
}