package router

import (
	"k8s_CICD/api/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/")
	api.Use()
	{
		api.POST("rollout", handler.Rollout)
		api.POST("apply", handler.Apply)
	}

	return r
}
