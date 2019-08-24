package controler

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type file struct {
	Filename string `json:"filename"`
	Size int64 `json:"size"`
}
func Scan(c *gin.Context) {

	session := sessions.Default(c)
	name := session.Get("user")

	var filess []file
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"status": "Fail",
		})
	}

	for _, fi := range files {
		var fil file
		if !fi.IsDir() {
			fil = file{
				Filename: fi.Name(),
				Size:     fi.Size(),
			}
			filess = append(filess, fil)
		} else if name!=nil && fi.Name() == name.(string) {
			ff, err := ioutil.ReadDir(path + fi.Name())
			if err != nil {
				fmt.Println(err)
			}
			for _, f := range ff {
				if !f.IsDir() {
					fil = file{
						Filename: f.Name(),
						Size:     f.Size(),
					}
				}
				filess = append(filess, fil)
			}
		}
	}
	c.JSON(200, filess)
}
