package Log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Log(c *gin.Context)  {
	path :="./log/message.log"
	logfile,err:=os.OpenFile(path,os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Wrong")
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(c.Request.URL)
	if c.Errors!=nil {
		log.Println(c.Errors)
	}
}
