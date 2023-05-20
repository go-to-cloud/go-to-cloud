package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/gorm"
)

// GetUsersByOrg 获取指定组织下的用户
// orgId：所属组织
func GetUsersByOrg(orgId uint) ([]*User, error) {
	db := conf.GetDbClient()

	var org Org
	err := db.Preload("Users").Where([]uint{orgId}).First(&org).Error

	return org.Users, err
}

func GetOrgsByUser(userId uint) ([]*Org, error) {
	db := conf.GetDbClient()

	var user User
	err := db.Preload("Orgs").Where([]uint{userId}).First(&user).Error

	return user.Orgs, err
}

func UpdateMembersToOrg(orgId uint, add, del []uint) error {
	db := conf.GetDbClient()

	return db.Transaction(func(tx *gorm.DB) (err error) {
		org := Org{Model: Model{ID: orgId}}
		for _, u := range add {
			if err = tx.Model(&org).Association("Users").Append(&User{Model: Model{ID: u}}); err != nil {
				return err
			}
		}
		for _, u := range del {
			if err = tx.Model(&org).Association("Users").Delete(&User{Model: Model{ID: u}}); err != nil {
				return err
			}
		}

		return tx.Error
	})
}

func UpdateOrgsToMember(userId uint, add, del []uint) error {
	db := conf.GetDbClient()

	return db.Transaction(func(tx *gorm.DB) (err error) {
		user := User{Model: Model{ID: userId}}
		for _, o := range add {
			if err = tx.Model(&user).Association("Orgs").Append(&Org{Model: Model{ID: o}}); err != nil {
				return err
			}
		}
		for _, o := range del {
			if err = tx.Model(&user).Association("Orgs").Delete(&Org{Model: Model{ID: o}}); err != nil {
				return err
			}
		}

		return tx.Error
	})
}
