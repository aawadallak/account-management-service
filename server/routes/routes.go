package routes

import (
	"latest/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		register := main.Group("register")
		{
			register.POST("/", controller.StoreNewUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controller.LoginWithUserAndPassword)
		}

		verify := main.Group("verify")
		{
			verify.POST("/", controller.VerifyUserCode)
			verify.POST("/resend", controller.ResendVerifyCode)

		}

		lostAccount := main.Group("lostaccount")
		{
			lostAccount.POST("/username")
			lostAccount.POST("/password", controller.LostPassword)
		}

		changePassword := main.Group("password")
		{
			changePassword.POST("/change", controller.ChangePassword)
		}

		changeEmail := main.Group("email")
		{
			changeEmail.POST("/change", controller.ChangeEmail)
			changeEmail.POST("/confirm", controller.ConfirmEmailChange)
		}

	}
	return router
}
