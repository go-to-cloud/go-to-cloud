package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/models/deploy/k8s"
	"gorm.io/datatypes"
	"time"
)

// K8sRepo K8s部署环境
type K8sRepo struct {
	Model
	Name          string         `json:"name" gorm:"column:name;type:nvarchar(20)"`
	KubeConfig    string         `json:"kubeconfig" gorm:"column:kube_config;type:text"`
	ServerVersion string         `json:"serverVersion" gorm:"column:server_ver;type:varchar(10)"`
	CreatedBy     uint           `json:"createdBy" gorm:"column:created_by"`  // 仓库创建人
	BelongsTo     datatypes.JSON `json:"belongsTo" gorm:"column:belongs_to;"` // 所属组织
	Remark        string         `json:"remark" gorm:"column:remark;type:nvarchar(500)"`
}

func (m *K8sRepo) TableName() string {
	return "k8s_repo"
}

type MergedK8sRepoWithOrg struct {
	K8sRepo
	Org []OrgLite
}

type K8sRepoWithOrg struct {
	K8sRepo
	OrgLite
}

func QueryK8sRepoById(id uint) (*K8sRepo, error) {
	var repo K8sRepo
	tx := conf.GetDbClient().Model(&repo)

	tx.First(&repo, id)

	return returnWithError(&repo, tx.Error)
}

func QueryK8sRepo(orgs []uint, repoNamePattern string, pager *models.Pager) ([]MergedK8sRepoWithOrg, error) {
	var repo []K8sRepoWithOrg

	tx := conf.GetDbClient().Model(&K8sRepo{})

	tx = tx.Select("k8s_repo.*, org.Id AS orgId, org.Name AS orgName")
	tx = tx.Joins("INNER JOIN org ON JSON_CONTAINS(k8s_repo.belongs_to, cast(org.id as JSON), '$')")
	tx = tx.Where("org.ID IN ? AND org.deleted_at IS NULL", orgs)

	if len(repoNamePattern) > 0 {
		tx = tx.Where("k8s_repo.name like ?", repoNamePattern+"%")
	}

	if pager != nil && pager.PageSize > 0 && pager.PageIndex > 1 {
		tx = tx.Limit(pager.PageSize).Offset((pager.PageIndex - 1) * pager.PageSize)
	}

	tx = tx.Order("created_at desc")

	err := tx.Scan(&repo).Error

	if err == nil {
		return mergeK8sRepoOrg(repo)
	} else {
		return nil, err
	}
}

func mergeK8sRepoOrg(repos []K8sRepoWithOrg) ([]MergedK8sRepoWithOrg, error) {
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

	merged := make(map[uint]*MergedK8sRepoWithOrg)
	for _, repo := range repos {
		if merged[repo.ID] == nil {
			merged[repo.ID] = &MergedK8sRepoWithOrg{
				K8sRepo: K8sRepo{
					Model: Model{
						ID:        repo.ID,
						CreatedAt: repo.CreatedAt,
						UpdatedAt: repo.UpdatedAt,
						DeletedAt: repo.DeletedAt,
					},
					Name:          repo.Name,
					KubeConfig:    repo.KubeConfig,
					CreatedBy:     repo.CreatedBy,
					BelongsTo:     datatypes.JSON{},
					Remark:        repo.Remark,
					ServerVersion: repo.ServerVersion,
				},
				Org: r[repo.ID],
			}
		}
	}
	rlt := make([]MergedK8sRepoWithOrg, len(merged))
	counter := 0
	for _, m := range merged {
		rlt[counter] = *m
		counter++
	}
	return rlt, nil
}

func buildK8sRepo(model *k8s.K8s, userId uint, orgs []uint, gormModel *Model) (*K8sRepo, error) {

	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	repo := K8sRepo{
		Model:         *gormModel,
		Name:          model.Name,
		CreatedBy:     userId,
		BelongsTo:     datatypes.JSON(belongs),
		ServerVersion: model.ServerVersion,
		Remark:        model.Remark,
		KubeConfig:    *model.KubeConfig,
	}

	return &repo, nil
}

// BindK8sRepo 绑定K8s仓库
func BindK8sRepo(model *k8s.K8s, userId uint, orgs []uint) error {
	g := &Model{
		CreatedAt: time.Now(),
	}
	repo, err := buildK8sRepo(model, userId, orgs, g)
	if err != nil {
		return err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("updated_at").Create(&repo).Error
	return err
}

// UpdateK8sRepo 更新K8s仓库
func UpdateK8sRepo(model *k8s.K8s, userId uint, orgs []uint) error {
	g := &Model{
		UpdatedAt: time.Now(),
	}

	repo, err := buildK8sRepo(model, userId, orgs, g)
	if err != nil {
		return err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("created_at", "created_by", "kube_config").Where("id = ?", model.Id).Updates(&repo).Error
	return err
}

func DeleteK8sRepo(userId, repoId uint) error {

	tx := conf.GetDbClient()

	// TODO: 校验当前userId是否拥有数据删除权限

	err := tx.Delete(&K8sRepo{
		Model: Model{
			ID: repoId,
		},
	}).Error

	return err
}
