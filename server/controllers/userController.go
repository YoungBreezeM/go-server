package controllers

import (
	"server/log"
	"server/models"
	"server/services"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint true "用户ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUserInfo(c *gin.Context) {
	// 处理获取用户信息的逻辑
	userID := c.Param("id")

	// 根据userID获取用户信息
	user, err := services.GetUserInfo(userID)
	log.Log.Debug("sdss")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get user information",
		})
		return
	}
	time.Sleep(time.Second * 2)
	c.JSON(200, user)
}

func UpdateUserInfo(c *gin.Context) {
	// 处理更新用户信息的逻辑
	userID := c.Param("id")

	// 解析请求体中的用户信息
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	// 更新用户信息
	if err := services.UpdateUserInfo(userID, &user); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update user information",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User information updated successfully",
	})
}
