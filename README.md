# GO-TO-CLOUD 代码极速上云工具

[![Build Status](https://github.com/go-to-cloud/go-to-cloud/actions/workflows/codecov.yml/badge.svg?branch=main)](https://github.com/go-to-cloud/go-to-cloud/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/go-to-cloud/go-to-cloud/branch/main/graph/badge.svg?token=9Y81AN6KUA)](https://codecov.io/gh/go-to-cloud/go-to-cloud)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/go-to-cloud/go-to-cloud/blob/main/LICENSE)

## Intro 介绍

### [官网](https://www.gotocloud.vip) | [在线体验](https://75o056r201.goho.co)

> 体验账号：gotocloud / gtcgct123!

**Go-To-Cloud**，简称`GTC`，是一款快速将代码构建并发布到K8S的上云工具。不需要学习专业的k8s、容器镜像知识也能轻松上手使用，工具的部署也做到了极简，抛开了大多数CI/CD工具对第三方工具(`jenkins`、`ArgoCD`等)的依赖，仅需MySQL、镜像仓库，并指定用于构建和部署的K8S机器，即可搭建一套自己的CI/CD服务平台。

> Go-To-Cloud正在奔跑中成长，非常渴望倾听你们的[**声音**](https://github.com/go-to-cloud/go-to-cloud/issues)

## Features 功能

- 项目管理
  - 创建项目，并配置项目的代码分支、构建及部署环境
  - 构建项目：将项目打包为镜像放到镜像仓库并打上tag
  - 部署项目：配置项目运行的环境、端口、健康检查等参数，并将项目镜像`发布`或`回滚`到目标k8s环境
  - 监控项目：查看在k8s中运行的项目实例，并可`伸缩`、`重启`实例，以及查看`容器日志`和进入容器内部执行`Shell`命令
- 配置管理
  - 代码仓库：用于创建项目的时候指定代码及分支，以及构建过程中拉取代码
  - 构建节点：用于执行构建和打包镜像指令的k8s环境
  - 制品仓库：构建的镜像存放的仓库，管理构建的结果、复制镜像地址，查看历史构建记录
  - 部署环境：项目最终运行的k8s环境
- 应用监控：伸缩、重启及管理应用
- 账号管理：
  - 创建用户：系统的使用者
  - 创建组织：项目是由`组织`维护，同一`组织`里的用户可以查看及管理归属的`项目`
  - 权限管理：项目权限是基于`RBAC`模型管理，[查看介绍](https://github.com/go-to-cloud/go-to-cloud/tree/main/internal/auth)
  
## Install 安装

[使用文档](https://docs.gotocloud.vip)

## Develop 开发环境

### 开发环境

- go `1.19+`
- mysql `8.0+`

### 运行环境
- 制品仓库

  如果没有，可以利用`docker`在本地搭建一个
```shell
# 本地搭建registry:v2
docker run -d -p 5000:5000 --restart always --name registry registry:2
```

注意需要在docker中配置允许非加密地址
```yaml
"insecure-registries": [
  "127.0.0.1:5000",
  .... 其他地址
]
```

- 代码仓库

  目前支持 `gitlab` `gitea` `github`

- 构建节点
  k8s

- 部署环境(非必需)

  如果仅构建镜像，则不需要配置`部署环境`，部署环境目前仅支持k8s，试用时可以与构建节点用同一个k8s集群

### 准备工作

1. 创建数据库
```sql
create database gotocloud character set utf8mb4;
```

2. 修改配置
> [配置位于: con/appsettings.yaml](https://github.com/go-to-cloud/go-to-cloud/blob/main/conf/appsettings.yaml)

```yaml
db: # 数据库配置
  user: # 数据库用户名
  password: # 数据库密码
  host: # 数据库地址
  schema: gotocloud # 数据库名，默认不用修改
jwt: # jwt配置，建议修改security，其他可以保持不变
  security: thisisunsafeuntilyouchangit 
  realm: GOTOCLOUD
  idkey: id
builder: # 打包镜像的服务，一般不用修改
  kaniko: go-to-cloud-docker.pkg.coding.net/devops/kaniko/executor:v1.9.1-debug
```
3. 运行
> 切换至项目根目录（main.go所在目录）
```bash
go mod tidy
go run .
```

4. 前端

> 项目前端位于另一个仓库: [前端仓库](https://github.com/go-to-cloud/go-to-cloud-frontend)


## About 联系方式

E-Mail: go-to-cloud@qq.com
