package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/builder"
	"gorm.io/datatypes"
	"os"
	"testing"
)

func TestGetBuildNodesById(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.NotNil(t, actualNode)

	assert.Equal(t, model.Name, actualNode.Name)
	assert.Equal(t, model.Remark, actualNode.Remark)
	assert.Equal(t, model.Workspace, actualNode.K8sWorkerSpace)
}

func TestDeleteBuilderNode(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.NotNil(t, actualNode)

	assert.NoError(t, DeleteBuilderNode(1, node))

	nilNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, nilNode.ID)
}

func TestDecryptKubeConfig(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	kubecfg := "abcdefg123"

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: kubecfg,
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.NotNil(t, actualNode)

	assert.NotEqual(t, kubecfg, actualNode.K8sKubeConfigEncrypted)
	assert.Empty(t, actualNode.k8sKubeConfigDecrypted)
	assert.Equal(t, kubecfg, *actualNode.DecryptKubeConfig())
	assert.NotEmpty(t, actualNode.k8sKubeConfigDecrypted)
	assert.Equal(t, actualNode.k8sKubeConfigDecrypted, *actualNode.DecryptKubeConfig())

}

func TestUpdateBuilderNode(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer func() {
			os.Unsetenv("UnitTestEnv")
		}()
		prepareDb()
	}

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	node, err := NewBuilderNode(model, 1, []uint{2, 3})
	assert.NoError(t, err)
	assert.NotNil(t, node)

	model.Id = node
	model.Name = "test2"
	model.MaxWorkers = 4
	err = UpdateBuilderNode(model, 1, []uint{3, 4})
	assert.NoError(t, err)

	actualNode, err := GetBuildNodesById(node)
	assert.NoError(t, err)
	assert.Equal(t, "test", actualNode.Name)  // not changed
	assert.Equal(t, 4, actualNode.MaxWorkers) // changed
	assert.Equal(t, datatypes.JSON("[3,4]"), actualNode.BelongsTo)
}

// testGetBuildNodesOnK8sByOrgId 需要sqlite 支持json_contains方法
// 需要考虑是否调整sql
func testGetBuildNodesOnK8sByOrgId(t *testing.T) {
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

	model := &builder.OnK8sModel{
		Name:       "test",
		MaxWorkers: 3,
		Workspace:  "ws",
		KubeConfig: "xcy",
		Orgs:       []uint{1},
		Remark:     "remark",
	}

	_, err = NewBuilderNode(model, 1, nil)
	assert.NoError(t, err)

	n, err := GetBuildNodesOnK8sByOrgId([]uint{1}, "", nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(n))
}

func TestMergeOrg(t *testing.T) {
	repos := make([]BuilderNodeWithOrg, 0)

	repos = append(repos, BuilderNodeWithOrg{
		BuilderNode: BuilderNode{
			Model: Model{ID: 1},
		},
		OrgLite: OrgLite{
			OrgId:   1,
			OrgName: "1",
		},
	})
	repos = append(repos, BuilderNodeWithOrg{
		BuilderNode: BuilderNode{
			Model: Model{ID: 1},
		},
		OrgLite: OrgLite{
			OrgId:   1,
			OrgName: "1",
		},
	})

	a, err := mergeBuilderNodeOrg(repos)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(a))

	repos = append(repos, BuilderNodeWithOrg{
		BuilderNode: BuilderNode{
			Model: Model{ID: 2},
		},
		OrgLite: OrgLite{
			OrgId:   2,
			OrgName: "2",
		},
	})

	a, err = mergeBuilderNodeOrg(repos)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(a))
}
