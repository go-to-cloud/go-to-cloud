package buildEnv

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/builder/lang"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

type cmd struct {
	UnitTest  string `json:"unitTest"`
	LintCheck string `json:"lintCheck"`
}

var envCmd map[string]*cmd

// BuildCmd 构建命令
// @Tags BuildConfigure
// @Description 构建命令
// @Success 200 {object} cmd
// @Router /api/configure/build/cmd [get]
// @Param   env     query     string     true  "Build Env"     example("dot-net-3.1")
// @Security JWT
func BuildCmd(ctx *gin.Context) {
	env := ctx.Query("env")

	v := envCmd[env]
	if v == nil {
		msg := fmt.Sprintf("language '%s' not supported", env)
		response.Fail(ctx, http.StatusNotFound, &msg)
	} else {
		response.Success(ctx, v)
	}
}
func init() {
	envCmd = make(map[string]*cmd)

	dotNetCmd := &cmd{
		UnitTest:  "dotnet test --collect:\"XPlat Code Coverage\" --logger \"html;logfilename=testresults.html\"",
		LintCheck: "dotnet format --verify-no-changes --report .",
	}
	envCmd[lang.DotNet3] = dotNetCmd
	envCmd[lang.DotNet5] = dotNetCmd
	envCmd[lang.DotNet6] = dotNetCmd
	envCmd[lang.DotNet7] = dotNetCmd

	golangCmd := &cmd{
		UnitTest: "go test -cover -test.short ./... | tee testresults.txt",
		LintCheck: `wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s 
                    ./bin/golangci-lint run ./... | tee lintcheck-result.txt`,
	}
	envCmd[lang.Go116] = golangCmd
	envCmd[lang.Go117] = golangCmd
	envCmd[lang.Go118] = golangCmd
	envCmd[lang.Go119] = golangCmd
	envCmd[lang.Go120] = golangCmd

	envCmd[lang.NodeJs14] = &cmd{}
	envCmd[lang.NodeJs16] = &cmd{}
	envCmd[lang.NodeJs18] = &cmd{}

	envCmd[lang.MavenJDK8] = &cmd{}
}
