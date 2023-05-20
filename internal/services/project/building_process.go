package project

import "time"

type buildingPipeline struct {
	NodeId     uint   // 节点ID
	BuildId    uint   // pipeline_history.ID
	PipelineId uint   // pipeline.ID
	TaskName   string // pod name
	StartAt    time.Time
}

var buildings map[uint][]buildingPipeline // 构建中的列表；key:构建节点ID；value：正在进行中的列表

func init() {
	buildings = make(map[uint][]buildingPipeline)
}
