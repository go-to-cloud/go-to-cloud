package artifact

import (
	"github.com/stretchr/testify/assert"
	pkgRegistry "go-to-cloud/internal/pkg/artifact/registry"
	"testing"
)
import artifactModel "go-to-cloud/internal/models/artifact"

func TestPing_Should_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	//ok, err := Ping(&artifactModel.Testing{
	//	IsSecurity: false,
	//	Kaniko:        "81.68.216.88:8080",
	//	User:       "admin",
	//	Password:   "some password",
	//})
	// nexus
	tester := &artifactModel.Testing{
		IsSecurity: false,
		Url:        "localhost:8081/repository/gotocloud",
		User:       "user",
		Password:   "pwd",
	}
	ok, err := Ping(tester)
	assert.True(t, ok)

	hub, err := pkgRegistry.GetRegistryHub(tester.IsSecurity, &tester.Url, &tester.User, &tester.Password)
	assert.NoError(t, err)

	str, err := hub.Repositories()
	assert.NotEmpty(t, str)
	assert.NoError(t, err)
}
