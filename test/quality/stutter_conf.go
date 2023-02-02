package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type StutterConfig struct {
	Switch           bool                 `json:"switch"`
	DecisionInterval int                  `json:"decision_interval"`
	CoverMap         map[string]CoverInfo `json:"cover_map"`
}

type CoverInfo struct {
	SNodeA []RangeLimit `json:"snode_a"`
	SNodeB []RangeLimit `json:"snode_b"`
}

type RangeLimit struct {
	LowerLimit             float64 `json:"lowerLimit"`             // 卡顿率下限
	UpperLimit             float64 `json:"upperLimit"`             // 卡顿率上限
	PlanLineControlRate    float64 `json:"planLineControlRate"`    // 规划线的控线比例
	AdjustLineControlRate  float64 `json:"adjustLineControlRate"`  // 调度线的控线比例
	DoDecisionWhenOverLine bool    `json:"doDecisionWhenOverLine"` // 上一分钟平均带宽超过规划线时，是否做决策
}

func main() {
	aa, bb := 11, 22
	if aa == 11 {
		aa, bb = 33, 44
	}
	fmt.Println(aa, bb)
	//zttvi_p := "zttvi-p"
	//zttapk_s := "zttapk-s"
	//default_ := "default"
	//dynamic_stutter_range = 0.00,0.01,1.00,1.06,0
	//dynamic_stutter_range = 0.01,0.03,1.00,1.04,0
	//dynamic_stutter_range = 0.03,0.05,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.05,0.07,1.00,1.00,0
	//dynamic_stutter_range = 0.07,0.09,0.95,0.95,1
	//dynamic_stutter_range = 0.09,0.11,0.90,0.95,1
	//dynamic_stutter_range = 0.11,0.13,0.85,0.95,1
	//dynamic_stutter_range = 0.13,0.15,0.80,0.90,1
	//dynamic_stutter_range = 0.15,1.00,0.75,0.90,1
	finalMap := make(map[string]CoverInfo)
	key, coverInfo := translateNew(getDefaultCoverConf())
	finalMap[key] = coverInfo
	key, coverInfo = translateNew(getZttvi_p())
	finalMap[key] = coverInfo
	key, coverInfo = translateNew(getZttapk_s())
	finalMap[key] = coverInfo
	input := &StutterConfig{
		Switch:           true,
		DecisionInterval: 60,
		CoverMap:         finalMap,
	}
	data, err := json.Marshal(input)
	fmt.Println(string(data), err)
}

func getDefaultCoverConf() ([]string, []string, string) {
	//dynamic_stutter_range = 0.00,0.05,1.00,1.06,0
	//dynamic_stutter_range = 0.05,0.07,1.00,1.04,0
	//dynamic_stutter_range = 0.07,0.09,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.09,0.11,1.00,1.00,0
	//dynamic_stutter_range = 0.11,0.13,0.95,0.95,1
	//dynamic_stutter_range = 0.13,0.15,0.90,0.95,1
	//dynamic_stutter_range = 0.15,0.20,0.85,0.90,1
	//dynamic_stutter_range = 0.20,0.25,0.80,0.90,1
	//dynamic_stutter_range = 0.25,0.30,0.75,0.85,1
	//dynamic_stutter_range = 0.30,1.00,0.70,0.85,1
	arrayA := []string{
		"0.00,0.05,1.00,1.06,0",
		"0.05,0.07,1.00,1.04,0",
		"0.07,0.09,1.00,1.02,0",
		"0.11,0.13,0.95,0.95,1",
		"0.13,0.15,0.90,0.95,1",
		"0.15,0.20,0.85,0.90,1",
		"0.20,0.25,0.80,0.90,1",
		"0.25,0.30,0.75,0.85,1",
		"0.30,1.00,0.70,0.85,1",
	}
	//dynamic_stutter_range = 0.00,0.06,1.00,1.06,0
	//dynamic_stutter_range = 0.06,0.08,1.00,1.04,0
	//dynamic_stutter_range = 0.08,0.10,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.10,0.12,1.00,1.00,0
	//dynamic_stutter_range = 0.12,0.14,0.95,0.95,1
	//dynamic_stutter_range = 0.14,0.16,0.90,0.95,1
	//dynamic_stutter_range = 0.16,0.19,0.85,0.90,1
	//dynamic_stutter_range = 0.19,0.22,0.80,0.90,1
	//dynamic_stutter_range = 0.22,0.25,0.75,0.85,1
	//dynamic_stutter_range = 0.25,1.00,0.70,0.85,1
	arrayB := []string{
		"0.00,0.06,1.00,1.06,0",
		"0.06,0.08,1.00,1.04,0",
		"0.08,0.10,1.00,1.02,0",
		"0.12,0.14,0.95,0.95,1",
		"0.14,0.16,0.90,0.95,1",
		"0.16,0.19,0.85,0.90,1",
		"0.19,0.22,0.80,0.90,1",
		"0.22,0.25,0.75,0.85,1",
		"0.25,1.00,0.70,0.85,1",
	}
	return arrayA, arrayB, "default"
}

func translateNew(arrayA, arrayB []string, cover string) (string, CoverInfo) {
	snodeA := make([]RangeLimit, 0)
	snodeB := make([]RangeLimit, 0)
	snodeA = translateArray(arrayA)
	snodeB = translateArray(arrayB)
	return cover, CoverInfo{
		SNodeA: snodeA,
		SNodeB: snodeB,
	}
}
func getZttapk_s() ([]string, []string, string) {
	//dynamic_stutter_range = 0.00,0.05,1.00,1.06,0
	//dynamic_stutter_range = 0.05,0.07,1.00,1.04,0
	//dynamic_stutter_range = 0.07,0.09,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.09,0.11,1.00,1.00,0
	//dynamic_stutter_range = 0.11,0.13,0.95,0.95,1
	//dynamic_stutter_range = 0.13,0.15,0.90,0.95,1
	//dynamic_stutter_range = 0.15,0.17,0.85,0.90,1
	//dynamic_stutter_range = 0.17,0.20,0.80,0.90,1
	//dynamic_stutter_range = 0.20,1.00,0.75,0.85,1
	arrayA := []string{
		"0.00,0.05,1.00,1.06,0",
		"0.05,0.07,1.00,1.04,0",
		"0.07,0.09,1.00,1.02,0",
		"0.11,0.13,0.95,0.95,1",
		"0.13,0.15,0.90,0.95,1",
		"0.15,0.17,0.85,0.90,1",
		"0.17,0.20,0.80,0.90,1",
		"0.20,1.00,0.75,0.85,1",
	}
	//dynamic_stutter_range = 0.00,0.01,1.00,1.06,0
	//dynamic_stutter_range = 0.01,0.03,1.00,1.04,0
	//dynamic_stutter_range = 0.03,0.05,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.05,0.07,1.00,1.00,0
	//dynamic_stutter_range = 0.07,0.09,0.95,0.95,1
	//dynamic_stutter_range = 0.09,0.11,0.90,0.95,1
	//dynamic_stutter_range = 0.11,0.13,0.85,0.90,1
	//dynamic_stutter_range = 0.13,0.15,0.80,0.90,1
	//dynamic_stutter_range = 0.15,1.00,0.75,0.85,1
	arrayB := []string{
		"0.00,0.01,1.00,1.06,0",
		"0.01,0.03,1.00,1.04,0",
		"0.03,0.05,1.00,1.02,0",
		"0.07,0.09,0.95,0.95,1",
		"0.09,0.11,0.90,0.95,1",
		"0.11,0.13,0.85,0.90,1",
		"0.13,0.15,0.80,0.90,1",
		"0.15,1.00,0.75,0.85,1",
	}
	return arrayA, arrayB, "zttapk-s"
}
func getZttvi_p() ([]string, []string, string) {
	//dynamic_stutter_range = 0.00,0.01,1.00,1.06,0
	//dynamic_stutter_range = 0.01,0.03,1.00,1.04,0
	//dynamic_stutter_range = 0.03,0.05,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.05,0.07,1.00,1.00,0
	//dynamic_stutter_range = 0.07,0.09,0.95,0.95,1
	//dynamic_stutter_range = 0.09,0.11,0.90,0.95,1
	//dynamic_stutter_range = 0.11,0.13,0.85,0.95,1
	//dynamic_stutter_range = 0.13,0.15,0.80,0.90,1
	//dynamic_stutter_range = 0.15,1.00,0.75,0.90,1
	arrayA := []string{
		"0.00,0.01,1.00,1.06,0",
		"0.01,0.03,1.00,1.04,0",
		"0.03,0.05,1.00,1.02,0",
		"0.07,0.09,0.95,0.95,1",
		"0.09,0.11,0.90,0.95,1",
		"0.11,0.13,0.85,0.95,1",
		"0.13,0.15,0.80,0.90,1",
		"0.15,1.00,0.75,0.90,1",
	}
	//dynamic_stutter_range = 0.00,0.01,1.00,1.06,0
	//dynamic_stutter_range = 0.01,0.02,1.00,1.04,0
	//dynamic_stutter_range = 0.02,0.03,1.00,1.02,0
	//#不处理 dynamic_stutter_range = 0.03,0.04,1.00,1.00,0
	//dynamic_stutter_range = 0.04,0.05,0.95,0.95,1
	//dynamic_stutter_range = 0.05,0.06,0.90,0.95,1
	//dynamic_stutter_range = 0.06,0.07,0.85,0.95,1
	//dynamic_stutter_range = 0.07,0.08,0.80,0.90,1
	//dynamic_stutter_range = 0.08,1.00,0.75,0.90,1
	arrayB := []string{
		"0.00,0.01,1.00,1.06,0",
		"0.01,0.02,1.00,1.04,0",
		"0.02,0.03,1.00,1.02,0",
		"0.04,0.05,0.95,0.95,1",
		"0.05,0.06,0.90,0.95,1",
		"0.06,0.07,0.85,0.95,1",
		"0.07,0.08,0.80,0.90,1",
		"0.08,1.00,0.75,0.90,1",
	}
	return arrayA, arrayB, "zttvi-p"
}

func translateArray(arr []string) []RangeLimit {
	result := make([]RangeLimit, 0)
	for _, s := range arr {
		nums := strings.Split(s, ",")
		// 解析参数值
		lowerLimit, _ := strconv.ParseFloat(nums[0], 64)
		UpperLimit, _ := strconv.ParseFloat(nums[1], 64)
		PlanLineRate, _ := strconv.ParseFloat(nums[2], 64)
		AdjustLineRate, _ := strconv.ParseFloat(nums[3], 64)
		doDecisionWhenOverLine, _ := strconv.ParseInt(nums[4], 10, 64)
		result = append(result, RangeLimit{
			LowerLimit:             lowerLimit,
			UpperLimit:             UpperLimit,
			PlanLineControlRate:    PlanLineRate,
			AdjustLineControlRate:  AdjustLineRate,
			DoDecisionWhenOverLine: doDecisionWhenOverLine > 0,
		})
	}
	return result
}
