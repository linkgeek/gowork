package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.POST("/info", func(c *gin.Context) {
		uname := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")
		c.String(200, "%s, %s", uname, age)
	})
	r.Run()
}
// curl -X POST "http://127.0.0.1:8080/info" -d "name=giant&age=28"