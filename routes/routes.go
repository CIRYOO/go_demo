package routes

import (
	"awesomeProject/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// 用户路由组
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/login", handlers.LoginUser)
	}
}
