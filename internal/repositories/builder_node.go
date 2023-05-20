package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/utils"
	"gorm.io/datatypes"
	"time"
)

type BuilderNode struct {
	Model
	BelongsTo              datatypes.JSON `json:"belongs_to" gorm:"column:belongs_to"`                                       // 所属机构
	Name                   string         `json:"name" gorm:"column:name;type:nvarchar(50)"`                                 // 节点名称
	NodeType               int            `json:"node_type" gorm:"column:node_type"`                                         // 节点类型；1：k8s；2：docker；3: windows；4：linux；5：macos
	MaxWorkers             int            `json:"max_workers" gorm:"column:max_workers"`                                     // 同时执行任务数量；0:不限；其他值：同时构建任务上限
	K8sWorkerSpace         string         `json:"k8s_worker_space" gorm:"column:k8s_worker_space;type:nvarchar(50)"`         // k8s名字空间
	K8sKubeConfigEncrypted string         `json:"k8s_kubeconfig_encrypted" gorm:"column:k8s_kubeconfig_encrypted;type:text"` // 已加密kubeconfig
	k8sKubeConfigDecrypted string         `gorm:"-"`
	Remark                 string         `json:"remark" gorm:"column:remark;type:nvarchar(200)"`
	AgentVersion           string         `json:"agent_version" gorm:"column:agent_version;type:nvarchar(200)"`
}

func (m *BuilderNode) TableName() string {
	return "builder_nodes"
}

type BuilderNodeWithOrg struct {
	BuilderNode
	OrgLite
}

type MergedBuilderNodeWithOrg struct {
	BuilderNode
	Org []OrgLite
}

func GetBuildNodesById(id uint) (*BuilderNode, error) {
	db := conf.GetDbClient()

	tx := db.Model(&BuilderNode{})

	var agent BuilderNode
	tx = tx.Where("id = ?", id)
	tx = tx.First(&agent)

	return returnWithError(&agent, tx.Error)
}

func GetBuildNodesOnK8sByOrgId(orgs []uint, repoNamePattern string, pager *models.Pager) ([]MergedBuilderNodeWithOrg, error) {
	var repo []BuilderNodeWithOrg

	tx := conf.GetDbClient().Model(&BuilderNode{})

	tx = tx.Select("builder_nodes.*, org.Id AS orgId, org.Name AS orgName")
	tx = tx.Joins("INNER JOIN org ON JSON_CONTAINS(builder_nodes.belongs_to, cast(org.id as JSON), '$')")
	tx = tx.Where("builder_nodes.node_type = ? AND org.deleted_at IS NULL", builder.K8s)

	if orgs != nil && len(orgs) > 0 {
		tx = tx.Where("org.ID IN ?", orgs)
	}

	if len(repoNamePattern) > 0 {
		tx = tx.Where("builder_nodes.name like ?", repoNamePattern+"%")
	}

	if pager != nil && pager.PageSize > 0 && pager.PageIndex > 1 {
		tx = tx.Limit(pager.PageSize).Offset((pager.PageIndex - 1) * pager.PageSize)
	}

	tx = tx.Order("created_at desc")

	err := tx.Scan(&repo).Error

	if err == nil {
		return mergeBuilderNodeOrg(repo)
	} else {
		return nil, err
	}
}

func mergeBuilderNodeOrg(repos []BuilderNodeWithOrg) ([]MergedBuilderNodeWithOrg, error) {
	r := make(map[uint][]OrgLite)
	for _, repo := range repos {
		x := r[repo.ID]
		if x == nil {
			r[repo.ID] = make([]OrgLite, 0)
		}
		r[repo.ID] = append(r[repo.ID], OrgLite{
			OrgId:   repo.OrgId,
			OrgName: repo.OrgName,
		})
	}

	merged := make(map[uint]*MergedBuilderNodeWithOrg)
	for _, repo := range repos {
		if merged[repo.ID] == nil {
			merged[repo.ID] = &MergedBuilderNodeWithOrg{
				BuilderNode: BuilderNode{
					Model: Model{
						ID:        repo.ID,
						CreatedAt: repo.CreatedAt,
						UpdatedAt: repo.UpdatedAt,
						DeletedAt: repo.DeletedAt,
					},
					Name:                   repo.Name,
					BelongsTo:              datatypes.JSON{},
					MaxWorkers:             repo.MaxWorkers,
					AgentVersion:           repo.AgentVersion,
					K8sWorkerSpace:         repo.K8sWorkerSpace,
					Remark:                 repo.Remark,
					NodeType:               repo.NodeType,
					K8sKubeConfigEncrypted: repo.K8sKubeConfigEncrypted,
				},
				Org: r[repo.ID],
			}
		}
	}
	rlt := make([]MergedBuilderNodeWithOrg, len(merged))
	counter := 0
	for _, m := range merged {
		rlt[counter] = *m
		counter++
	}
	return rlt, nil
}

func (m *BuilderNode) EncryptKubeConfig() {
	m.K8sKubeConfigEncrypted = utils.Base64AesEny(m.k8sKubeConfigDecrypted)
}
func (m *BuilderNode) DecryptKubeConfig() *string {
	m.k8sKubeConfigDecrypted = utils.Base64AesEnyDecode(m.K8sKubeConfigEncrypted)
	return &m.k8sKubeConfigDecrypted
}

func buildRepoModel(model *builder.OnK8sModel, _ uint, orgs []uint, gormModel *Model) (*BuilderNode, error) {
	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	repo := BuilderNode{
		Model:                  *gormModel,
		BelongsTo:              datatypes.JSON(belongs),
		Name:                   model.Name,
		NodeType:               int(builder.K8s),
		MaxWorkers:             model.MaxWorkers,
		K8sWorkerSpace:         model.Workspace,
		k8sKubeConfigDecrypted: model.KubeConfig,
		Remark:                 model.Remark,
	}
	repo.EncryptKubeConfig()

	return &repo, nil
}

func NewBuilderNode(node *builder.OnK8sModel, userId uint, orgs []uint) (uint, error) {
	g := &Model{
		CreatedAt: time.Now(),
	}
	repo, err := buildRepoModel(node, userId, orgs, g)
	if err != nil {
		return 0, err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("updated_at").Create(&repo).Error
	if err != nil {
		return 0, err
	} else {
		return repo.ID, nil
	}
}

func DeleteBuilderNode(userId, nodeId uint) error {

	tx := conf.GetDbClient()

	// TODO: 校验当前userId是否拥有数据删除权限

	err := tx.Delete(&BuilderNode{
		Model: Model{
			ID: nodeId,
		},
	}).Error

	return err
}

// UpdateBuilderNode 更新构建节点信息
func UpdateBuilderNode(model *builder.OnK8sModel, userId uint, orgs []uint) error {
	g := &Model{
		UpdatedAt: time.Now(),
	}

	repo, err := buildBuilderNode(model, userId, orgs, g)
	if err != nil {
		return err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("created_at", "created_by", "k8s_kubeconfig_encrypted").Where("id = ?", model.Id).Updates(&repo).Error
	return err
}

func buildBuilderNode(model *builder.OnK8sModel, userId uint, orgs []uint, gormModel *Model) (*BuilderNode, error) {

	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	repo := BuilderNode{
		Model:      *gormModel,
		BelongsTo:  datatypes.JSON(belongs),
		MaxWorkers: model.MaxWorkers,
		Remark:     model.Remark,
	}

	return &repo, nil
}
