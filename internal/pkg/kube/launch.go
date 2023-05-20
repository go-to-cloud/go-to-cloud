package kube

import (
	"fmt"
	"go-to-cloud/conf"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

func (client *Client) Launch(appDeployConfig *AppDeployConfig) error {
	deploy, service, err := GetYamlFromTemple(appDeployConfig)

	if err != nil {
		return err
	}

	if conf.Environment.IsDevelopment() {
		fmt.Println(*deploy)
		if len(*service) > 0 {
			fmt.Println("---")
			fmt.Println(*service)
		}
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	deployCfg := appsv1.DeploymentApplyConfiguration{}

	if err := DecodeYaml(deploy, &deployCfg); err != nil {
		fmt.Println(err)
		return err
	}
	if _, e := client.ApplyDeployment(&appDeployConfig.Namespace, &deployCfg); e != nil {
		fmt.Println(e)
		return e
	}

	if len(*service) > 0 {
		serviceCfg := corev1.ServiceApplyConfiguration{}
		if err := DecodeYaml(service, &serviceCfg); err != nil {
			fmt.Println(err)
			return err
		}

		if _, e := client.ApplyService(&appDeployConfig.Namespace, &serviceCfg); e != nil {
			fmt.Println(e)
			return e
		}
	}

	return nil
}
