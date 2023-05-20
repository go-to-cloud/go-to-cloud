package deploy

import (
	"go-to-cloud/internal/utils"
)

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type DeploymentCondition struct {
	Type    string          `json:"type,omitempty"`
	Status  ConditionStatus `json:"status"`
	Message string          `json:"message"`
}

// DeploymentDescription 裁剪后的Deployment Spec
type DeploymentDescription struct {
	Id              uint                  `json:"id"`              // deployment.ID
	Name            string                `json:"name"`            // 应用名称
	Namespace       string                `json:"namespace"`       // 名字空间
	Replicas        uint                  `json:"replicas"`        // 副本数
	AvailablePods   uint                  `json:"availablePods"`   // 可用副本数
	UnavailablePods uint                  `json:"unavailablePods"` // 不可用副本数
	CreatedAt       utils.JsonTime        `json:"createdAt"`       // 创建时间
	Conditions      []DeploymentCondition `json:"conditions"`      // 状态
}
