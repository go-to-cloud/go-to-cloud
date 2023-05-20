package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"sync"
)

var db *gorm.DB

func getDbConnectionString(c *Conf) *string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Db.User, c.Db.Password, c.Db.Host, c.Db.Schema)
	return &dsn
}

var onceDb sync.Once

// GetDbClient 获取数据库连接对象
func GetDbClient() *gorm.DB {
	onceDb.Do(func() {
		if "1" == os.Getenv("UnitTestEnv") {
			db, _ = getInMemoryDbClient()
			db = db.Debug()
		} else {
			dsn := getDbConnectionString(getConf())
			_db, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			})
			if err != nil {
				panic(err)
			}

			if Environment.IsDevelopment() {
				_db = _db.Debug()
			}
			db = _db
		}
	})
	return db
}

// getInMemoryDbClient 获取内存数据库对象，仅限单元测试使用
func getInMemoryDbClient() (*gorm.DB, error) {
	if client, err := gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return client, nil
	}
}
