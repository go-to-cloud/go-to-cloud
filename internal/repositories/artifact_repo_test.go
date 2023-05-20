package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/artifact"
	"os"
	"testing"
)

func TestMergeArtifactRepoOrg(t *testing.T) {
	repos := make([]ArtifactRepoWithOrg, 0)

	repos = append(repos, ArtifactRepoWithOrg{
		ArtifactRepo: ArtifactRepo{
			Model: Model{ID: 1},
		},
		OrgLite: OrgLite{
			OrgId:   1,
			OrgName: "1",
		},
	})
	repos = append(repos, ArtifactRepoWithOrg{
		ArtifactRepo: ArtifactRepo{
			Model: Model{ID: 1},
		},
		OrgLite: OrgLite{
			OrgId:   1,
			OrgName: "1",
		},
	})

	merged, err := mergeArtifactRepoOrg(repos)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(merged))

	repos = append(repos, ArtifactRepoWithOrg{
		ArtifactRepo: ArtifactRepo{
			Model: Model{ID: 2},
		},
		OrgLite: OrgLite{
			OrgId:   2,
			OrgName: "1",
		},
	})
	merged, err = mergeArtifactRepoOrg(repos)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(merged))
}

func TestBindK8sRepo(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	err := CreateOrg(&orgName, &orgRemark)
	assert.NoError(t, err)

	err = BindArtifactRepo(&artifact.Artifact{
		Name:   orgName,
		Remark: orgRemark,
	}, 1, []uint{1})
	assert.NoError(t, err)

	err = UpdateArtifactRepo(&artifact.Artifact{
		Testing: artifact.Testing{Id: 1},
		Name:    "abc",
	}, 1, []uint{1})
	assert.NoError(t, err)

	err = DeleteArtifactRepo(1, 1)

	// Query depends JSON_CONTAINS
	//a, err := QueryArtifactRepo([]uint{1}, "")
	//assert.NoError(t, err)
	//assert.Equal(t, 1, len(a))
	//assert.Equal(t, orgName, a[0].Name)
	//assert.Equal(t, orgRemark, a[0].Remark)
}

func TestGetArtifactRepoByID(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	err := CreateOrg(&orgName, &orgRemark)
	assert.NoError(t, err)

	err = BindArtifactRepo(&artifact.Artifact{
		Testing: artifact.Testing{
			Type:       artifact.Maven,
			IsSecurity: true,
			Url:        "a",
			User:       "b",
			Password:   "c",
		},
		Name:   orgName,
		Remark: orgRemark,
	}, 1, []uint{1})
	assert.NoError(t, err)

	url, account, password, isSecurity, origin, err := GetArtifactRepoByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "a", *url)
	assert.Equal(t, "b", *account)
	assert.Equal(t, "c", *password)
	assert.Equal(t, true, isSecurity)
	assert.EqualValues(t, artifact.Maven, origin)
}
