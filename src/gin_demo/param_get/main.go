package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.GET("/info", func(c *gin.Context) {
		uname := c.Query("uname")
		age := c.DefaultQuery("age", "18")
		c.String(200, "%s, %s", uname, age)
	})
	r.Run()
}
// curl -X GET "http://127.0.0.1:8080/info?uname=giant"