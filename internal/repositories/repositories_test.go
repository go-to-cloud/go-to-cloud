package repositories

import (
	"go-to-cloud/conf"
)

func prepareDb() {
	if conf.GetDbClient().Migrator().HasTable(Org{}) {
		conf.GetDbClient().Migrator().DropTable(&Org{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&Org{})

	if conf.GetDbClient().Migrator().HasTable(User{}) {
		conf.GetDbClient().Migrator().DropTable(&User{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&User{})

	if conf.GetDbClient().Migrator().HasTable(&PipelineSteps{}) {
		conf.GetDbClient().Migrator().DropTable(&PipelineSteps{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&PipelineSteps{})

	if conf.GetDbClient().Migrator().HasTable(&Pipeline{}) {
		conf.GetDbClient().Migrator().DropTable(&Pipeline{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&Pipeline{})

	if conf.GetDbClient().Migrator().HasTable(&PipelineHistory{}) {
		conf.GetDbClient().Migrator().DropTable(&PipelineHistory{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&PipelineHistory{})

	if conf.GetDbClient().Migrator().HasTable(&BuilderNode{}) {
		conf.GetDbClient().Migrator().DropTable(&BuilderNode{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&BuilderNode{})

	if conf.GetDbClient().Migrator().HasTable(&ArtifactRepo{}) {
		conf.GetDbClient().Migrator().DropTable(&ArtifactRepo{})
	}
	conf.GetDbClient().Migrator().AutoMigrate(&ArtifactRepo{})
}
