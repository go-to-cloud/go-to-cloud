package project

import (
	"github.com/patrickmn/go-cache"
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/repositories"
	"strconv"
	"strings"
	"time"
)

// projectArtifacts 内存缓存（如果使用多副本，可能会导致前端出现数据不一致的情况）
var projectArtifacts *cache.Cache

// ListArtifactsByKeywords 根据制品名称查询
func ListArtifactsByKeywords(projectId uint, queryString *string) ([]struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}, error) {
	var matching []artifact.FullName
	if c, _ := projectArtifacts.Get(strconv.Itoa(int(projectId))); c != nil && len(c.([]artifact.FullName)) > 0 {
		matching = c.([]artifact.FullName)
	} else {
		if repo, err := repositories.QueryLatestImagesByProjectId(projectId); err != nil {
			return nil, err
		} else {
			matching = make([]artifact.FullName, len(repo))
			for i, image := range repo {
				matching[i] = artifact.FullName{Id: image.Id, Name: image.Name, Tag: image.Tag, Address: image.Address}
			}
			projectArtifacts.Set(strconv.Itoa(int(projectId)), matching, 0)
		}
	}

	matched := make([]struct {
		Name string `json:"name"`
		Id   uint   `json:"id"`
	}, 0)
	for _, a := range matching {
		if strings.Contains(a.Name, *queryString) {
			matched = append(matched, struct {
				Name string `json:"name"`
				Id   uint   `json:"id"`
			}{Name: a.Name, Id: a.Id})
		}
	}
	return matched, nil
}

func init() {
	projectArtifacts = cache.New(time.Minute*3, 0)
}
