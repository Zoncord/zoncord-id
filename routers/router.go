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
			signin.GET("", handlers.DefaultSignIn)
			signin.POST("", handlers.PostSignIn)
			signin.PUT("", handlers.DefaultSignIn)
			signin.DELETE("", handlers.DefaultSignIn)
			signin.PATCH("", handlers.DefaultSignIn)
		}

		signup := v1.Group("/signup")
		{
			signup.POST("", handlers.PostSignUp)
			signup.GET("", handlers.DefaultSignUp)
			signup.PUT("", handlers.DefaultSignUp)
			signup.DELETE("", handlers.DefaultSignUp)
			signup.PATCH("", handlers.DefaultSignUp)
		}
	}
	return r
}
