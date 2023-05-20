package kube

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"strings"
	"time"
)

const DeploymentLabelSelector string = "gotocloud"

var deploymentCache *cache.Cache

func init() {
	deploymentCache = cache.New(time.Second*5, 0)
}

// GetDeployments 获取部署工作负载
func (client *Client) GetDeployments(ctx context.Context, k8sId uint, ns string, deploymentsId *map[uint]bool, force bool) ([]deploy.DeploymentDescription, error) {
	var rlt []deploy.DeploymentDescription

	if tmp, ok := deploymentCache.Get(fmt.Sprintf("%d.%s", k8sId, ns)); !ok || force {
		d, err := client.clientSet.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{
			LabelSelector: fmt.Sprintf("deployed=%s", DeploymentLabelSelector),
		})
		if err != nil {
			return nil, err
		}
		rlt = make([]deploy.DeploymentDescription, len(d.Items))
		for i, item := range d.Items {
			rlt[i] = deploy.DeploymentDescription{
				Id: func() uint {
					s := item.GetLabels()["appId"]
					idStr := strings.TrimPrefix(s, fmt.Sprintf("%s-", DeploymentLabelSelector))
					if id, err := strconv.ParseUint(idStr, 0, 64); err != nil {
						return 0
					} else {
						return uint(id)
					}
				}(),
				Name:            item.Name,
				Namespace:       item.Namespace,
				Replicas:        uint(*item.Spec.Replicas),
				AvailablePods:   uint(item.Status.AvailableReplicas),
				UnavailablePods: uint(item.Status.UnavailableReplicas),
				CreatedAt:       utils.JsonTime(item.ObjectMeta.GetCreationTimestamp().Time),
				Conditions: func() []deploy.DeploymentCondition {
					cond := make([]deploy.DeploymentCondition, len(item.Status.Conditions))
					for i2, condition := range item.Status.Conditions {
						cond[i2] = deploy.DeploymentCondition{
							Type:    string(condition.Type),
							Status:  deploy.ConditionStatus(condition.Status),
							Message: condition.Message,
						}
					}
					return cond
				}(),
			}
		}

		deploymentCache.Set(fmt.Sprintf("%d.%s", k8sId, ns), rlt, 0)
	} else {
		rlt = tmp.([]deploy.DeploymentDescription)
	}

	if deploymentsId == nil {
		return rlt, nil
	} else {
		found := make([]deploy.DeploymentDescription, 0)
		for i, description := range rlt {
			if (*deploymentsId)[description.Id] {
				found = append(found, rlt[i])
			}
		}
		return found, nil
	}
}
