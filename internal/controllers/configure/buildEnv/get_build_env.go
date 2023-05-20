package buildEnv

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/builder/lang"
	"go-to-cloud/internal/pkg/response"
)

type envGroup struct {
	Label   string `json:"label"`
	Options []struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"options"`
}

// BuildEnv 构建环境
// @Tags BuildConfigure
// @Description 构建环境
// @Success 200 {array} envGroup
// @Router /api/configure/build/env [get]
// @Security JWT
func BuildEnv(ctx *gin.Context) {
	response.Success(ctx, []envGroup{
		{
			Label: ".Net",
			Options: []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			}{
				{
					Value: lang.DotNet3,
					Label: ".NET Core 3.1",
				}, {
					Value: lang.DotNet5,
					Label: ".NET 5",
				}, {
					Value: lang.DotNet6,
					Label: ".NET 6",
				}, {
					Value: lang.DotNet7,
					Label: ".NET 7",
				},
			},
		},
		{
			Label: "Golang",
			Options: []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			}{
				{
					Value: lang.Go116,
					Label: "Go 1.16",
				}, {
					Value: lang.Go117,
					Label: "Go 1.17",
				}, {
					Value: lang.Go118,
					Label: "Go 1.18",
				}, {
					Value: lang.Go119,
					Label: "Go 1.19",
				}, {
					Value: lang.Go120,
					Label: "Go 1.20",
				},
			},
		},
		{
			Label: "NodeJS",
			Options: []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			}{
				{
					Value: lang.NodeJs16,
					Label: "NodeJS 16",
				}, {
					Value: lang.NodeJs14,
					Label: "NodeJS 14",
				}, {
					Value: lang.NodeJs18,
					Label: "NodeJS 18",
				},
			},
		},
		{
			Label: "Java",
			Options: []struct {
				Value string `json:"value"`
				Label string `json:"label"`
			}{
				{
					Value: lang.MavenJDK8,
					Label: "Maven JDK8",
				},
			},
		},
	})
}
