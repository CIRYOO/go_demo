package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("hello world")
	r := gin.Default()
	config.InitDB()
	// 启动服务
	routes.SetupRouter(r)
	_ = r.Run(":9090")
}
