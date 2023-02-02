package main

import (
	"fmt"
	"sort"
)

type AffinityInfo struct {
	CardIp     string
	UsedWeight int64
}

var AffinityMgr = make(map[string][]AffinityInfo, 6)

//[
//{CardIp:111.122.60.173 IpCard:ppp17233827 UsedWeight:0 LastUsedTime:1663904059}
//{CardIp:111.122.60.189 IpCard:ppp17233825 UsedWeight:0 LastUsedTime:1663904059}
//{CardIp:111.122.60.52 IpCard:ppp17233823 UsedWeight:0 LastUsedTime:1663904059}
//{CardIp:111.122.60.9 IpCard:ppp17233833 UsedWeight:2506393 LastUsedTime:1663904059}
//{CardIp:111.122.61.112 IpCard:ppp17233829 UsedWeight:12907676 LastUsedTime:1663904059}]
func main() {
	var newAffinityMgr = make(map[string][]AffinityInfo, 6)

	funcxxx(newAffinityMgr)
	AffinityMgr = newAffinityMgr

	arr := []string{"name2", "name3", "name1", "name6", "name4", "name5"}
	sort.Strings(arr)
	for name := range newAffinityMgr {
		sort.Slice(newAffinityMgr[name], func(i, j int) bool {
			if newAffinityMgr[name][i].UsedWeight == newAffinityMgr[name][j].UsedWeight {
				return newAffinityMgr[name][i].CardIp < newAffinityMgr[name][j].CardIp
			}
			return newAffinityMgr[name][i].UsedWeight > newAffinityMgr[name][j].UsedWeight
		})
	}
	// 按cache组名称排序
	//sort.Slice(result, func(i, j int) bool {
	//	if result[i].UsedWeight == result[j].UsedWeight {
	//		return result[i].CardIp < result[j].CardIp
	//	}
	//	return result[i].UsedWeight > result[j].UsedWeight
	//})
	fmt.Printf("hhhh")
}

func funcxxx(mgr map[string][]AffinityInfo) {
	result := make([]AffinityInfo, 0, 5)
	result = append(result,
		AffinityInfo{"111.122.60.173", 2},
		AffinityInfo{"111.122.60.52", 3},
		AffinityInfo{"111.122.60.189", 4},
		AffinityInfo{"111.122.60.9", 3},
		AffinityInfo{"111.122.61.112", 1})
	mgr["name2"] = result
	mgr["name3"] = result
	mgr["name1"] = result
	mgr["name6"] = result
	mgr["name4"] = result
	mgr["name5"] = result
}
