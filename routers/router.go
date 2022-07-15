package routers

import (
	"ZoncordID/handlers"
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
		}
	}
	return r
}
