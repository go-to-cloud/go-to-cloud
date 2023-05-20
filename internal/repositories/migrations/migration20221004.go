package migrations

import (
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
)

type migration20221004 struct {
}

func (m *migration20221004) Up(db *gorm.DB) error {

	if !db.Migrator().HasTable(&repo.CodeRepo{}) {
		err := db.AutoMigrate(&repo.CodeRepo{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.ArtifactRepo{}) {
		err := db.AutoMigrate(&repo.ArtifactRepo{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.ArtifactDockerImages{}) {
		err := db.AutoMigrate(&repo.ArtifactDockerImages{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.K8sRepo{}) {
		err := db.AutoMigrate(&repo.K8sRepo{})
		if err != nil {
			return err
		}
	}

	//if !db.Migrator().HasTable(&repo.GitRepo{}) {
	//	err := db.AutoMigrate(&repo.GitRepo{})
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	if !db.Migrator().HasTable(&repo.Project{}) {
		err := db.AutoMigrate(&repo.Project{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.ProjectSourceCode{}) {
		err := db.AutoMigrate(&repo.ProjectSourceCode{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.Pipeline{}) {
		err := db.AutoMigrate(&repo.Pipeline{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.PipelineSteps{}) {
		err := db.AutoMigrate(&repo.PipelineSteps{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.BuilderNode{}) {
		err := db.AutoMigrate(&repo.BuilderNode{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.PipelineHistory{}) {
		err := db.AutoMigrate(&repo.PipelineHistory{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.Deployment{}) {
		err := db.AutoMigrate(&repo.Deployment{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&repo.DeploymentHistory{}) {
		err := db.AutoMigrate(&repo.DeploymentHistory{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *migration20221004) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&repo.CodeRepo{},
		&repo.ArtifactRepo{},
		&repo.ArtifactDockerImages{},
		&repo.K8sRepo{},
		//&repo.GitRepo{},
		&repo.Project{},
		&repo.ProjectSourceCode{},
		&repo.Pipeline{},
		&repo.PipelineSteps{},
		&repo.BuilderNode{},
		&repo.PipelineHistory{},
		&repo.Deployment{},
		&repo.DeploymentHistory{},
	)
}
