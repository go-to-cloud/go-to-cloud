package kube

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"strings"
	"text/template"
)

type ConfigValidation interface {
	validate(args ...string) error
}

// DecodeYaml 解析yaml文件，支持kind: Deployment / Service / Pod
func DecodeYaml[T Kinds](yml *string, conf *T) error {
	buf := bytes.NewBufferString(*yml)
	decoder := yaml.NewYAMLOrJSONDecoder(buf, 4096)
	ext := runtime.RawExtension{}
	if err := decoder.Decode(&ext); err != nil {
		if err == io.EOF {
			return fmt.Errorf("文件内容为空：%v", err)
		} else {
			return fmt.Errorf("解析失败：%v", err)
		}
	}

	ext.Raw = bytes.TrimSpace(ext.Raw)
	if len(ext.Raw) == 0 {
		return fmt.Errorf("文件内容为空")
	}

	return json.Unmarshal(ext.Raw, conf)
}

// validate 校验入参，并做些适当的调整，比如将名称调整为小写
func (cfg *AppDeployConfig) validate() error {
	cfg.Name = strings.ReplaceAll(strings.ToLower(cfg.Name), "_", "-")
	if cfg.Ports != nil && len(cfg.Ports) > 0 {
		cfg.PortType = "ClusterIP"
		for i, p := range cfg.Ports {

			if p.ServicePort <= 0 || p.ServicePort > 65535 {
				return errors.New("服务端口范围 1~65535")
			}

			if p.ContainerPort <= 0 || p.ContainerPort > 65535 {
				return errors.New("容器端口范围 1~65535")
			}

			if p.NodePort > 0 {
				cfg.PortType = "NodePort"
				if p.NodePort < 30000 || p.NodePort > 32767 {
					return errors.New("NodePort端口范围：30000～32767；0时表示使用ClusterIP")
				}
			}

			if len(strings.TrimSpace(cfg.Ports[i].PortName)) == 0 {
				cfg.Ports[i].PortName = strings.Replace(fmt.Sprintf("p-%2d", i), " ", "0", 1)
			}
		}
	}

	if len(cfg.Image) == 0 {
		return errors.New("镜像地址没有配置")
	}

	if cfg.Replicas < 0 {
		return errors.New("副本数量不允许小于0")
	}

	if cfg.RollingUpdate != nil {
		if err := cfg.RollingUpdate.validate(); err != nil {
			return err
		}
	}

	if cfg.ResourceLimit != nil {
		if err := cfg.ResourceLimit.validate(); err != nil {
			return err
		}
	}

	if cfg.Liveness != nil {
		if err := cfg.Liveness.validate("存活探针"); err != nil {
			return err
		}
	}
	if cfg.Readiness != nil {
		if err := cfg.Readiness.validate("就绪探针"); err != nil {
			return err
		}
	}

	return nil
}

// GetYamlFromTemple 从模板创建service.yaml
// @appName 应用名称
// @servicePort 服务暴露端口(port）
// @containerPort 容器端口（targetPort)
func GetYamlFromTemple(cfg *AppDeployConfig) (deployment, service *string, err error) {

	if err = cfg.validate(); err != nil {
		return nil, nil, err
	}

	template1, template2 := YamlTplDeployment, YamlTplService
	tpls := [2]*string{&template1, &template2}

	// 按约定，0：Deployment；1：Service
	parsed := make([]*string, 2)

	for i, t := range tpls {

		tpl, err := template.New("k8s").Parse(*t)
		if err != nil {
			return nil, nil, err
		}

		buf := new(bytes.Buffer)
		err = tpl.Execute(buf, *cfg)

		if err != nil {
			return nil, nil, err
		}

		yml := strings.TrimSpace(buf.String())
		parsed[i] = &yml
	}

	return parsed[0], parsed[1], nil
}

// ProbeConfigure 存活和就绪探针配置
type ProbeConfigure struct {
	Path             string // 路径
	Port             int    // 端口
	Delay            int    // 延迟探测时间
	Period           int    // 探测频率
	Timeout          int    // 探测超时
	SuccessThreshold int    // 成功阈值，探测失败后，连续成功几次会认为是成功；（liveness为常量1）；最小值为1
	FailureThreshold int    // 失败阈值，控制成功后，连续失败多少次会认为是失败；默认值3，最小值1
}

func (cfg *ProbeConfigure) validate(args ...string) error {

	if len(args) == 0 {
		args = []string{"探针"}
	}

	if len(cfg.Path) == 0 {
		return fmt.Errorf("%s路径未配置", args[0])
	}
	if cfg.Timeout <= 0 {
		return fmt.Errorf("%s超时时间不能小于0", args[0])
	}

	if cfg.Delay <= 0 {
		return fmt.Errorf("%s延时探测时间不能小于0", args[0])
	}

	if cfg.Period <= 0 {
		return fmt.Errorf("%s探测频率不能小于0", args[0])
	}

	if cfg.Port <= 0 || cfg.Port >= 65536 {
		return fmt.Errorf("%s探测端口范围：1~65535", args[0])
	}
	if cfg.SuccessThreshold <= 0 {
		return fmt.Errorf("%s成功阈值不能小于0", args[0])
	}
	if cfg.FailureThreshold <= 0 {
		return fmt.Errorf("%s失败阈值不能小于0", args[0])
	}

	return nil
}

// RollingUpdateStrategy 滚动发布策略
type RollingUpdateStrategy struct {
	MaxSurge       int // 	最大激增数；即发布过程中，最多允许超出期望副本数的 pod 实例的数量；取值范围：1～100
	MaxUnavailable int //	最大无效数：滚动升级期间，相对于期望副本数能够允许有多少 pod 实例处于不可用状态；取值范围：1～100
}

func (cfg *RollingUpdateStrategy) validate() error {
	if cfg.MaxSurge <= 0 || cfg.MaxSurge >= 100 {
		return errors.New("最大激发数范围：1～100")
	}
	if cfg.MaxUnavailable <= 0 || cfg.MaxUnavailable >= 100 {
		return errors.New("最大无效数范围：1～100")
	}
	return nil
}

// ResLimits 资源限制
type ResLimits struct {
	CpuRequest *int `json:"cpuRequest"` // CPU分配资源；> 0，单位m
	CpuLimits  *int `json:"cpuLimits"`  // CPU资源上限；> 0，单位m
	MemRequest *int `json:"memRequest"` // 内存分配资源：单位：Mi
	MemLimits  *int `json:"memLimits"`  // 内存资源上限；单位：Mi
}

func (cfg *ResLimits) validate() error {
	if cfg.MemLimits != nil && *cfg.MemLimits <= 0 {
		return errors.New("内存分配上限不能小于或等于0M")
	}
	if cfg.MemRequest != nil && *cfg.MemRequest <= 0 {
		return errors.New("内存分配不能小于或等于0M")
	}
	if cfg.MemLimits != nil && cfg.MemRequest != nil && *cfg.MemLimits < *cfg.MemRequest {
		return errors.New("内存分配上限不能小于内存分配数量")
	}

	if cfg.CpuLimits != nil && (*cfg.CpuLimits <= 0) {
		return errors.New("CPU分配上限不能少于或等于0")
	}
	if cfg.CpuRequest != nil && (*cfg.CpuRequest <= 0) {
		return errors.New("CPU分配量不能少于或等于0")
	}

	if cfg.CpuLimits != nil && cfg.CpuRequest != nil && *cfg.CpuLimits < *cfg.CpuRequest {
		return errors.New("CPU分配上限不能小于CPU分配数量")
	}

	return nil
}

type DependContainer struct {
	ContainerName string // 容器名称
	Namespace     string // 容器所在名字空间
}

type Port struct {
	ServicePort   int    `json:"servicePort"`   // 应用运行端口(service.spec.ports.port)
	ContainerPort int    `json:"containerPort"` // 容器暴露端口(service.spec.ports.targetPort)
	NodePort      int    `json:"nodePort"`      // NodePort(30000～32767）
	PortName      string `json:"portName"`      // PortName，按规则P01～P99
}

type EnvVar struct {
	Name      string
	Value     string
	ValueFrom interface{} // TODO: Not Supported YET；参考：https://kubernetes.io/zh-cn/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
}

type ConfigMapVolume struct {
	Name          string
	ConfigMapName string
	MountPath     string
	SubPath       string
}

// ContainerExtract 额外信息，作为Containers节点下的附加信息，通常放Args、Command
type ContainerExtract struct {
	Key    string
	Values []string
}

// AppDeployConfig 应用部署配置
type AppDeployConfig struct {
	LabelSelector     string
	AppId             string // 对应Deployment ID
	Namespace         string // 名字空间
	Name              string //	应用名称
	Ports             []Port // 端口
	PortType          string // 端口类型：NodePort 或者 ClusterIP
	Image             string // 镜像地址
	ContainerExtracts []ContainerExtract
	Env               []EnvVar               // 环境变量
	Replicas          int                    //	运行副本数
	Liveness          *ProbeConfigure        // 存活检查
	Readiness         *ProbeConfigure        // 就绪检查
	RollingUpdate     *RollingUpdateStrategy // 滚动发布策略
	ResourceLimit     *ResLimits             //	资源限制
	Dependencies      []DependContainer      // 依赖容器
	NodeSelector      []NodeSelector         // 指定选择
	ConfigMaps        []ConfigMapVolume      // ConfigMap卷
}

type NodeSelector struct {
	LabelName  string
	LabelValue string
}

// YamlTplService Service ClusterIP方式Yaml模板
const YamlTplService = `
{{- if .Ports}}
apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}-service
spec:
  selector:
    app: {{.Name}}
  ports:
{{- range .Ports}}
  - protocol: TCP
    port: {{.ServicePort}}
    targetPort: {{.ServicePort}}
    name: {{.PortName}}
{{- if .NodePort}}
    nodePort: {{.NodePort}}
{{- end}}
{{- end}}
  type: {{.PortType}}
{{- end}}
`

// YamlTplDeployment Deployment Yaml模板
const YamlTplDeployment = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}-deployment
  labels:
    app: {{.Name}}
    deployed: {{.LabelSelector}}
    appId: {{.LabelSelector}}-{{.AppId}}
spec:
  replicas: {{.Replicas}}
{{- if .RollingUpdate}}
{{- with .RollingUpdate}}
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: {{.MaxSurge}}%
      maxUnavailable: {{.MaxUnavailable}}%
{{- end}}
{{- end}}
  selector:
    matchLabels:
      app: {{.Name}}
  template:
    metadata:
      labels:
        app: {{.Name}}
    spec:
{{- if .ConfigMaps}}
      volumes:
{{- range .ConfigMaps}}
      - configMap:
          defaultMode: 420
          name: {{.Name}}
        name: {{.ConfigMapName}}
{{- end}}
{{- end}}
{{- if .NodeSelector}}
      nodeSelector:
{{- range .NodeSelector}}
        {{.LabelName}}: {{.LabelValue}}
{{- end}}
{{- end}}
      containers:
      - name: {{.Name}}
        image: {{.Image}}
        imagePullPolicy: IfNotPresent
{{- if .ContainerExtracts}}
{{- range .ContainerExtracts}}
        {{.Key}}:
{{- range .Values}}
          - "{{.}}"
{{- end}}
{{- end}}
{{- end}}
{{- if .Env}}
        env:
{{- range .Env}}
{{- if .Value}}
        - name: {{.Name}}
          value: "{{.Value}}"
{{- end}}
{{- end}}
{{- end}}
{{- if .Ports}}
        ports:
{{- range .Ports}}
          - name: {{.PortName}}
            containerPort: {{.ContainerPort}}
{{- end}}
{{- end}}
{{- if .ConfigMaps}}
        volumeMounts:
{{- range .ConfigMaps}}
        - mountPath: {{.MountPath}}
          name: {{.ConfigMapName}}
          subPath: {{.SubPath}}
{{- end}}
{{- end}}
{{- if .ResourceLimit}}
        resources:
{{- with .ResourceLimit}}
          requests:
{{- if .CpuRequest}}
            cpu: {{.CpuRequest}}m
{{- end}}
{{- if .MemRequest}}
            memory: {{.MemRequest}}Mi
{{- end}}
          limits:
{{- if .CpuLimits}}
            cpu: {{.CpuLimits}}m
{{- end}}
{{- if .MemLimits}}
            memory: {{.MemLimits}}Mi
{{- end}}
{{- end}}
{{- end}}
{{- if .Liveness }}
{{- with .Liveness}}
        livenessProbe:
          httpGet:
            path: {{.Path}}
            port: {{.Port}}
          initialDelaySeconds: {{.Delay}}
          periodSeconds: {{.Period}}
          failureThreshold: {{.FailureThreshold}}
          successThreshold: {{.SuccessThreshold}}
{{- end}}
{{- end}}
{{- if .Readiness}}
{{- with .Readiness}}
        readinessProbe:
          httpGet:
            path: {{.Path}}
            port: {{.Port}}
          initialDelaySeconds: {{.Delay}}
		  periodSeconds: {{.Period}}
          failureThreshold: {{.FailureThreshold}}
          successThreshold: {{.SuccessThreshold}}
{{- end}}
{{- end}}
{{- if .Dependencies }}
    initContainers:
{{- range .Dependencies}}
      - name: init-service-{{$.Name}}
        image: busybox:1.28
        command: ['sh', '-c', "until nslookup {{.ContainerName}}.{{.Namespace}}.svc.cluster.local; do echo waiting for {{.ContainerName}}; sleep 2; done"]
{{- end}}
{{- end}}
`
