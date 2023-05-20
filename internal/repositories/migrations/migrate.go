package migrations

import (
	"go-to-cloud/conf"
	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

var migrations []Migration

func init() {
	// 迁移对象必需按从旧到新的顺序添加
	migrations = []Migration{
		&Migration20220831{},
		&migration20220921{},
		&migration20221004{},
	}
}

func AutoMigrate() {
	db := conf.GetDbClient()

	// TODO：检查 '__migration'表中是否有最新变更记录，有则跳过迁移
	Migrate(db)
}

// Migrate 数据库变更同步
func Migrate(db *gorm.DB) {

	for i := 0; i < len(migrations); i++ {
		migrations[i].Up(db)
	}
}

// Rollback 数据库变更回滚
func Rollback(db *gorm.DB) {

	for i := len(migrations) - 1; i >= 0; i-- {
		migrations[i].Down(db)
	}
}
