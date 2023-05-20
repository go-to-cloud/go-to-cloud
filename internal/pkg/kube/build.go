package kube

import (
	"bytes"
	"fmt"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/pipeline"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"strings"
	"text/template"
)

const LatestTag = "latest"

type Step struct {
	CommandType  pipeline.PlanStepType
	CommandText  string
	Command      string // 当StepType是Image时，Command无意义
	Dockerfile   string
	Context      string
	ArtifactName string // 制品名称
	ArtifactTag  string // 制品Tag
	Registry     struct {
		Url      string
		User     string
		Password string
		Security bool
	}
} // 构建步骤

type PodSpecConfig struct {
	LabelFlag    string
	LabelBuildId string
	BuildId      uint
	Namespace    string
	TaskName     string // pod name
	SourceCode   string // git url
	Branch       string // git branch
	Sdk          string // sdk 基础镜像
	Steps        []Step
}

func (m PodSpecConfig) IsBuildImage() bool {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return true
		}
	}
	return false
}

// GetBaseImage 获取构建镜像的基础镜像
func (m PodSpecConfig) GetBaseImage() string {
	return *conf.GetBuildImage()
}

func (m PodSpecConfig) GetDockerImageName() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return fmt.Sprintf("%s:%s", step.ArtifactName, step.ArtifactTag)
		}
	}
	return ""
}

func (m PodSpecConfig) GetLatestDockerImage() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return fmt.Sprintf("%s:%s", step.ArtifactName, LatestTag)
		}
	}
	return ""
}

func (m PodSpecConfig) GetDockerfile() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return step.Dockerfile
		}
	}
	return ""
}

func (m PodSpecConfig) GetContext() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return step.Context
		}
	}
	return ""
}

func (m PodSpecConfig) GetRegistryUrl() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			s1 := strings.TrimPrefix(step.Registry.Url, "http://")
			s2 := strings.TrimPrefix(s1, "https://")
			return s2
		}
	}
	return ""
}

func (m PodSpecConfig) GetRegistryUser() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return step.Registry.User
		}
	}
	return ""
}

func (m PodSpecConfig) GetRegistryPassword() string {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.Image {
			return step.Registry.Password
		}
	}
	return ""
}

func (m PodSpecConfig) IsExistsQaSteps() bool {
	for _, step := range m.Steps {
		if step.CommandType == pipeline.LintCheck || step.CommandType == pipeline.UnitTest {
			return true
		}
	}
	return false
}

// Build 构建任务
func (client *Client) Build(podSpecConfig *PodSpecConfig) error {
	spec, err := makeTemplate(podSpecConfig)

	if err != nil {
		return err
	}

	if conf.Environment.IsDevelopment() {
		fmt.Println(*spec)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	podCfg := corev1.PodApplyConfiguration{}

	if err := DecodeYaml(spec, &podCfg); err != nil {
		fmt.Println(err)
		return err
	}
	if _, e := client.ApplyPod(&podSpecConfig.Namespace, &podCfg); e != nil {
		fmt.Println(e)
		return e
	}

	return nil
}

func makeTemplate(spec *PodSpecConfig) (*string, error) {
	tpl, err := template.New("k8s-pod").Parse(BuildTemplate)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, *spec)
	if err != nil {
		return nil, err
	}

	yml := strings.TrimSpace(buf.String())

	return &yml, nil
}

const BuildTemplate string = `
apiVersion: v1
kind: Pod
metadata:
  name: {{.TaskName}}
  labels:
    builder: {{.LabelFlag}}
{{- if .LabelBuildId}}
    {{.LabelBuildId}}: {{.LabelBuildId}}-{{.BuildId}}
{{- end}}
spec:
    initContainers:
    - name: coderepo
      image: alpine/git
      imagePullPolicy: IfNotPresent
      command: 
      - /bin/sh
      - -ec
      - |
        git clone -b {{.Branch}} {{.SourceCode}} /git
{{- if .IsBuildImage}}
        echo '{"auths":{"{{.GetRegistryUrl}}":{"username":"{{.GetRegistryUser}}", "password":"{{.GetRegistryPassword}}"}}}' > /kaniko/.docker/config.json
{{- end}}
      tty: true
      volumeMounts:
      - name: workdir
        mountPath: "/git"
{{- if .IsBuildImage}}
      - name: kaniko-config
        mountPath: "/kaniko/.docker"
{{- end}}
    containers:
{{- if .IsBuildImage}}
    - name: image
      image: {{.GetBaseImage}}
      args:
        - "--dockerfile={{.GetDockerfile}}"
        - "--context=/workdir/{{.GetContext}}"
        - "--skip-tls-verify"
        - "--insecure"
        - "--destination={{.GetRegistryUrl}}/{{.GetDockerImageName}}"
        - "--destination={{.GetRegistryUrl}}/{{.GetLatestDockerImage}}"
      tty: true
      volumeMounts:
      - name: workdir
        mountPath: "/workdir"
      - name: kaniko-config
        mountPath: "/kaniko/.docker"
{{- end}}
{{- if .IsExistsQaSteps}}
    - name: compile
      image: {{.Sdk}}
      imagePullPolicy: IfNotPresent
      tty: true
      lifecycle: 
        preStop:
          exec:
            command: ["/bin/sh", "-c", "echo '构建任务完成'"]
      command:
      - /bin/sh
      - -ec
      - |
        cd /workdir
        set -e
{{- range .Steps}}
{{- if .Command}}
        echo ">>>{{.CommandText}} 开始"
        if {{.Command}};then
          echo ">>>{{.CommandText}} 成功"
        else
          echo ">>>{{.CommandText}} 失败"
        fi
        echo "\n"
{{- end}}
{{- end}}
      volumeMounts:
      - name: workdir
        mountPath: "/workdir"
{{- end}}
    restartPolicy: OnFailure
    volumes:
    - name: workdir
      emptyDir: {}
{{- if .IsBuildImage}}
    - name: kaniko-config
      emptyDir: {}
{{- end}}
`
