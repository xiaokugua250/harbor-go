// harbor golang客户端

package client

import (
	"crypto/tls"
	"github.com/bitly/go-simplejson"
	"net/http"
	"strconv"

	"errors"

	"conf"
)

var (
	harborClient  http.Client
	harborBaseURL string
	harborAdmin   string
	harborPasswd  string
)

func InitHarbor(config *conf.Config) error {
	harborClient = http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	harborBaseURL = config.Get("harbor.baseurl")

	if harborBaseURL == "" {
		return errors.New("必须设置harbor_baseURL")
	}

	harborAdmin = config.Get("harbor.admin")
	if harborAdmin == "" {
		return errors.New("必须设置harbor_admin")
	}
	harborPasswd = config.Get("harbor.password")
	if harborPasswd == "" {
		return errors.New("必须设置harbor.password")
	}
	harborBaseURL = "https://" + harborAdmin + ":" + harborPasswd + "@" + harborBaseURL
	return nil
}

// todo 添加search 参数
func HarborProjectSearch(params string) (*simplejson.Json, error) {
	// API URL  /search
	request_Search_url := harborBaseURL + "/api/search?q=" + params
	request_harbor_Search, err := http.NewRequest(http.MethodGet, request_Search_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_Search)
	if err != nil {
		return nil, err
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborConfigurations() (*simplejson.Json, error) {
	// API URL  /configurations
	request_Configurations_url := harborBaseURL + "/api/configurations"

	request_harbor_Configurations, err := http.NewRequest(http.MethodGet, request_Configurations_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_Configurations)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborStatistics() (*simplejson.Json, error) {
	// API URL  /statistics
	request_Statistics_url := harborBaseURL + "/api/statistics"

	request_harbor_Statistics, err := http.NewRequest(http.MethodGet, request_Statistics_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_Statistics)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

// todo harbor 1.5.x以上支持该方法
func HarborLables() (*simplejson.Json, error) {
	// API URL /labels
	request_Lables_url := harborBaseURL + "/api/labels"
	request_harbor_Lables, err := http.NewRequest(http.MethodGet, request_Lables_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_Lables)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

//
func HarborProjects() (*simplejson.Json, error) {
	// API URL /projects
	request_AllProjects_url := harborBaseURL + "/api/projects"
	request_harbor_AllProject, err := http.NewRequest(http.MethodGet, request_AllProjects_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_AllProject)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborSingleProject(project_id int) (*simplejson.Json, error) {
	// APU URL /projects/{project_id}
	request_singleProject_url := harborBaseURL + "/api/projects/" + strconv.Itoa(project_id)
	request_harbor_singleProject, err := http.NewRequest(http.MethodGet, request_singleProject_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_singleProject)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

//
func HarborProjectsRepositoriesTags(project, repository string, details_id int) (*simplejson.Json, error) {
	// API URL /api/repositories/{projects}/{repostory}/tags?detail=1

	request_ProjectRepositoryTags_url := harborBaseURL + "/api/repositories/" + project + "/" + repository + "/tags?details=" + strconv.Itoa(details_id)

	request_harbor_ProjectRepositoryTags_url, err := http.NewRequest(http.MethodGet, request_ProjectRepositoryTags_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_ProjectRepositoryTags_url)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborProjectsRepositoriesDeleteByTagName(project, repository, tagName string) (*simplejson.Json, error) {
	// API URL /api/repositories/{projects}/{repostory}/tags?detail=1

	request_ProjectRepositoryTags_url := harborBaseURL + "/api/repositories/" + project + "/" + repository + "/tags/" + tagName
	request_harbor_ProjectRepositoryTags_url, err := http.NewRequest(http.MethodDelete, request_ProjectRepositoryTags_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_ProjectRepositoryTags_url)
	if err != nil {
		return nil, err
	}

	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}

	responseJson := simplejson.New()
	responseJson.Set("delete", "success")
	return responseJson, nil
	// 	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborRepositories(page, page_size, project_id int) (*simplejson.Json, error) {
	// API URL /repositories

	request_AllRepositories_url := harborBaseURL + "/api/repositories?page=" + strconv.Itoa(page) + "&page_size=" + strconv.Itoa(page_size) + "&project_id=" + strconv.Itoa(project_id)

	request_harbor_AllRepositories, err := http.NewRequest(http.MethodGet, request_AllRepositories_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_AllRepositories)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborProjectMembers(project_id int) (*simplejson.Json, error) {
	// API URL /api/projects/{project_id}/members
	request_ProjectMembers_url := harborBaseURL + "/api/projects/" + strconv.Itoa(project_id) + "/members"
	request_harbor_ProjectMembers, err := http.NewRequest(http.MethodGet, request_ProjectMembers_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_ProjectMembers)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborCurrentUser() (*simplejson.Json, error) {
	// API URL /api/users/current
	request_CurrentUser_url := harborBaseURL + "/api/users/current"
	request_harbor_CurrentUser, err := http.NewRequest(http.MethodGet, request_CurrentUser_url, nil)
	if err != nil {
		return nil, err
	}
	harbor_Response, err := harborClient.Do(request_harbor_CurrentUser)
	if err != nil {
		return nil, err
	}
	if harbor_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(harbor_Response.StatusCode))
	}
	return simplejson.NewFromReader(harbor_Response.Body)
}

func HarborSystemInfo() (*simplejson.Json, error) {
	// API URL /systeminfo
	request_harbor_systeminfo_url := harborBaseURL + "/api/systeminfo"
	request_harbor_systeminfo, err := http.NewRequest(http.MethodGet, request_harbor_systeminfo_url, nil)
	if err != nil {
		return nil, err
	}
	systeminfo_Response, err := harborClient.Do(request_harbor_systeminfo)
	if err != nil {
		return nil, err
	}
	if systeminfo_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(systeminfo_Response.StatusCode))
	}
	return simplejson.NewFromReader(systeminfo_Response.Body)

}

func HarborSystemInfovolumes() (*simplejson.Json, error) {
	// API URL /systeminfo/volumes
	request_harbor_systeminfo_volumes_url := harborBaseURL + "/api/systeminfo/volumes"
	request_harbor_systeminfo_volumes, err := http.NewRequest(http.MethodGet, request_harbor_systeminfo_volumes_url, nil)
	if err != nil {
		return nil, err
	}
	systeminfo_Response, err := harborClient.Do(request_harbor_systeminfo_volumes)
	if err != nil {
		return nil, err
	}
	if systeminfo_Response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP返回结果有误，HTTP状态码为:" + strconv.Itoa(systeminfo_Response.StatusCode))
	}
	return simplejson.NewFromReader(systeminfo_Response.Body)

}
