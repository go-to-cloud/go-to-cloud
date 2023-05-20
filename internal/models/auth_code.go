package models

type AuthCode int

const (
	// MainMenuProject 主菜单->项目
	MainMenuProject            AuthCode = 10000
	ResProjectNew              AuthCode = 10100 // 新建项目
	ResProjectDelete           AuthCode = 10201 // 删除项目
	ResProjectUpdate           AuthCode = 10202 // 更新项目信息
	ResProjectSourceCode       AuthCode = 10310 // 管理模块：代码仓库
	ResProjectSourceCodeImport AuthCode = 10311 // 管理模块：导入代码仓库
	ResProjectSourceDelete     AuthCode = 10312 // 管理模块：删除代码仓库
	ResProjectCI               AuthCode = 10320 // 管理模块：持续集成
	ResProjectCINew            AuthCode = 10321 // 管理模块：新建构建计划
	ResProjectCIStart          AuthCode = 10322 // 管理模块：开始构建
	ResProjectCIHistory        AuthCode = 10323 // 管理模块：构建计划
	ResProjectCIDelete         AuthCode = 10324 // 管理模块：删除计划
	ResProjectCD               AuthCode = 10330 // 管理模块：持续部署
	ResProjectCDNew            AuthCode = 10331 // 管理模块：新建部署方案
	ResProjectCDStart          AuthCode = 10332 // 管理模块：开始部署/重新部署
	ResProjectCDHistory        AuthCode = 10333 // 管理模块：部署历史
	ResProjectCDRollback       AuthCode = 10334 // 管理模块：回滚部署
	ResProjectCDDelete         AuthCode = 10334 // 管理模块：删除部署方案
	ResProjectCDMonitor        AuthCode = 10335 // 管理模块：转到应用监控
	ResProjectArtifacts        AuthCode = 10340 // 管理模块：制品仓库

	// MainMenuConfigure 主菜单->配置
	MainMenuConfigure                 AuthCode = 20000
	SubMenuConfigureCodeRepo          AuthCode = 20010 // 代码仓库
	ResConfigureCodeRepoBind          AuthCode = 20011 // 绑定代码仓库
	ResConfigureCodeRepoUpdate        AuthCode = 20012 // 编辑代码仓库
	ResConfigureCodeRepoRemove        AuthCode = 20013 // 移除代码仓库
	SubMenuConfigureBuildNode         AuthCode = 20020 // 构建节点
	ResConfigureBuildNodeBind         AuthCode = 20021 // 安装构建节点
	ResConfigureBuildNodeUpdate       AuthCode = 20022 // 更新构建节点
	ResConfigureBuildNodeRemove       AuthCode = 20023 // 卸载构建节点
	SubMenuConfigureArtifactRepo      AuthCode = 20030 // 制品仓库
	ResConfigureArtifactRepoBind      AuthCode = 20031 // 绑定制品仓库
	ResConfigureArtifactRepoUpdate    AuthCode = 20032 // 更新制品仓库
	ResConfigureArtifactRepoRemove    AuthCode = 20033 // 移除制品仓库
	ResConfigureArtifactDeleteHistory AuthCode = 20034 // 删除历史制品
	SubMenuConfigureDeployRepo        AuthCode = 20040 // 部署环境
	ResConfigureDeployBind            AuthCode = 20041 // 绑定部署环境
	ResConfigureDeployUpdate          AuthCode = 20042 // 更新部署环境
	ResConfigureDeployRemove          AuthCode = 20043 // 移除部署环境

	// MainMenuMonitor 主菜单->应用监控
	MainMenuMonitor          AuthCode = 30000
	ResourceMonitorScale     AuthCode = 30010 // 伸缩副本
	ResourceMonitorRestart   AuthCode = 30011 // 重新启动
	ResourceMonitorDelete    AuthCode = 30012 // 删除应用
	ResourceMonitorShell     AuthCode = 30013 // 容器Shell
	ResourceMonitorDeletePod AuthCode = 30014 // 删除容器

	// MainMenuUsers 主菜单->用户管理
	MainMenuUsers AuthCode = 40000
	SubMenuUser   AuthCode = 40010 // 用户管理
	SubMenuOrgs   AuthCode = 40020 // 组织管理
)

type Kind string

type KindPair struct {
	Key     Kind   `json:"key"`
	ValueCN string `json:"valueCN"`
	ValueEN string `json:"valueEN"`
}

const (
	Root  Kind = "root"
	Dev   Kind = "dev"
	Ops   Kind = "ops"
	Guest Kind = "guest"
)
