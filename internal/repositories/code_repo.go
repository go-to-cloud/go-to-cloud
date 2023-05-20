package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/models/scm"
	"gorm.io/datatypes"
	"time"
)

type CodeRepo struct {
	Model
	Name        string         `json:"name" gorm:"column:name;type:varchar(200)"`
	ScmOrigin   int            `json:"scmOrigin" gorm:"column:scm_origin"`                       // 代码仓库来源；gitlab(0);github(1);gitee(2);gitea(3)
	IsPublic    int8           `json:"isPublic" gorm:"column:is_public"`                         // 是否公开仓库
	AccessToken string         `json:"accessToken" gorm:"column:access_token;type:varchar(200)"` // 访问令牌 PAT
	Url         string         `json:"url" gorm:"column:url;type:varchar(500)"`                  // SCM平台地址（非项目仓库地址）
	CreatedBy   uint           `json:"createdBy" gorm:"column:created_by"`                       // 仓库创建人
	BelongsTo   datatypes.JSON `json:"belongsTo" gorm:"column:belongs_to;type:text"`             // SCM所属组织
	Remark      string         `json:"remark" gorm:"column:remark;type:varchar(200)"`
}

func (m *CodeRepo) TableName() string {
	return "code_repo"
}

type CodeRepoWithOrg struct {
	CodeRepo
	OrgLite
}

type MergedCodeRepoWithOrg struct {
	CodeRepo
	Org []OrgLite
}

func QueryCodeRepo(orgs []uint, repoNamePattern string, pager *models.Pager, imported bool) ([]MergedCodeRepoWithOrg, error) {
	var repo []CodeRepoWithOrg

	tx := conf.GetDbClient().Model(&CodeRepo{})

	tx = tx.Select("code_repo.*, org.Id AS orgId, org.Name AS orgName")
	tx = tx.Joins("INNER JOIN org ON JSON_CONTAINS(code_repo.belongs_to, cast(org.id as JSON), '$')")
	if imported {
		tx = tx.Joins("INNER JOIN project_source_code psc on code_repo.id = psc.code_repo_id and psc.deleted_at is null")
	}
	tx = tx.Where("org.ID IN ? AND org.deleted_at IS NULL", orgs)

	if len(repoNamePattern) > 0 {
		tx = tx.Where("code_repo.name like ?", repoNamePattern+"%")
	}

	if pager != nil && pager.PageSize > 0 && pager.PageIndex > 1 {
		tx = tx.Limit(pager.PageSize).Offset((pager.PageIndex - 1) * pager.PageSize)
	}

	tx = tx.Order("created_at desc")

	err := tx.Scan(&repo).Error

	if err == nil {
		return mergeCodeRepoOrg(repo)
	} else {
		return nil, err
	}
}

func mergeCodeRepoOrg(repos []CodeRepoWithOrg) ([]MergedCodeRepoWithOrg, error) {
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

	merged := make(map[uint]*MergedCodeRepoWithOrg)
	for _, repo := range repos {
		if merged[repo.ID] == nil {
			merged[repo.ID] = &MergedCodeRepoWithOrg{
				CodeRepo: CodeRepo{
					Model: Model{
						ID:        repo.ID,
						CreatedAt: repo.CreatedAt,
						UpdatedAt: repo.UpdatedAt,
						DeletedAt: repo.DeletedAt,
					},
					Name:        repo.Name,
					ScmOrigin:   repo.ScmOrigin,
					IsPublic:    repo.IsPublic,
					AccessToken: repo.AccessToken,
					Url:         repo.Url,
					CreatedBy:   repo.CreatedBy,
					BelongsTo:   datatypes.JSON{},
					Remark:      repo.Remark,
				},
				Org: r[repo.ID],
			}
		}
	}
	rlt := make([]MergedCodeRepoWithOrg, len(merged))
	counter := 0
	for _, m := range merged {
		rlt[counter] = *m
		counter++
	}
	return rlt, nil
}

func buildCodeRepo(model *scm.Scm, userId uint, orgs []uint, gormModel *Model) (*CodeRepo, error) {
	isPublic := int8(0)
	if model.IsPublic {
		isPublic = 1
	} else {
		isPublic = 0
	}
	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	repo := CodeRepo{
		Model:       *gormModel,
		ScmOrigin:   int(model.Origin),
		Name:        model.Name,
		IsPublic:    isPublic,
		AccessToken: *model.Token,
		Url:         model.Url,
		CreatedBy:   userId,
		BelongsTo:   datatypes.JSON(belongs),
		Remark:      model.Remark,
	}

	return &repo, nil
}

// BindCodeRepo 绑定代码仓库
func BindCodeRepo(model *scm.Scm, userId uint, orgs []uint) error {
	g := &Model{
		CreatedAt: time.Now(),
	}
	repo, err := buildCodeRepo(model, userId, orgs, g)
	if err != nil {
		return err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("updated_at").Create(&repo).Error
	return err
}

// UpdateCodeRepo 更新代码仓库
func UpdateCodeRepo(model *scm.Scm, userId uint, orgs []uint) error {
	g := &Model{
		UpdatedAt: time.Now(),
	}

	repo, err := buildCodeRepo(model, userId, orgs, g)
	if err != nil {
		return err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("created_at", "created_by").Where("id = ?", model.Id).Updates(&repo).Error
	return err
}

func DeleteCodeRepo(userId, repoId uint) error {

	tx := conf.GetDbClient()

	// TODO: 校验当前userId是否拥有数据删除权限

	err := tx.Delete(&CodeRepo{
		Model: Model{
			ID: repoId,
		},
	}).Error

	return err
}
