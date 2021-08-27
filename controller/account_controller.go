package controller

import (
	"latest/dto"
	"latest/services"

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {

	var user dto.UserUpdatePassword

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": "Can'\t bind JSON",
		})
		return
	}

	s := services.UserService{}

	err = s.ChangePassword(user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ChangeEmail(c *gin.Context) {

	var user dto.UserUpdateEmail

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": "Can'\t bind JSON",
		})
		return
	}
	s := services.UserService{}

	err = s.ChangeEmail(user)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.AbortWithStatus(200)
}

func ConfirmEmailChange(c *gin.Context) {

	oldEmail := c.Query("oldEmail")
	newEmail := c.Query("newEmail")
	code := c.Query("code")

	s := services.UserService{}

	err := s.ConfirmEmailChange(newEmail, oldEmail, code)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatus(200)

}

func LostPassword(c *gin.Context) {

	email := c.Query("email")

	svc := services.UserService{}

	err := svc.LostPassword(email)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}
