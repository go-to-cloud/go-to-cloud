package conf

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"sync"
)

var once sync.Once

type Conf struct {
	Builder struct {
		Kaniko string
	}
	Db struct {
		User     string
		Password string
		Host     string
		Schema   string
	}
	Jwt JWT
}

var conf *Conf

// getConf 读取配置
// 默认从配置文件取，如果配置文件中的db节点内容为空，则从环境变量取
// 如果配置文件不存在，则db从环境变量取，其他值使用默认值
func getConf() *Conf {
	once.Do(func() {
		if conf == nil {
			filePath := getConfFilePath()
			conf = getConfiguration(filePath)
		}
	})
	return conf
}

func isDbEmpty(conf *Conf) bool {
	return len(conf.Db.Host) == 0 || len(conf.Db.User) == 0 || len(conf.Db.Schema) == 0 || len(conf.Db.Password) == 0
}

// getConfiguration 读取配置
// 优先从配置文件读取，如果数据库相关配置为空，则从环境变量读取
func getConfiguration(filePath *string) *Conf {
	if file, err := os.ReadFile(*filePath); err != nil {
		return getConfFromEnv()
	} else {
		c := Conf{}
		err := yaml.Unmarshal(file, &c)
		if err != nil {
			return getConfFromEnv()
		}

		// 检查db节点是否为空
		if isDbEmpty(&c) {
			db := getConfFromEnv()
			c.Db.Host = db.Db.Host
			c.Db.User = db.Db.User
			c.Db.Schema = db.Db.Schema
			c.Db.Password = db.Db.Password
		}
		return &c
	}
}

func getConfFromEnv() *Conf {
	// 初始化conf
	dbUser := os.Getenv("dbuser")
	dbPwd := os.Getenv("dbpwd")
	dbUrl := os.Getenv("dburl")
	schema := os.Getenv("schema")

	conf := &Conf{
		Builder: struct {
			Kaniko string
		}{"go-to-cloud-docker.pkg.coding.net/devops/kaniko/executor:v1.9.1-debug"},
		Db: struct {
			User     string
			Password string
			Host     string
			Schema   string
		}{
			dbUser, dbPwd, dbUrl, schema,
		},
		Jwt: JWT{"thisisunsafeuntilyouchangit", "GOTOCLOUD", "id"},
	}

	return conf
}
