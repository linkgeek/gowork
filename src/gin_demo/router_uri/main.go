package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id": c.Param("id"),
		})
	})
	r.Run()
}
// curl -X GET "http://127.0.0.1:8080/giant/88"