package controller

import (
	"latest/dto"
	"latest/services"

	"github.com/gin-gonic/gin"
)

func StoreNewUser(c *gin.Context) {

	var user dto.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": "Can'\t bind JSON",
		})
		return
	}

	s := services.UserService{}

	_, err = s.Store(user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(201)
}

func LoginWithUserAndPassword(c *gin.Context) {

	var user dto.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": "Can'\t bind JSON",
		})
		return
	}

	s := services.UserService{}

	token, err := s.LoginWithUserAndPassword(user.Username, user.Password)

	if err != nil {
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func VerifyUserCode(c *gin.Context) {

	code := c.Query("code")

	email := c.Query("email")

	s := services.UserService{}

	err := s.VerifyCode(code, email)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.Status(204)
}

func ResendVerifyCode(c *gin.Context) {

	email := c.Query("email")

	s := services.UserService{}

	err := s.ResendVerifyCode(email)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.Status(204)

}
