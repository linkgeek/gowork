package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.icon", "./favicon.icon")
	r.Run()
}

// go build -o router_static && ./router_static
// win: go build -o router_static.exe && router_static.exe