package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name" binding:"required"`
	Age int `form:"age" binding:"required,gt=10"`
	Addr string `form:"addr" binding:"required"`
}

func main()  {
	r:=gin.Default()
	r.GET("/testing", testing)
	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(404, "%v", err)
		c.Abort()
		return
	}
	c.String(200, "%v", person)
}

// curl -X GET "http://127.0.0.1:8080/testing?name=giant&addr=shenzhen&birthday=1991-02-01"