package main

import (
	"fmt"
	"k8s_CICD/router"
	"k8s_CICD/util/display/figlet"
	"k8s_CICD/util/file"
	"k8s_CICD/util/key/generation"
	"strings"
)

func main() {
	figlet.Logo()

	r := router.InitRouter()

	key, err := generation.GenerateRandomString(32)
	if err != nil {
		panic(err.Error())
	}

	err = file.WriteFile(file.KeyPath, key)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(figlet.ColorSize("\n"+strings.Repeat("=", 100)+"\n", figlet.FgHiGreen))
	fmt.Println(figlet.ColorSize("【密钥key只出现一次请注意保留】: "+key, figlet.FgHiWhite))
	fmt.Println(figlet.ColorSize("\n"+strings.Repeat("=", 100), figlet.FgHiGreen))
	err = r.Run(":8080")
	if err != nil {
		panic("【启动失败】" + err.Error())
	}
}
