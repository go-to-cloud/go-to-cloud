package builder

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/pkg/kube"
	"strings"
	"testing"
	"text/template"
)

func TestBuildPodSpec(t *testing.T) {

	tpl, err := template.New("k8s").Parse(kube.BuildTemplate)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, kube.PodSpecConfig{
		TaskName:   "Test",
		SourceCode: "DFDF.git",
		Branch:     "test-branch",
		Sdk:        "6.0",
		Steps: []kube.Step{
			{
				Command: "c1",
			},
			{
				Command: "c2",
			},
		},
	})

	assert.NoError(t, err)

	yml := strings.TrimSpace(buf.String())

	assert.True(t, len(yml) > 0)
}
