package kube

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/pipeline"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"testing"
)

func TestGetRegistryUrl(t *testing.T) {

	host := "mockurl.com/registry/v2"
	p1 := &PodSpecConfig{
		Steps: []Step{
			{
				CommandType: pipeline.Image,
				Registry: struct {
					Url      string
					User     string
					Password string
					Security bool
				}{Url: "https://mockurl.com/registry/v2", User: "regUser", Password: "regPwd", Security: true},
			},
		},
	}
	assert.Equal(t, host, p1.GetRegistryUrl())

	p2 := &PodSpecConfig{
		Steps: []Step{
			{
				CommandType: pipeline.Image,
				Registry: struct {
					Url      string
					User     string
					Password string
					Security bool
				}{Url: "http://mockurl.com/registry/v2", User: "regUser", Password: "regPwd", Security: false},
			},
		},
	}
	assert.Equal(t, host, p2.GetRegistryUrl())

	p3 := &PodSpecConfig{
		Steps: []Step{
			{
				CommandType: pipeline.LintCheck,
				Registry: struct {
					Url      string
					User     string
					Password string
					Security bool
				}{Url: "http://mockurl.com/registry/v2", User: "regUser", Password: "regPwd", Security: false},
			},
		},
	}
	assert.Equal(t, "", p3.GetRegistryUrl())
}

func TestParseTemplate(t *testing.T) {
	podSpecConfig := &PodSpecConfig{
		Namespace:  "testNs",
		TaskName:   "testTask",
		SourceCode: "testSource",
		Branch:     "test-branch",
		Sdk:        "testSdk",
		Steps: []Step{
			{
				CommandType: pipeline.Image,
				CommandText: "image",
				Command:     "cli 1",
				Dockerfile:  "dockerfile.file",
				Registry: struct {
					Url      string
					User     string
					Password string
					Security bool
				}{Url: "http://mockurl.com/registry/v2", User: "regUser", Password: "regPwd", Security: false},
			},
			{
				CommandType: pipeline.LintCheck,
				CommandText: "image",
				Command:     "cli 2",
			},
		},
	}

	spec, err := makeTemplate(podSpecConfig)
	assert.NoError(t, err)

	err = DecodeYaml(spec, &corev1.PodApplyConfiguration{})
	fmt.Println(*spec)
	assert.NoError(t, err)
}
