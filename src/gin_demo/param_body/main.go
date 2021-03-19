package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.POST("/info", func(c *gin.Context) {
		bodyByte, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))

		uname := c.PostForm("name")
		age := c.DefaultPostForm("age", "18") // 此方法可以设置默认值

		c.String(http.StatusOK, "%s, %s, %s", string(bodyByte), uname, age)
	})
	r.Run(":8080")
}
// curl -X POST "http://127.0.0.1:8080/info" -d "name=giant&age=28"