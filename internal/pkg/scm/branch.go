package scm

import (
	"context"
	"fmt"
	scm2 "github.com/drone/go-scm/scm"
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
	"net/url"
	"strings"
)

// ListBranches 列出代码分支
func ListBranches(projectId, sourceCodeId uint) ([]scm.Branch, error) {
	sourceCode, err := repositories.GetProjectSourceCodeById(projectId, sourceCodeId)
	if err != nil {
		return nil, err
	}
	if client, err := newClient(scm.Type(sourceCode.CodeRepo.ScmOrigin), false, &sourceCode.CodeRepo.Url, &sourceCode.CodeRepo.AccessToken); err != nil {
		return nil, err
	} else {
		var repo, prefix string
		if sourceCode.CodeRepo.ScmOrigin == int(scm.Github) || sourceCode.CodeRepo.ScmOrigin == int(scm.Gitee) {
			if u, err := url.Parse(sourceCode.GitUrl); err != nil {
				prefix = "https://github.com"
			} else {
				prefix = fmt.Sprintf("%s://%s/", u.Scheme, u.Host)
			}
		} else {
			prefix = "/"
		}

		repo = strings.TrimPrefix(
			strings.TrimPrefix(
				strings.TrimSuffix(sourceCode.GitUrl, ".git"),
				sourceCode.CodeRepo.Url,
			),
			prefix)
		branches, _, err := client.Git.ListBranches(
			context.Background(),
			repo,
			scm2.ListOptions{
				Page: 0,
				Size: 100,
			})

		rlt := make([]scm.Branch, len(branches))
		for i, branch := range branches {
			rlt[i] = scm.Branch{
				Reference: branch,
			}
		}
		if err != nil {
			return nil, err
		}

		return rlt, nil
	}
}
