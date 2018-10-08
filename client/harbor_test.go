package client

import (
	"fmt"
	"harbor-go/conf"

	"testing"
)

func TestHarborSystemInfo(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborSystemInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}

func TestHarborSystemInfovolumes(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborSystemInfovolumes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}
func TestHarborStatistics(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborStatistics()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}

func TestHarborProjects(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborProjects()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}

func TestHarborProjectSearch(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborProjectSearch("params")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}

func TestHarborProjectsRepositoriesTags(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborProjectsRepositoriesTags("project", "respository", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}

func TestHarborProjectMembers(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborProjectMembers(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}

func TestHarborRepositories(t *testing.T) {
	// 从配置文件加载配置信息
	config := new(conf.Config)
	err := config.LoadHarborConfig("harbor", "../")
	if err != nil {
		fmt.Println(err)
	}
	err = InitHarbor(config)
	if err != nil {
		fmt.Println(err)
	}

	harborJson, err := HarborRepositories(1, 15, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(harborJson)
}
