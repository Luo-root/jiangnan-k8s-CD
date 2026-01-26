package router

import (
	"k8s_CICD/api/handler"
	"k8s_CICD/util/key/verify"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/")
	api.Use(verify.Verify())
	{
		api.POST("rollout", handler.Rollout)
		api.POST("apply", handler.Apply)
	}

	return r
}
