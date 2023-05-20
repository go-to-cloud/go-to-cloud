package registry

import (
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	url2 "net/url"
	"strings"
)

func GetRegistryHub(isSecurity bool, url, user, password *string) (hub *registry.Registry, err error) {
	schema := "https"
	if !isSecurity {
		schema = "http"
	}

	registryUrl := fmt.Sprintf("%s://%s", schema, strings.TrimSuffix(*url, "/"))
	var u *url2.URL
	u, err = url2.Parse(registryUrl)
	if err != nil {
		return
	}

	hub, err = registry.New(fmt.Sprintf("%s://%s", schema, u.Host), *user, *password)

	return
}
