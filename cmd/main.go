package main

import (
	"k8s_CICD/router"
	"k8s_CICD/util/file"
	"k8s_CICD/util/key/generation"
)

func main() {
	r := router.InitRouter()

	res, err := generation.GenerateRandomString(32)
	if err != nil {
		panic(err.Error())
	}

	err = file.WriteFile(file.KeyPath, res)
	if err != nil {
		panic(err.Error())
	}

	err = r.Run(":8080")
	if err != nil {
		panic("【启动失败】" + err.Error())
	}
}
