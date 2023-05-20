package builder

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	pipeline2 "go-to-cloud/internal/models/pipeline"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	s := "{\"dockerfile\":\"Dockerfile\",\"registry\":\"go-to-cloud-docker.pkg.coding.net/devops/kaniko\",\"isSecurity\":true,\"account\":\"13805720368\",\"password\":\"DO#we498\"}"
	var abc pipeline2.ArtifactScript
	err := json.Unmarshal([]byte(s), &abc)

	assert.NoError(t, err)
}
