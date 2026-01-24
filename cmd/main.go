package main

import (
	"k8s_CICD/router"
)

func main() {
	r := router.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		panic("【启动失败】" + err.Error())
	}
}
