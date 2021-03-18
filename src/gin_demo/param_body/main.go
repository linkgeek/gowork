package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.POST("/info", func(c *gin.Context) {
		uname := c.PostForm("uname")
		age := c.DefaultPostForm("age", "18") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "200",
			"uname": uname,
			"age":    age,
		})
	})
	r.Run(":8080")
}
// curl -X POST "http://127.0.0.1:8080/info" -d "{"name": "giant"}"