package users

import (
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func GetOrgList() ([]user.Org, error) {
	if orgs, err := repositories.GetOrgs(); err != nil {
		return nil, err
	} else {
		rlt := make([]user.Org, len(orgs))
		for i, org := range orgs {
			rlt[i] = user.Org{
				Id:          org.ID,
				CreatedAt:   org.CreatedAt.Format("2006-01-02 15:04:05"),
				Name:        org.Name,
				MemberCount: uint(len(org.Users)),
				Remark:      org.Remark,
			}
		}
		return rlt, nil
	}
}

// JoinOrg 加入组织，如果当前成员在users中不存在，则移除
func JoinOrg(orgId uint, users []uint) error {
	if us, err := repositories.GetUsersByOrg(orgId); err != nil {
		return err
	} else {
		old := make([]uint, len(us))
		for i, u := range us {
			old[i] = u.ID
		}

		oldSet := utils.New(old...)
		newSet := utils.New(users...)

		delSet := utils.Minus(oldSet, newSet)      // 差集：移除
		comSet := utils.Complement(oldSet, newSet) // 补集：追加

		return repositories.UpdateMembersToOrg(orgId, utils.List(comSet), utils.List(delSet))
	}
}

// JoinOrgs 加入组织，如果当前成员在users中不存在，则移除
func JoinOrgs(orgId []uint, userId uint) error {
	if orgs, err := repositories.GetOrgsByUser(userId); err != nil {
		return err
	} else {
		old := make([]uint, len(orgs))
		for i, u := range orgs {
			old[i] = u.ID
		}

		oldSet := utils.New(old...)
		newSet := utils.New(orgId...)

		delSet := utils.Minus(oldSet, newSet)      // 差集：移除
		comSet := utils.Complement(oldSet, newSet) // 补集：追加

		return repositories.UpdateOrgsToMember(userId, utils.List(comSet), utils.List(delSet))
	}
}

func ResetPassword(userId uint, oldPassword, newPassword *string, force bool) (*string, error) {

	if force {
		pwd := utils.StrongPasswordGen(12)
		return pwd, repositories.ResetPassword(userId, pwd)
	} else {
		return newPassword, repositories.ResetPasswordWithCheckOldPassword(userId, oldPassword, newPassword)
	}
}
