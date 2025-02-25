package core

import (
	"github.com/skyhackvip/risk_engine/internal/log"
)

type EndNode struct {
	Info NodeInfo
}

func NewEndNode(name string) *EndNode {
	return &EndNode{
		Info: NodeInfo{Name: name, Kind: TypeEnd.String()},
	}
}

func (node EndNode) GetName() string {
	return node.Info.Name
}

func (node EndNode) GetType() NodeType {
	return GetNodeType(node.Info.Kind)
}

func (node EndNode) GetInfo() NodeInfo {
	return node.Info
}

func (node EndNode) BeforeParse(ctx *PipelineContext) error {
	return nil
}

func (node EndNode) AfterParse(ctx *PipelineContext, result *NodeResult) error {
	return nil
}

func (node EndNode) Parse(ctx *PipelineContext) (*NodeResult, error) {
	log.Info("======[trace] End======")
	info := node.GetInfo()
	nodeResult := &NodeResult{Id: info.Id, Name: info.Name, Kind: node.GetType(), Tag: info.Tag, Label: info.Label, IsBlock: true}
	return nodeResult, nil
}
