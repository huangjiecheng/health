package model

type DecisionInfo struct {
	AdjustRatio float64 `json:"adjust_ratio"` // 调度线控线比例
	PlanRatio   float64 `json:"plan_ratio"`   // 规划线控线比例
	AdjustLine  int64   `json:"adjust_line"`  // 决策调度线
	PlanLine    int64   `json:"plan_line"`    // 规划线
}
