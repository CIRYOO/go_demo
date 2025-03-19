package handlers

import (
	db "awesomeProject/config"
	"awesomeProject/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUser(c *gin.Context) {
	// 查询用户是否存在
	name := c.Query("name")
	var user models.User
	result := db.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
		}
		return
	}
}
