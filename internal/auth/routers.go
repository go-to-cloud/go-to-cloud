package auth

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/configure/artifact"
	"go-to-cloud/internal/controllers/configure/buildEnv"
	"go-to-cloud/internal/controllers/configure/builder"
	"go-to-cloud/internal/controllers/configure/deploy/k8s"
	"go-to-cloud/internal/controllers/configure/scm"
	"go-to-cloud/internal/controllers/monitor"
	"go-to-cloud/internal/controllers/projects"
	"go-to-cloud/internal/controllers/users"
	"go-to-cloud/internal/models"
)

type RestfulMethod string

const (
	PUT    RestfulMethod = "PUT"
	GET    RestfulMethod = "GET"
	DELETE RestfulMethod = "DELETE"
	POST   RestfulMethod = "POST"
)

type RouterMap struct {
	Url     string
	Methods []RestfulMethod
	Func    func(ctx *gin.Context)
	Kinds   []models.Kind
}

var RouterMaps []RouterMap

func init() {
	// Dev具备Guest的所有权限，Ops具备Dev的所有权限
	RouterMaps = make([]RouterMap, 0)

	// webapi
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/auths", []RestfulMethod{GET}, auth.GetAuthCodes, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/kinds", []RestfulMethod{GET}, users.AllKinds, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user", []RestfulMethod{PUT, POST}, users.UpsertUser, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/info", []RestfulMethod{GET}, users.Info, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/org/list", []RestfulMethod{GET}, users.OrgList, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/org", []RestfulMethod{PUT, POST}, users.UpsertOrg, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/org/:orgId", []RestfulMethod{DELETE}, users.DeleteOrg, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/list", []RestfulMethod{GET}, users.List, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/joined/:orgId", []RestfulMethod{GET}, users.Joined, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/join/:orgId", []RestfulMethod{PUT}, users.Join, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId/orgs/joined", []RestfulMethod{GET}, users.Belonged, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId/join", []RestfulMethod{PUT}, users.Belong, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId", []RestfulMethod{DELETE}, users.DeleteUser, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId/password/reset", []RestfulMethod{PUT}, users.ResetPassword, []models.Kind{models.Dev}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/build/env", []RestfulMethod{GET}, buildEnv.BuildEnv, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/build/cmd", []RestfulMethod{GET}, buildEnv.BuildCmd, []models.Kind{models.Ops}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo", []RestfulMethod{GET}, scm.QueryCodeRepos, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo/bind", []RestfulMethod{POST}, scm.BindCodeRepo, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo", []RestfulMethod{PUT}, scm.UpdateCodeRepo, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo/:id", []RestfulMethod{DELETE}, scm.RemoveCodeRepo, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo/testing", []RestfulMethod{POST}, scm.Testing, []models.Kind{models.Dev}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/testing", []RestfulMethod{POST}, artifact.Testing, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/bind", []RestfulMethod{POST}, artifact.BindArtifactRepo, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact", []RestfulMethod{PUT}, artifact.UpdateArtifactRepo, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact", []RestfulMethod{GET}, artifact.QueryArtifactRepo, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/:id", []RestfulMethod{DELETE}, artifact.RemoveArtifactRepo, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/:id", []RestfulMethod{GET}, artifact.QueryArtifactItems, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/image/:imageId", []RestfulMethod{DELETE}, artifact.DeleteImage, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/images/hashId/:hashId", []RestfulMethod{DELETE}, artifact.DeleteImageByHashId, []models.Kind{models.Ops}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s/testing", []RestfulMethod{POST}, k8s.Testing, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s/bind", []RestfulMethod{POST}, k8s.BindK8sRepo, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s", []RestfulMethod{PUT}, k8s.UpdateK8sRepo, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s", []RestfulMethod{GET}, k8s.QueryK8sRepos, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s/:id", []RestfulMethod{DELETE}, k8s.RemoveK8sRepo, []models.Kind{models.Ops}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/install/k8s", []RestfulMethod{POST}, builder.K8sInstall, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/nodes/k8s", []RestfulMethod{GET}, builder.QueryNodesOnK8s, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/nodes/k8s/available", []RestfulMethod{GET}, builder.QueryAvailableNodesOnK8s, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/node/:id", []RestfulMethod{DELETE}, builder.Uninstall, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/node", []RestfulMethod{PUT}, builder.UpdateBuilderNode, []models.Kind{models.Ops}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/projects", []RestfulMethod{POST}, projects.Create, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId", []RestfulMethod{DELETE}, projects.DeleteProject, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/list", []RestfulMethod{GET}, projects.List, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/coderepo", []RestfulMethod{GET}, projects.CodeRepo, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects", []RestfulMethod{PUT}, projects.UpdateProject, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/import", []RestfulMethod{POST}, projects.ImportSourceCode, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/imported", []RestfulMethod{GET}, projects.ListImportedSourceCode, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/sourcecode/:id", []RestfulMethod{DELETE}, projects.DeleteSourceCode, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/src/:sourceCodeId", []RestfulMethod{GET}, projects.ListBranches, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline", []RestfulMethod{POST}, projects.NewBuildPlan, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline", []RestfulMethod{GET}, projects.QueryBuildPlan, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/:pipelineId/history", []RestfulMethod{GET}, projects.QueryBuildPlanHistory, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/state", []RestfulMethod{GET}, projects.QueryBuildPlanState, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/:id", []RestfulMethod{DELETE}, projects.DeleteBuildPlan, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/:id/build", []RestfulMethod{POST}, projects.StartBuildPlan, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:id", []RestfulMethod{PUT}, projects.Deploying, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:id/rollback/:historyId", []RestfulMethod{PUT}, projects.Rollback, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/apps", []RestfulMethod{GET}, projects.QueryDeployments, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/app/:deploymentId/history", []RestfulMethod{GET}, projects.QueryDeploymentHistory, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/app", []RestfulMethod{POST}, projects.CreateDeployment, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:id", []RestfulMethod{DELETE}, projects.DeleteDeployment, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:k8sRepoId/namespaces", []RestfulMethod{GET}, projects.QueryNamespaces, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/env", []RestfulMethod{GET}, projects.QueryDeploymentEnv, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifacts/:querystring", []RestfulMethod{GET}, projects.QueryArtifacts, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifact/:artifactId/tags", []RestfulMethod{GET}, projects.QueryArtifactTags, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifact/:artifactId", []RestfulMethod{GET}, projects.QueryArtifactItemsByProjectId, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifact", []RestfulMethod{GET}, projects.QueryArtifactsByProjectId, []models.Kind{models.Guest}})

	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/query", []RestfulMethod{GET}, monitor.Query, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/restart", []RestfulMethod{PUT}, monitor.Restart, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/delete", []RestfulMethod{PUT}, monitor.DeletePod, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/scale", []RestfulMethod{PUT}, monitor.Scale, []models.Kind{models.Ops}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/pods/:deploymentId", []RestfulMethod{GET}, monitor.QueryPods, []models.Kind{models.Guest}})
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/delete/:deploymentId", []RestfulMethod{DELETE}, monitor.DeleteDeployment, []models.Kind{models.Ops}})

	// websocket
	RouterMaps = append(RouterMaps, RouterMap{"/ws/monitor/:k8s/pod/:deploymentId/:podName/log", []RestfulMethod{GET}, monitor.DisplayLog, []models.Kind{models.Dev}})
	RouterMaps = append(RouterMaps, RouterMap{"/ws/monitor/:k8s/pod/:deploymentId/:podName/shell", []RestfulMethod{GET}, monitor.Interactive, []models.Kind{models.Dev}})
}
