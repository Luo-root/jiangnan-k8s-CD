package verify

import (
	"k8s_CICD/model"
	"k8s_CICD/util/file"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Authorization")

		k, err := file.ReadFile(file.KeyPath)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code: model.ReadFail,
				Msg:  "【密钥读取异常】" + err.Error(),
				Data: nil,
			})
			// 当在中间件或处理器中调用 c.Abort() 后，Gin 会停止执行当前请求后续的所有中间件和处理器函数
			c.Abort()
			return
		}
		if key != k {
			c.JSON(http.StatusOK, model.Response{
				Code: model.KeyFail,
				Msg:  "【密钥校验失败】",
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
