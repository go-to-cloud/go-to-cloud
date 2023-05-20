package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArtifactDockerImages_GetHashedCode(t *testing.T) {
	img := &ArtifactDockerImages{
		ArtifactRepoID: 100,
		Name:           "testing:latest",
		Tag:            "latest",
	}

	assert.Equal(t, "100,testing", img.GetHashedCode())
}

func TestTableName(t *testing.T) {
	img := &ArtifactDockerImages{}
	assert.Equal(t, "artifact_docker_images", img.TableName())
}
