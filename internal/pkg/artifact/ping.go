package artifact

import (
	"go-to-cloud/internal/models/artifact"
	pkgRegistry "go-to-cloud/internal/pkg/artifact/registry"
)

func Ping(testing *artifact.Testing) (bool, error) {

	hub, err := pkgRegistry.GetRegistryHub(testing.IsSecurity, &testing.Url, &testing.User, &testing.Password)
	if err == nil && hub != nil {
		return true, hub.Ping()
	} else {
		return false, err
	}
}
