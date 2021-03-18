package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "router generic")
	})
	r.Run()
}
// curl -X GET "http://127.0.0.1:8080/giant/88"