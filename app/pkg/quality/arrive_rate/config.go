package arrive_rate

import "health/share"

type Config struct {
	Switch           bool                            `json:"switch"`
	DecisionInterval int                             `json:"decision_interval"`
	CoverMap         map[share.CoverName]*RangeLimit `json:"cover_map"`
}

// RangeLimit 到达率控线策略配置
type RangeLimit struct {
	ArriveRateList            []float64 `json:"arrive_rate_list"`
	LtAvgBwAdjustLineRateList []float64 `json:"lt_avg_bw_adjust_line_rate_list"`
	GtAvgBwAdjustLineRateList []float64 `json:"gt_avg_bw_adjust_line_rate_list"`
	LtAvgBwPlanLineRateList   []float64 `json:"lt_avg_bw_plan_line_rate_list"`
	GtAvgBwPlanLineRateList   []float64 `json:"gt_avg_bw_plan_line_rate_list"`
	SardinePassCount          int64     `json:"sardine_pass_count"`
	SilentRate                float64   `json:"silent_rate"`
	RecoveryRate              float64   `json:"recovery_rate"`
	RecoveryAdjustLineRate    float64   `json:"recovery_adjust_line_rate"`
}
