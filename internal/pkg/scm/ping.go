package scm

import (
	"context"
	"errors"
	"github.com/drone/go-scm/scm"
	scmModels "go-to-cloud/internal/models/scm"
	"net/http"
	"strconv"
)

// Ping 测试SCM服务是否可用
func Ping(testing *scmModels.Testing) (bool, error) {
	if client, err := newClient(testing.Origin, testing.IsPublic, &testing.Url, testing.Token); err != nil {
		return false, err
	} else {
		_, resp, err := client.Repositories.List(context.Background(), scm.ListOptions{Page: 0, Size: 1})

		if err != nil || resp == nil {
			return false, err
		}

		if http.StatusOK <= resp.Status && resp.Status < http.StatusBadRequest {
			return true, nil
		}

		if resp.Status == http.StatusUnauthorized || resp.Status == http.StatusForbidden {
			return false, errors.New("connected, but token is not allowed")
		}

		return false, errors.New("connection failed, returns: " + strconv.Itoa(resp.Status))
	}
}
