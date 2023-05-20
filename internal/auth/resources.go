package auth

import (
	"go-to-cloud/internal/models"
	"strconv"
)

var groupPolicies [][]string
var resourcePolicies [][][]string

func GroupPolicies() [][]string {
	return groupPolicies
}

func ResourcePolicies() [][][]string {
	return resourcePolicies
}

func init() {
	groupPolicies = [][]string{
		{string(models.Root), "*"},
		{string(models.Ops), string(models.Dev)},
		{string(models.Dev), string(models.Guest)},
	}

	resourcePolicies = [][][]string{
		{{string(models.Guest), strconv.Itoa(int(models.MainMenuProject)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectNew)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectDelete)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectUpdate)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectSourceCode)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectSourceCodeImport)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectSourceDelete)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCI)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCINew)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCIStart)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCIHistory)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResProjectCIDelete)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCD)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCDNew)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCDStart)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCDHistory)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCDRollback)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCDDelete)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectCDMonitor)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResProjectArtifacts)), "RESOURCE"}},

		{{string(models.Dev), strconv.Itoa(int(models.MainMenuConfigure)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.SubMenuConfigureCodeRepo)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResConfigureCodeRepoBind)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResConfigureCodeRepoUpdate)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureCodeRepoRemove)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.SubMenuConfigureBuildNode)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResConfigureBuildNodeBind)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureBuildNodeUpdate)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureBuildNodeRemove)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.SubMenuConfigureArtifactRepo)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureArtifactRepoBind)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureArtifactRepoUpdate)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureArtifactRepoRemove)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureArtifactDeleteHistory)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.SubMenuConfigureDeployRepo)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureDeployBind)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureDeployUpdate)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResConfigureDeployRemove)), "RESOURCE"}},

		{{string(models.Dev), strconv.Itoa(int(models.MainMenuMonitor)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResourceMonitorScale)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResourceMonitorRestart)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResourceMonitorDelete)), "RESOURCE"}},
		{{string(models.Dev), strconv.Itoa(int(models.ResourceMonitorShell)), "RESOURCE"}},
		{{string(models.Ops), strconv.Itoa(int(models.ResourceMonitorDeletePod)), "RESOURCE"}},

		{{string(models.Root), strconv.Itoa(int(models.MainMenuUsers)), "RESOURCE"}},
		{{string(models.Root), strconv.Itoa(int(models.SubMenuUser)), "RESOURCE"}},
		{{string(models.Root), strconv.Itoa(int(models.SubMenuOrgs)), "RESOURCE"}},
	}
}
