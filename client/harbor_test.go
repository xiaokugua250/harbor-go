package client

import (
	"fmt"
	"nebula/common/conf"
	"strings"
	"testing"
)

func TestHarborHttpClient(t *testing.T) {
	// 	HarborHttpClient("https://10.127.48.30", nil, "/api/systeminfo")
	service := "harbor"
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig(service)
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	/*systemInfo,err :=HarborSystemInfo()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(systemInfo)*/
	harborJson, err := HarborSystemInfo()
	if err != nil {
		fmt.Println("err happend ===>", err)
	}
	fmt.Println(harborJson)
}

func TestImageNameSpiliet(t *testing.T) {
	imageFullName := "nscc-gz.cn/nscc/spark:2.1.0"

	projects := strings.Split(imageFullName, "/")
	repostory := strings.Split(strings.Split(imageFullName, "/")[2], ":")[1]
	fmt.Println(projects[1])
	fmt.Println(repostory)
}
