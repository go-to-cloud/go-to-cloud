package kube

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"os"
	"testing"
	"text/template"
)

func TestValidate(t *testing.T) {
	config := AppDeployConfig{
		Name:     "HelloWorld",
		Image:    "Empty",
		Replicas: 1,
		Ports: []Port{
			{
				ServicePort:   0,
				ContainerPort: 10,
			},
		},
	}
	assert.ErrorContains(t, config.validate(), "服务端口范围 1~65535")

	config.Ports[0].ServicePort = 10
	config.Ports[0].ContainerPort = 0
	assert.ErrorContains(t, config.validate(), "容器端口范围 1~65535")
	config.Ports[0].ContainerPort = 10

	config.Image = ""
	assert.ErrorContains(t, config.validate(), "镜像地址没有配置")
	config.Image = "123"

	config.Replicas = -1
	assert.ErrorContains(t, config.validate(), "副本数量不允许小于0")
	config.Replicas = 1
}

func TestRollingUpValidate(t *testing.T) {
	config := &RollingUpdateStrategy{
		MaxSurge:       -1,
		MaxUnavailable: 0,
	}
	assert.ErrorContains(t, config.validate(), "最大激发数范围：1～100")
	config.MaxSurge = 101
	config.MaxUnavailable = 0
	assert.ErrorContains(t, config.validate(), "最大激发数范围：1～100")

	config.MaxSurge = 1
	config.MaxUnavailable = -1
	assert.ErrorContains(t, config.validate(), "最大无效数范围：1～100")
	config.MaxSurge = 1
	config.MaxUnavailable = 101
	assert.ErrorContains(t, config.validate(), "最大无效数范围：1～100")
}

func TestResourceLimitValidate(t *testing.T) {
	zero := 0
	nonzero := 1
	greater := 2

	config := &ResLimits{}

	config.MemLimits = &zero
	assert.ErrorContains(t, config.validate(), "内存分配上限不能小于或等于0M")
	config.MemLimits = &nonzero
	config.MemRequest = &zero
	assert.ErrorContains(t, config.validate(), "内存分配不能小于或等于0M")

	config.MemLimits = &nonzero
	config.MemRequest = &greater
	assert.ErrorContains(t, config.validate(), "内存分配上限不能小于内存分配数量")

	config.MemRequest = nil
	config.MemLimits = nil

	config.CpuLimits = &zero
	assert.ErrorContains(t, config.validate(), "CPU分配上限不能少于或等于0")
	config.CpuLimits = &nonzero
	config.CpuRequest = &zero
	assert.ErrorContains(t, config.validate(), "CPU分配量不能少于或等于0")

	config.CpuLimits = &nonzero
	config.CpuRequest = &greater
	assert.ErrorContains(t, config.validate(), "CPU分配上限不能小于CPU分配数量")
}
func TestYamlTemplateParser(t *testing.T) {

	cpuRequest, cpuLimit, memRequest, memLimit := 1, 2, 3, 4
	config := AppDeployConfig{
		Name:     "HelloWorld",
		Image:    "Empty",
		Replicas: 1,
		Ports: []Port{
			{
				ServicePort:   10,
				ContainerPort: 10,
			},
			{
				ServicePort:   11,
				ContainerPort: 11,
				NodePort:      91,
			},
		},
		ResourceLimit: &ResLimits{
			CpuRequest: &cpuRequest,
			CpuLimits:  &cpuLimit,
			MemRequest: &memRequest,
			MemLimits:  &memLimit,
		},
		RollingUpdate: &RollingUpdateStrategy{
			MaxSurge:       12,
			MaxUnavailable: 30,
		},
		Dependencies: []DependContainer{
			{
				ContainerName: "C1",
				Namespace:     "helloWorld",
			},
			{
				ContainerName: "c2",
				Namespace:     "helloWorld2",
			},
		},
	}

	assert.Error(t, config.validate())

	tpl, err := template.New("deploy").Parse(YamlTplDeployment)
	assert.NoError(t, err)

	tpl, err = template.New("service").Parse(YamlTplService)
	assert.NoError(t, err)

	assert.Equal(t, "NodePort", config.PortType)
	err = tpl.Execute(os.Stdout, config)

	config.Ports[1].NodePort = 0
	assert.NoError(t, config.validate())
	err = tpl.Execute(os.Stdout, config)
	assert.Equal(t, "ClusterIP", config.PortType)

	assert.NoError(t, err)
}

func TestYamlTemplateCheck(t *testing.T) {

	cpuRequest, cpuLimit, memRequest, memLimit := 1, 2, 3, 4

	config := AppDeployConfig{
		Name:     "HelloWorld",
		Image:    "Nginx:latest",
		Replicas: 1,
		ResourceLimit: &ResLimits{
			CpuRequest: &cpuRequest,
			CpuLimits:  &cpuLimit,
			MemRequest: &memRequest,
			MemLimits:  &memLimit,
		},
		RollingUpdate: &RollingUpdateStrategy{
			MaxSurge:       12,
			MaxUnavailable: 30,
		},
		Dependencies: []DependContainer{
			{
				ContainerName: "C1",
				Namespace:     "helloWorld",
			},
			{
				ContainerName: "c2",
				Namespace:     "helloWorld2",
			},
		},
	}

	assert.NoError(t, config.validate())
}

func TestDecodeYaml(t *testing.T) {
	podCfg := corev1.PodApplyConfiguration{}
	spec := ""
	err := DecodeYaml(&spec, &podCfg)
	assert.EqualValues(t, fmt.Errorf("文件内容为空：%v", io.EOF), err)
}

func TestGetYamlFromTemple(t *testing.T) {

	cpuRequest, cpuLimit, memRequest, memLimit := 1, 2, 3, 4

	config := AppDeployConfig{
		Name:     "HelloWorld",
		Image:    "Nginx:latest",
		Replicas: 1,
		ResourceLimit: &ResLimits{
			CpuRequest: &cpuRequest,
			CpuLimits:  &cpuLimit,
			MemRequest: &memRequest,
			MemLimits:  &memLimit,
		},
		RollingUpdate: &RollingUpdateStrategy{
			MaxSurge:       12,
			MaxUnavailable: 30,
		},
		Dependencies: []DependContainer{
			{
				ContainerName: "C1",
				Namespace:     "helloWorld",
			},
			{
				ContainerName: "c2",
				Namespace:     "helloWorld2",
			},
		},
		Ports: []Port{
			{
				ServicePort:   10,
				ContainerPort: 10,
			},
			{
				ServicePort:   11,
				ContainerPort: 11,
				NodePort:      30010,
			},
		},
	}

	t1, t2, err := GetYamlFromTemple(&config)

	assert.NoError(t, err)
	assert.NotEmpty(t, *t1)
	assert.NotEmpty(t, *t2)
}

func TestProbValidate(t *testing.T) {
	prob := &ProbeConfigure{
		Path:             "path",
		Timeout:          10,
		Delay:            10,
		Period:           10,
		Port:             30010,
		SuccessThreshold: 10,
		FailureThreshold: 10,
	}
	err := prob.validate("abc")

	assert.NoError(t, err)

	name := "abc"

	prob.Path = ""
	err = prob.validate(name)
	assert.ErrorContains(t, err, name+"路径未配置")
	prob.Path = "DDD"

	prob.Timeout = 0
	err = prob.validate(name)
	assert.ErrorContains(t, err, name+"超时时间不能小于0")
	prob.Timeout = 10

	prob.Delay = 0
	err = prob.validate(name)
	assert.ErrorContains(t, err, name+"延时探测时间不能小于0")
	prob.Delay = 10

	prob.Period = 0
	err = prob.validate(name)
	assert.ErrorContains(t, err, name+"探测频率不能小于0")
	prob.Period = 10

	prob.Port = 0
	err = prob.validate(name)
	assert.ErrorContains(t, err, name+"探测端口范围")
	prob.Port = 65536
	assert.ErrorContains(t, err, name+"探测端口范围")
	prob.Port = 10

	prob.SuccessThreshold = 0
	err = prob.validate(name)
	assert.ErrorContains(t, err, name+"成功阈值不能小于0")
	prob.SuccessThreshold = 10

	name = "探针"
	prob.FailureThreshold = 0
	err = prob.validate()
	assert.ErrorContains(t, err, name+"失败阈值不能小于0")
	prob.FailureThreshold = 10
}
