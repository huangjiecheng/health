package main

import (
	"fmt"
	"sort"
)

type AffinityInfo struct {
	CardIp       string
	IpCard       string
	UsedWeight   int64
	Priority     int64
	LastUsedTime int64
}
type CacheGroupInfo struct {
	Name             string
	ExpectWeight     int64
	ExpectWeightRate int64
	Qps              int64
	AffinityInfo     *AffinityInfo
}

var (
	//GlobalConfig = &CacheGroupInfo{} // 通过同步配置中心得到质量决策配置
	GlobalConfig *CacheGroupInfo // 通过同步配置中心得到质量决策配置
)

func main() {
	//GlobalConfig = &CacheGroupInfo{}
	GlobalConfig.AffinityInfo = &AffinityInfo{}

	cacheRatio := float64(714) / float64(4286)
	demoNum := int64(100 * 1000 * cacheRatio)
	demoNum2 := 50 * demoNum / 1000
	//q := float64(714) / float64(4286)
	//demoNum := int64(100 * 1000 * q)
	fmt.Println("adasdasdsadasdddd", demoNum, demoNum2)
	//{"viewUnit.Weight": 100, "reqNum": 714, "totalReqByView": 4286, "info.Ratio": 0}

	result := make([]*CacheGroupInfo, 0, 3)
	result = append(result, &CacheGroupInfo{
		Name:         "dx-chongqing-chongqing-17-cache-2",
		ExpectWeight: 2000,
	})
	result = append(result, &CacheGroupInfo{
		Name:         "dx-neimenggu-huhehaote-9-cache-1",
		ExpectWeight: 5000,
	})
	result = append(result, &CacheGroupInfo{
		Name:         "dx-hubei-yichang-7-cache-3",
		ExpectWeight: 3000,
	})
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	newAffinityMgr := make(map[string][]AffinityInfo, 3)
	newAffinityMgr["dx-chongqing-chongqing-17-cache-2"] = []AffinityInfo{
		{CardIp: "114.101.81.116", IpCard: "ppp17627012", UsedWeight: 0, Priority: 1, LastUsedTime: 1666769633},
		{CardIp: "114.101.82.214", IpCard: "ppp17627034", UsedWeight: 0, Priority: 4, LastUsedTime: 1666769649},
		{CardIp: "114.101.82.194", IpCard: "ppp17627032", UsedWeight: 0, Priority: 2, LastUsedTime: 1666769649},
		{CardIp: "114.101.80.218", IpCard: "ppp17627030", UsedWeight: 5195700, Priority: 1, LastUsedTime: 1666769653},
		{CardIp: "114.101.82.118", IpCard: "ppp17627010", UsedWeight: 770767, Priority: 1, LastUsedTime: 1666769654},
		{CardIp: "114.101.80.194", IpCard: "ppp17626978", UsedWeight: 0, Priority: 0, LastUsedTime: 1666769649},
		{CardIp: "114.101.81.113", IpCard: "ppp17627018", UsedWeight: 0, Priority: -1, LastUsedTime: 1666769625},
		{CardIp: "114.101.80.26", IpCard: "ppp17626964", UsedWeight: 0, Priority: -4, LastUsedTime: 1666769629},
		{CardIp: "114.101.81.105", IpCard: "ppp17627024", UsedWeight: 0, Priority: -4, LastUsedTime: 1666769629},
		{CardIp: "114.101.83.79", IpCard: "ppp17627014", UsedWeight: 0, Priority: -6, LastUsedTime: 1666769621},
		{CardIp: "114.101.80.30", IpCard: "ppp17626972", UsedWeight: 0, Priority: -7, LastUsedTime: 1666769613},
		{CardIp: "114.101.82.119", IpCard: "ppp17627016", UsedWeight: 0, Priority: -8, LastUsedTime: 1666769601},
		{CardIp: "114.101.83.184", IpCard: "ppp17627002", UsedWeight: 0, Priority: -8, LastUsedTime: 1666769601},
		{CardIp: "114.101.83.231", IpCard: "ppp17626990", UsedWeight: 0, Priority: -11, LastUsedTime: 1666769585},
		{CardIp: "114.101.81.88", IpCard: "ppp17626988", UsedWeight: 0, Priority: -13, LastUsedTime: 1666769581},
		{CardIp: "114.101.81.155", IpCard: "ppp17626984", UsedWeight: 0, Priority: -14, LastUsedTime: 1666769589},
	}
	for name := range newAffinityMgr {
		sort.Slice(newAffinityMgr[name], func(i, j int) bool {
			if newAffinityMgr[name][i].Priority != newAffinityMgr[name][j].Priority {
				return newAffinityMgr[name][i].Priority > newAffinityMgr[name][j].Priority
			}
			if newAffinityMgr[name][i].LastUsedTime != newAffinityMgr[name][j].LastUsedTime {
				return newAffinityMgr[name][i].LastUsedTime > newAffinityMgr[name][j].LastUsedTime
			}
			return newAffinityMgr[name][i].CardIp < newAffinityMgr[name][j].CardIp
		})
	}
	fmt.Println("aaaa")
}
