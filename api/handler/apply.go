package handler

import (
	"k8s_CICD/api/service"
	"k8s_CICD/model"
	"k8s_CICD/model/kube_param/command_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Apply(c *gin.Context) {
	var parameter command_model.ApplyParameter
	if err := c.ShouldBindJSON(&parameter); err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: model.ParameterFail,
			Msg:  "【apply参数不匹配】" + err.Error(),
			Data: nil,
		})
		return
	}
	err := service.ApplyService(&parameter)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: model.ApplyFail,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code: model.Success,
		Msg:  "【Apply 成功】",
		Data: "",
	})
}
