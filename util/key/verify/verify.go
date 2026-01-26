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
				Code: model.ParameterFail,
				Msg:  "【密钥读取异常】" + err.Error(),
				Data: nil,
			})
			return
		}
		if key != k {
			c.JSON(http.StatusOK, model.Response{
				Code: model.ParameterFail,
				Msg:  "【密钥校验失败】",
				Data: nil,
			})
			return
		}
	}
}
