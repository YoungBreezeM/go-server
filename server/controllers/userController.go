package controllers

import (
	"server/constant"
	"server/db"
	"server/log"
	"server/models"
	"server/services"

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
func GetCurrentUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	u, err := services.GetUserInfoByToken(token)
	//
	if err != nil {
		log.Log.Errorln(err)
		c.JSON(500, models.R[string]{
			Status:  0,
			Message: err.Error(),
		})
		return
	}
	//
	c.JSON(200, models.R[db.User]{
		Status:  0,
		Data:    *u,
		Message: constant.UNAITHORIZED,
	})

}
