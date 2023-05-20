package migrations

import (
	"encoding/json"
	"go-to-cloud/internal/models"
	repo "go-to-cloud/internal/repositories"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Migration20220831 struct {
}

func (m *Migration20220831) Up(db *gorm.DB) error {

	userOrgRelNotExists := false

	if !db.Migrator().HasTable(&repo.User{}) {
		if err := db.AutoMigrate(&repo.User{}); err != nil {
			return err
		}
		userOrgRelNotExists = true
	}
	if !db.Migrator().HasTable(&repo.Org{}) {
		if err := db.AutoMigrate(&repo.Org{}); err != nil {
			return err
		}
		userOrgRelNotExists = true
	}

	if userOrgRelNotExists {
		org := &repo.Org{
			Name: "ROOT",
		}
		db.Debug().Create(org)

		user := &repo.User{
			Account:  models.RootUserName,
			RealName: "系统管理员",
			Kind: func() datatypes.JSON {
				s, _ := json.Marshal([]string{string(models.Root)})
				return s
			}(),
			Pinyin:     "xitongguanliyuan",
			PinyinInit: "xtgly",
			Orgs:       []*repo.Org{org},
		}
		initRootPassword := "root"
		if err := user.SetPassword(&initRootPassword); err != nil {
			return err
		}
		// ignore error check
		db.Debug().Create(user)
		db.Debug().Save(user)

		guest := &repo.User{
			Account:  "guest",
			RealName: "游客", Kind: func() datatypes.JSON {
				s, _ := json.Marshal([]string{string(models.Guest)})
				return s
			}(),
			Pinyin:     "youke",
			PinyinInit: "yk",
			Orgs:       []*repo.Org{org},
		}
		initRootPassword2 := "guest"
		if err := guest.SetPassword(&initRootPassword2); err != nil {
			return err
		}
		if err := db.Debug().Create(guest).Error; err != nil {
			return err
		}
		if err := db.Debug().Save(guest).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m *Migration20220831) Down(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		&repo.Org{},
		&repo.User{},
	); err != nil {
		return err
	}

	return db.Migrator().DropTable("orgs_users_rel")
}
