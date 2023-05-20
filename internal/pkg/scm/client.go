package scm

import (
	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/gitea"
	"github.com/drone/go-scm/scm/driver/gitee"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/driver/gitlab"
	"github.com/drone/go-scm/scm/transport"
	scm2 "go-to-cloud/internal/models/scm"
	"net/http"
	"strings"
)

// newClient 获取scm客户端
func newClient(origin scm2.Type, isPublic bool, uri, token *string) (client *scm.Client, err error) {
	switch origin {
	case scm2.Github:
		client = github.NewDefault()
		break
	case scm2.Gitlab:
		client, err = gitlab.New(*uri)
		break
	case scm2.Gitee:
		client = gitee.NewDefault()
		break
	case scm2.Gitea:
		client, err = gitea.New(*uri)
	}

	if client != nil {
		client.Client = scmHttpClient(origin, isPublic, token)
	}
	return
}

func scmHttpClient(origin scm2.Type, isPublic bool, token *string) *http.Client {
	if isPublic || token == nil || len(strings.TrimSpace(*token)) == 0 {
		return &http.Client{}
	}

	switch origin {
	case scm2.Gitlab:
		return &http.Client{
			Transport: &transport.PrivateToken{Token: *token},
		}
	default:
		return &http.Client{
			Transport: &transport.BearerToken{Token: *token},
		}
	}
}
