package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name string `form:"name"`
	Addr string `form:"addr"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main()  {
	r:=gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind err: %v", err)
	}
}

// curl -X GET "http://127.0.0.1:8080/testing?name=giant&addr=shenzhen&birthday=1991-02-01"
// curl -X POST "http://127.0.0.1:8080/testing" -d "name=giant&addr=shenzhen"
// curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/testing" -d "{\"name\":\"giant\", \"addr\":\"28\", \"birthday\":\"2020-03-19\"}"