package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.POST("/info", func(c *gin.Context) {
		uname := c.PostForm("uname")
		age := c.DefaultPostForm("age", "18")
		c.String(200, "%s, %s", uname, age)
	})
	r.Run()
}
// curl -X POST "http://127.0.0.1:8080/info" -d "{uname: "giant", age: 18}"