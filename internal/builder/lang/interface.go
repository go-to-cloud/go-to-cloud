package lang

import (
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/repositories"
)

const (
	DotNet3 = "dot-net-3.1"
	DotNet5 = "dot-net-5"
	DotNet6 = "dot-net-6"
	DotNet7 = "dot-net-7"

	Go116 = "go-1.16"
	Go117 = "go-1.17"
	Go118 = "go-1.18"
	Go119 = "go-1.19"
	Go120 = "go-1.20"

	NodeJs14 = "nodejs-14"
	NodeJs16 = "nodejs-16"
	NodeJs18 = "nodejs-18"

	MavenJDK8 = "maven-jdk-8"
)

type Tpl interface {
	Sdk(env string) string //获取sdk镜像地址
	Steps(env string, steps []repositories.PipelineSteps) map[pipeline.PlanStepType]string
}

type TplCommon struct {
}

func (m *TplCommon) Steps(_ string, steps []repositories.PipelineSteps) map[pipeline.PlanStepType]string {
	rlt := make(map[pipeline.PlanStepType]string)
	for _, step := range steps {
		rlt[step.Type] = step.Script
	}
	return rlt
}
