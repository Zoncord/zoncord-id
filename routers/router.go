package routers

import (
	"github.com/Zoncord/zoncord-id/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		signin := v1.Group("/signin")
		{
			signin.GET("", handlers.MethodNotAllowed)
			signin.POST("", handlers.PostSignIn)
			signin.PUT("", handlers.MethodNotAllowed)
			signin.DELETE("", handlers.MethodNotAllowed)
			signin.PATCH("", handlers.MethodNotAllowed)
		}

		signup := v1.Group("/signup")
		{
			signup.POST("", handlers.PostSignUp)
			signup.GET("", handlers.MethodNotAllowed)
			signup.PUT("", handlers.MethodNotAllowed)
			signup.DELETE("", handlers.MethodNotAllowed)
			signup.PATCH("", handlers.MethodNotAllowed)
		}

		user := v1.Group("/user")
		{
			user.GET("", handlers.GetCurrentUserData)
			user.POST("", handlers.MethodNotAllowed)
			user.PUT("", handlers.MethodNotAllowed)
			user.DELETE("", handlers.MethodNotAllowed)
			user.PATCH("", handlers.MethodNotAllowed)
		}
	}
	return r
}
