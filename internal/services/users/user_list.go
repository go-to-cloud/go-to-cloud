package users

import (
	"encoding/json"
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func GetUserList() ([]user.User, error) {
	if users, err := repositories.GetAllUser(); err != nil {
		return nil, err
	} else {
		rlt := make([]user.User, len(users))
		for i, u := range users {
			rlt[i] = user.User{
				Id:         u.ID,
				CreatedAt:  utils.JsonTime(u.CreatedAt),
				RealName:   u.RealName,
				Account:    u.Account,
				Pinyin:     u.Pinyin,
				PinyinInit: u.PinyinInit,
				Email:      u.Email,
				Mobile:     u.Mobile,
				BelongsTo: func() []string {
					orgs := make([]string, len(u.Orgs))
					for i2, o := range u.Orgs {
						orgs[i2] = o.Name
					}
					return orgs
				}(),
				Kind: func() []string {
					var m []string
					json.Unmarshal(u.Kind, &m)
					return m
				}(),
			}
		}
		return rlt, nil
	}
}

func GetUsersByOrg(orgId uint) ([]user.User, error) {
	if us, err := repositories.GetUsersByOrg(orgId); err != nil {
		return nil, err
	} else {
		rlt := make([]user.User, len(us))
		for i, u := range us {
			rlt[i] = user.User{
				Id:        u.ID,
				CreatedAt: utils.JsonTime(u.CreatedAt),
				RealName:  u.RealName,
				Account:   u.Account,
			}
		}
		return rlt, nil
	}
}

func GetUserBelongs(userId uint) ([]user.Org, error) {
	if us, err := repositories.GetOrgsByUser(userId); err != nil {
		return nil, err
	} else {
		rlt := make([]user.Org, len(us))
		for i, o := range us {
			rlt[i] = user.Org{
				Id:        o.ID,
				CreatedAt: o.CreatedAt.Format("2006-01-02 15:04:05"),
				Name:      o.Name,
			}
		}
		return rlt, nil
	}
}
