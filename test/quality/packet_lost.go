package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type PkgLostConfig struct {
	Switch            bool                       `json:"switch"`
	Period            int                        `json:"period"`
	SilentRoundNumber int                        `json:"silent_round_number"`
	CoverMap          map[string]*PkgLostRateCfg `json:"cover_map"`
}

type PkgLostRateCfg struct {
	AdjustLineRateList []float64
	PlanLineRateList   []float64
	PkgLostRateList    []float64
}

type StrategyType uint8

const (
	DownloadSpeed StrategyType = 1 << iota
	PacketLost
	Arrive
	Stutter
)

const (
	local11  = `{\"arrive_rate\":{\"switch\":true,\"decision_interval\":60,\"cover_map\":{\"filetest\":{\"arrive_rate_list\":[0.93,0.94,0.95,0.96,0.97],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.7,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.6,0.7,0.8,0.9],\"sardine_pass_count\":1000,\"silent_rate\":0.92,\"recovery_rate\":0.98,\"recovery_adjust_line_rate\":1.05},\"zttapk-s\":{\"arrive_rate_list\":[0.6,0.65,0.7,0.75,0.8],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.7,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.6,0.7,0.8,0.9],\"sardine_pass_count\":1000,\"silent_rate\":0.7,\"recovery_rate\":0.8,\"recovery_adjust_line_rate\":1.05},\"zttvi-p\":{\"arrive_rate_list\":[0.7,0.75,0.8,0.85],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.7,0.8,0.9],\"sardine_pass_count\":2000,\"silent_rate\":0.8,\"recovery_rate\":0.85,\"recovery_adjust_line_rate\":1.05}}},\"packet_lost\":{\"switch\":true,\"period\":4,\"silent_round_number\":4,\"cover_map\":{\"default\":{\"AdjustLineRateList\":[1.05,1.04,1.03,1.02,1,0.95,0.95,0.95,0.95,0.9,0.9],\"PlanLineRateList\":[1,1,1,1,1,0.95,0.9,0.85,0.8,0.75,0.7],\"PkgLostRateList\":[0.01,0.02,0.03,0.04,0.05,0.07,0.09,0.11,0.15,0.2,10]},\"zttapk-s\":{\"AdjustLineRateList\":[1.05,1.04,1.03,1.02,1,0.95,0.95,0.95,0.95,0.9,0.9],\"PlanLineRateList\":[1,1,1,1,1,0.95,0.9,0.85,0.8,0.75,0.7],\"PkgLostRateList\":[0.01,0.02,0.03,0.04,0.05,0.07,0.09,0.11,0.15,0.2,10]},\"zttvi-p\":{\"AdjustLineRateList\":[1.05,1.02,1,0.95,0.95,0.95,0.95,0.95,0.95,0.9,0.9,0.9],\"PlanLineRateList\":[1,1,1,0.97,0.94,0.9,0.85,0.8,0.75,0.7,0.65,0.6],\"PkgLostRateList\":[0.01,0.02,0.025,0.03,0.04,0.05,0.07,0.09,0.11,0.15,0.2,10]}}},\"download_speed\":{\"switch\":true,\"sync_interval\":4,\"speed_limit_min_rate\":0.01,\"speed_limit_freeze_time\":120,\"cover_map\":{\"default\":[{\"speed\":1000000,\"rate\":0.05},{\"speed\":2000000,\"rate\":0.85},{\"speed\":3000000,\"rate\":0.9},{\"speed\":4000000,\"rate\":0.95},{\"speed\":5000000,\"rate\":1},{\"speed\":6000000,\"rate\":1},{\"speed\":7000000,\"rate\":1}],\"faiqiyi\":[{\"speed\":3000000,\"rate\":0.05},{\"speed\":5000000,\"rate\":0.1},{\"speed\":6000000,\"rate\":0.4},{\"speed\":7000000,\"rate\":0.6},{\"speed\":10000000,\"rate\":0.7},{\"speed\":16000000,\"rate\":0.8},{\"speed\":20000000,\"rate\":0.9}],\"filetest\":[{\"speed\":1000000,\"rate\":0.05},{\"speed\":2000000,\"rate\":0.85},{\"speed\":3000000,\"rate\":0.9},{\"speed\":4000000,\"rate\":0.95},{\"speed\":5000000,\"rate\":1},{\"speed\":6000000,\"rate\":1},{\"speed\":7000000,\"rate\":1}],\"pf001\":[{\"speed\":1000000,\"rate\":0.05},{\"speed\":2000000,\"rate\":0.85},{\"speed\":3000000,\"rate\":0.9},{\"speed\":4000000,\"rate\":0.95},{\"speed\":5000000,\"rate\":1},{\"speed\":6000000,\"rate\":1},{\"speed\":7000000,\"rate\":1}],\"pf003\":[{\"speed\":1000000,\"rate\":0.05},{\"speed\":2000000,\"rate\":0.85},{\"speed\":3000000,\"rate\":0.9},{\"speed\":4000000,\"rate\":0.95},{\"speed\":5000000,\"rate\":1},{\"speed\":6000000,\"rate\":1},{\"speed\":7000000,\"rate\":1}],\"zqctgdl\":[{\"speed\":1000000,\"rate\":0.05},{\"speed\":2000000,\"rate\":0.85},{\"speed\":3000000,\"rate\":0.9},{\"speed\":4000000,\"rate\":0.95},{\"speed\":5000000,\"rate\":1},{\"speed\":6000000,\"rate\":1},{\"speed\":7000000,\"rate\":1}],\"zttapk-s\":[{\"speed\":1000000,\"rate\":0.05},{\"speed\":2000000,\"rate\":0.85},{\"speed\":3000000,\"rate\":0.9},{\"speed\":4000000,\"rate\":0.95},{\"speed\":5000000,\"rate\":1},{\"speed\":6000000,\"rate\":1},{\"speed\":7000000,\"rate\":1}],\"zttvi-p\":[{\"speed\":3000000,\"rate\":0.05},{\"speed\":4000000,\"rate\":0.85},{\"speed\":5000000,\"rate\":0.9},{\"speed\":6000000,\"rate\":0.95},{\"speed\":7000000,\"rate\":1},{\"speed\":8000000,\"rate\":1},{\"speed\":9000000,\"rate\":1}]}}}`
	center11 = `{\"arrive_rate\":{\"switch\":true,\"decision_interval\":60,\"cover_map\":{\"default\":{\"arrive_rate_list\":[0.6,0.65,0.7,0.75,0.8],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.7,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.6,0.7,0.8,0.9],\"sardine_pass_count\":1000,\"silent_rate\":0.7,\"recovery_rate\":0.8,\"recovery_adjust_line_rate\":1.05},\"filetest\":{\"arrive_rate_list\":[0.93,0.94,0.95,0.96,0.97],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.7,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.6,0.7,0.8,0.9],\"sardine_pass_count\":1000,\"silent_rate\":0.92,\"recovery_rate\":0.98,\"recovery_adjust_line_rate\":1.05},\"zttapk-s\":{\"arrive_rate_list\":[0.6,0.65,0.7,0.75,0.8],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.7,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.6,0.7,0.8,0.9],\"sardine_pass_count\":1000,\"silent_rate\":0.7,\"recovery_rate\":0.8,\"recovery_adjust_line_rate\":1.05},\"zttvi-p\":{\"arrive_rate_list\":[0.7,0.75,0.8,0.85],\"lt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.95,1],\"gt_avg_bw_adjust_line_rate_list\":[0.85,0.9,0.95,1],\"lt_avg_bw_plan_line_rate_list\":[0.6,0.8,0.9,1],\"gt_avg_bw_plan_line_rate_list\":[0.5,0.7,0.8,0.9],\"sardine_pass_count\":2000,\"silent_rate\":0.8,\"recovery_rate\":0.85,\"recovery_adjust_line_rate\":1.05}}},\"packet_lost\":{\"switch\":true,\"period\":4,\"silent_round_number\":4,\"cover_map\":{\"default\":{\"AdjustLineRateList\":[1.05,1.04,1.03,1.02,1,0.95,0.95,0.95,0.95,0.9,0.9],\"PlanLineRateList\":[1,1,1,1,1,0.95,0.9,0.85,0.8,0.75,0.7],\"PkgLostRateList\":[0.01,0.02,0.03,0.04,0.05,0.07,0.09,0.11,0.15,0.2,10]},\"zttapk-s\":{\"AdjustLineRateList\":[1.05,1.04,1.03,1.02,1,0.95,0.95,0.95,0.95,0.9,0.9],\"PlanLineRateList\":[1,1,1,1,1,0.95,0.9,0.85,0.8,0.75,0.7],\"PkgLostRateList\":[0.01,0.02,0.03,0.04,0.05,0.07,0.09,0.11,0.15,0.2,10]},\"zttvi-p\":{\"AdjustLineRateList\":[1.05,1.02,1,0.95,0.95,0.95,0.95,0.95,0.95,0.9,0.9,0.9],\"PlanLineRateList\":[1,1,1,0.97,0.94,0.9,0.85,0.8,0.75,0.7,0.65,0.6],\"PkgLostRateList\":[0.01,0.02,0.025,0.03,0.04,0.05,0.07,0.09,0.11,0.15,0.2,10]}}},\"download_speed\":{\"switch\":true,\"sync_interval\":4,\"speed_limit_min_rate\":0.01,\"speed_limit_freeze_time\":120,\"cover_map\":{\"default\":[{\"speed\":7,\"rate\":1},{\"speed\":6,\"rate\":1},{\"speed\":5,\"rate\":1},{\"speed\":4,\"rate\":0.95},{\"speed\":3,\"rate\":0.9},{\"speed\":2,\"rate\":0.85},{\"speed\":1,\"rate\":0.05}],\"faiqiyi\":[{\"speed\":20,\"rate\":0.9},{\"speed\":16,\"rate\":0.8},{\"speed\":10,\"rate\":0.7},{\"speed\":7,\"rate\":0.6},{\"speed\":6,\"rate\":0.4},{\"speed\":5,\"rate\":0.1},{\"speed\":3,\"rate\":0.05}],\"filetest\":[{\"speed\":7,\"rate\":1},{\"speed\":6,\"rate\":1},{\"speed\":5,\"rate\":1},{\"speed\":4,\"rate\":0.95},{\"speed\":3,\"rate\":0.9},{\"speed\":2,\"rate\":0.85},{\"speed\":1,\"rate\":0.05}],\"pf001\":[{\"speed\":7,\"rate\":1},{\"speed\":6,\"rate\":1},{\"speed\":5,\"rate\":1},{\"speed\":4,\"rate\":0.95},{\"speed\":3,\"rate\":0.9},{\"speed\":2,\"rate\":0.85},{\"speed\":1,\"rate\":0.05}],\"zqctgdl\":[{\"speed\":7,\"rate\":1},{\"speed\":6,\"rate\":1},{\"speed\":5,\"rate\":1},{\"speed\":4,\"rate\":0.95},{\"speed\":3,\"rate\":0.9},{\"speed\":2,\"rate\":0.85},{\"speed\":1,\"rate\":0.05}],\"zttapk-s\":[{\"speed\":7,\"rate\":1},{\"speed\":6,\"rate\":1},{\"speed\":5,\"rate\":1},{\"speed\":4,\"rate\":0.95},{\"speed\":3,\"rate\":0.9},{\"speed\":2,\"rate\":0.85},{\"speed\":1,\"rate\":0.05}],\"zttvi-p\":[{\"speed\":9,\"rate\":1},{\"speed\":8,\"rate\":1},{\"speed\":7,\"rate\":1},{\"speed\":6,\"rate\":0.95},{\"speed\":5,\"rate\":0.9},{\"speed\":4,\"rate\":0.85},{\"speed\":3,\"rate\":0.05}]}}}`
)

type EnumType struct {
	FinalStrategy StrategyType `json:"final_strategy"` // 最终控线策略
}

// AlarmInfo 每个报警的内容
type AlarmInfo struct {
	Endpoint string                 `json:"endpoint"` // 主机名 如果传空，监控组件会取当前机器的主机名
	Step     int64                  `json:"step"`     // 采集时间步长 如果为空，默认取 60 单位 秒
	Value    int64                  `json:"value"`    // 默认value 当 fields 只有一个字段的时候，推荐只使用该值
	Tags     map[string]interface{} `json:"tags"`     // 用于标识报警的唯一性
	Fields   map[string]interface{} `json:"fields"`   // 主要数据
	Time     int64                  `json:"time"`     // 秒级时间戳 如果为空, 默认取当前时间
	Name     string                 `json:"name"`     // 监控数据表名, 必传
}

func Errorrrrr(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v...))
}

type ReportV2 struct {
	TTL  int
	Name string
}

func main() {
	r := &ReportV2{
		Name: "111",
		TTL:  111,
	}
	tempReport := *r
	cacheReport := &tempReport

	tempReport2 := *r
	cacheReport2 := &tempReport2
	cacheReport2.Name = "2222"
	fmt.Println(cacheReport)
	expectWeight := int64(8800)
	dValue := int64(2202)
	allowAdjustExpectWeightRate := 0.5
	rrr := expectWeight < int64(allowAdjustExpectWeightRate*float64(dValue))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), rrr)
	aaa := EnumType{
		FinalStrategy: Stutter,
	}
	fmt.Println("0000000000001" + strconv.Itoa(int(aaa.FinalStrategy)))
	finalMap := make(map[string]*PkgLostRateCfg)
	key, conf := defaultLost()
	finalMap[key] = conf
	key, conf = defaultzttvi_p()
	finalMap[key] = conf
	key, conf = defaultzttapk_s()
	finalMap[key] = conf
	input := &PkgLostConfig{
		Switch:            true,
		Period:            4,
		SilentRoundNumber: 4,
		CoverMap:          finalMap,
	}
	//tempReport := *input
	//cacheReport := &tempReport
	//cacheReport.Switch = false
	data, err := json.Marshal(input)
	fmt.Println(string(data), err)
}

func defaultLost() (string, *PkgLostRateCfg) {
	cover := "default"
	result := &PkgLostRateCfg{
		AdjustLineRateList: []float64{
			1.05, 1.04, 1.03, 1.02, 1.0, 0.95, 0.95, 0.95, 0.95, 0.90, 0.90,
		},
		PlanLineRateList: []float64{
			1.0, 1.0, 1.0, 1.0, 1.0, 0.95, 0.90, 0.85, 0.80, 0.75, 0.70,
		},
		PkgLostRateList: []float64{
			0.01, 0.02, 0.03, 0.04, 0.05, 0.07, 0.09, 0.11, 0.15, 0.2, 10.0,
		},
	}
	return cover, result
}

func defaultzttvi_p() (string, *PkgLostRateCfg) {
	cover := "zttvi-p"
	result := &PkgLostRateCfg{
		AdjustLineRateList: []float64{
			1.05, 1.02, 1.0, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95, 0.90, 0.90, 0.90,
		},
		PlanLineRateList: []float64{
			1.0, 1.0, 1.0, 0.97, 0.94, 0.90, 0.85, 0.80, 0.75, 0.70, 0.65, 0.60,
		},
		PkgLostRateList: []float64{
			0.01, 0.02, 0.025, 0.03, 0.04, 0.05, 0.07, 0.09, 0.11, 0.15, 0.2, 10.0,
		},
	}
	return cover, result
}

func defaultzttapk_s() (string, *PkgLostRateCfg) {
	cover := "zttapk-s"
	result := &PkgLostRateCfg{
		AdjustLineRateList: []float64{
			1.05, 1.04, 1.03, 1.02, 1.0, 0.95, 0.95, 0.95, 0.95, 0.90, 0.90,
		},
		PlanLineRateList: []float64{
			1.0, 1.0, 1.0, 1.0, 1.0, 0.95, 0.90, 0.85, 0.80, 0.75, 0.70,
		},
		PkgLostRateList: []float64{
			0.01, 0.02, 0.03, 0.04, 0.05, 0.07, 0.09, 0.11, 0.15, 0.2, 10.0,
		},
	}
	return cover, result
}

//func arriveZttvi_p() (string, ArriveConfig) {
//	cover := "zttvi-p"
//	return cover, ArriveConfig{
//		//arrive_rate = 0.7
//		//arrive_rate = 0.75
//		//arrive_rate = 0.8
//		//arrive_rate = 0.85
//		ArriveRateList: []float64{0.7, 0.75, 0.8, 0.85},
//		//#小于平均带宽的调度线
//		//lt_adjust_line_rate = 0.85
//		//lt_adjust_line_rate = 0.9
//		//lt_adjust_line_rate = 0.95
//		//lt_adjust_line_rate = 1
//		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.95, 1},
//		//#大于平均带宽的调度线
//		//gt_adjust_line_rate = 0.85
//		//gt_adjust_line_rate = 0.9
//		//gt_adjust_line_rate = 0.95
//		//gt_adjust_line_rate = 1
//		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.95, 1},
//		//#小于平均带宽的规划线
//		//lt_plan_line_rate = 0.6
//		//lt_plan_line_rate = 0.8
//		//lt_plan_line_rate = 0.9
//		//lt_plan_line_rate = 1
//		LtAvgBwPlanLineRateList: []float64{0.6, 0.8, 0.9, 1},
//		//#大于平均带宽的规划线
//		//gt_plan_line_rate = 0.5
//		//gt_plan_line_rate = 0.7
//		//gt_plan_line_rate = 0.8
//		//gt_plan_line_rate = 0.9
//		GtAvgBwPlanLineRateList: []float64{0.5, 0.7, 0.8, 0.9},
//		//#sardine放行总数
//		//sardine_pass_count = 2000
//		//#静默
//		//silent_rate = 0.8
//		//#恢复
//		//recovery_rate = 0.85
//		//#恢复后的调度线
//		//recovery_adjust_line_rate = 1.05
//		SardinePassCount:       2000,
//		SilentRate:             0.8,
//		RecoveryRate:           0.85,
//		RecoveryAdjustLineRate: 1.05,
//	}
//}
//
//func arrivezttapk_s() (string, ArriveConfig) {
//	cover := "zttapk-s"
//	return cover, ArriveConfig{
//		//#到达率
//		//arrive_rate = 0.93
//		//arrive_rate = 0.94
//		//arrive_rate = 0.95
//		//arrive_rate = 0.96
//		//arrive_rate = 0.97
//		ArriveRateList: []float64{0.93, 0.94, 0.95, 0.96, 0.97},
//		//#小于平均带宽的调度线
//		//lt_adjust_line_rate = 0.85
//		//lt_adjust_line_rate = 0.9
//		//lt_adjust_line_rate = 0.9
//		//lt_adjust_line_rate = 0.95
//		//lt_adjust_line_rate = 1
//		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
//		//#大于平均带宽的调度线
//		//gt_adjust_line_rate = 0.85
//		//gt_adjust_line_rate = 0.9
//		//gt_adjust_line_rate = 0.9
//		//gt_adjust_line_rate = 0.95
//		//gt_adjust_line_rate = 1
//		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
//		//#小于平均带宽的规划线
//		//lt_plan_line_rate = 0.6
//		//lt_plan_line_rate = 0.7
//		//lt_plan_line_rate = 0.8
//		//lt_plan_line_rate = 0.9
//		//lt_plan_line_rate = 1
//		LtAvgBwPlanLineRateList: []float64{0.6, 0.7, 0.8, 0.9, 1},
//		//#大于平均带宽的规划线
//		//gt_plan_line_rate = 0.5
//		//gt_plan_line_rate = 0.6
//		//gt_plan_line_rate = 0.7
//		//gt_plan_line_rate = 0.8
//		//gt_plan_line_rate = 0.9
//		GtAvgBwPlanLineRateList: []float64{0.5, 0.6, 0.7, 0.8, 0.9},
//		//#sardine放行总数
//		//sardine_pass_count = 1000
//		//#静默
//		//silent_rate = 0.92
//		//#恢复
//		//recovery_rate = 0.98
//		//#恢复后的调度线
//		//recovery_adjust_line_rate = 1.05
//		SardinePassCount:       1000,
//		SilentRate:             0.92,
//		RecoveryRate:           0.98,
//		RecoveryAdjustLineRate: 1.05,
//	}
//}
//
//func arrivefiletest() (string, ArriveConfig) {
//
//	//#小于平均带宽的调度线
//	//lt_adjust_line_rate = 0.85
//	//lt_adjust_line_rate = 0.9
//	//lt_adjust_line_rate = 0.9
//	//lt_adjust_line_rate = 0.95
//	//lt_adjust_line_rate = 1
//	//#大于平均带宽的调度线
//	//gt_adjust_line_rate = 0.85
//	//gt_adjust_line_rate = 0.9
//	//gt_adjust_line_rate = 0.9
//	//gt_adjust_line_rate = 0.95
//	//gt_adjust_line_rate = 1
//	//#sardine放行总数
//	//sardine_pass_count = 1000
//	//#静默
//	//silent_rate = 0.7
//	//#恢复
//	//recovery_rate = 0.8
//	//#恢复后的调度线
//	//recovery_adjust_line_rate = 1.05
//	cover := "filetest"
//	return cover, ArriveConfig{
//		//#到达率
//		//arrive_rate = 0.6
//		//arrive_rate = 0.65
//		//arrive_rate = 0.7
//		//arrive_rate = 0.75
//		//arrive_rate = 0.8
//		ArriveRateList: []float64{0.6, 0.65, 0.7, 0.75, 0.8},
//		//#小于平均带宽的调度线
//		//lt_adjust_line_rate = 0.85
//		//lt_adjust_line_rate = 0.9
//		//lt_adjust_line_rate = 0.9
//		//lt_adjust_line_rate = 0.95
//		//lt_adjust_line_rate = 1
//		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
//		//#大于平均带宽的调度线
//		//gt_adjust_line_rate = 0.85
//		//gt_adjust_line_rate = 0.9
//		//gt_adjust_line_rate = 0.9
//		//gt_adjust_line_rate = 0.95
//		//gt_adjust_line_rate = 1
//		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
//		//#小于平均带宽的规划线
//		//lt_plan_line_rate = 0.6
//		//lt_plan_line_rate = 0.7
//		//lt_plan_line_rate = 0.8
//		//lt_plan_line_rate = 0.9
//		//lt_plan_line_rate = 1
//		LtAvgBwPlanLineRateList: []float64{0.6, 0.7, 0.8, 0.9, 1},
//		//#大于平均带宽的规划线
//		//gt_plan_line_rate = 0.5
//		//gt_plan_line_rate = 0.6
//		//gt_plan_line_rate = 0.7
//		//gt_plan_line_rate = 0.8
//		//gt_plan_line_rate = 0.9
//		GtAvgBwPlanLineRateList: []float64{0.5, 0.6, 0.7, 0.8, 0.9},
//		//#sardine放行总数
//		//sardine_pass_count = 1000
//		//#静默
//		//silent_rate = 0.7
//		//#恢复
//		//recovery_rate = 0.8
//		//#恢复后的调度线
//		//recovery_adjust_line_rate = 1.05
//		SardinePassCount:       1000,
//		SilentRate:             0.7,
//		RecoveryRate:           0.8,
//		RecoveryAdjustLineRate: 1.05,
//	}
//}
