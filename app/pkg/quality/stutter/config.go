package stutter

import "health/share"

type Config struct {
	Switch           bool                          `json:"switch"`
	DecisionInterval int                           `json:"decision_interval"`
	CoverMap         map[share.CoverName]CoverInfo `json:"cover_map"`
}

type CoverInfo struct {
	SNodeA []*RangeLimit `json:"snode_a"`
	SNodeB []*RangeLimit `json:"snode_b"`
}

type RangeLimit struct {
	LowerLimit             float64 `json:"lowerLimit"`             // 卡顿率下限
	UpperLimit             float64 `json:"upperLimit"`             // 卡顿率上限
	PlanLineControlRate    float64 `json:"planLineControlRate"`    // 规划线的控线比例
	AdjustLineControlRate  float64 `json:"adjustLineControlRate"`  // 调度线的控线比例
	DoDecisionWhenOverLine bool    `json:"doDecisionWhenOverLine"` // 上一分钟平均带宽超过规划线时，是否做决策
}
