package demo

import (
	"fmt"
	"health/enum"
	"health/share"
)

// 响应体
type (
	HunterRespBody struct {
		Code int             `json:"code"`
		Msg  string          `json:"msg"`
		Data *CacheGroupData `json:"data"`
	}

	CacheGroupData struct {
		Cache                string                             `json:"cache"`
		Hostname             string                             `json:"hostname"`
		SardineStatsRealInfo []SardineReqInfo                   `json:"sardine_stats_real_info"` // 2轮共4秒的统计
		LastMinutePassInfo   LastMinutePassInfo                 `json:"last_minute_pass_info"`   // 上分钟放行信息
		DetectFailCardIps    map[share.CardIp]enum.IpStatusEnum `json:"detect_fail_card_ips"`
	}

	SardineReqInfo struct {
		CoverName               share.CoverName   `json:"cover_name"`
		RegionName              share.RegionName  `json:"region_name"`
		EndTimeMillSecond       share.TimestampMs `json:"end_time_mill_second"`       //结束时间(毫秒)
		PassReqCount            int64             `json:"pass_req_count"`             //放行请求数
		InterceptReqCount       int64             `json:"intercept_req_count"`        //拦截请求数
		RemotePassReqCount      int64             `json:"remote_pass_req_count"`      // 远程sardine放行请求数
		RemoteInterceptReqCount int64             `json:"remote_intercept_req_count"` // 远程sardine拦截请求数
	}

	LastMinutePassInfo struct {
		StartTimeMillSecond share.TimestampMs                `json:"start_time_mill_second"`
		CoverMap            map[share.CoverName]ReqCountInfo `json:"cover_map"`
		IPMap               map[share.CardIp]ReqCountInfo    `json:"ip_map"`
	}

	ReqCountInfo struct {
		PassReqCount            int64 `json:"pass_req_count"`             // 放行请求数
		InterceptReqCount       int64 `json:"intercept_req_count"`        // 拦截请求数
		RemotePassReqCount      int64 `json:"remote_pass_req_count"`      // 远程放行请求数
		RemoteInterceptReqCount int64 `json:"remote_intercept_req_count"` // 远程拦截请求数
	}
)

const (
	respBody = `{
    "code": 0,
    "msg": "OK",
    "data": {
        "cache": "dx-lt-yd-hebei-shijiazhuang-10-cache-8",
        "hostname": "dx-lt-yd-hebei-shijiazhuang-10-124-236-69-79",
        "s_node_status": 1,
        "sardine_stats_real_info": [
            {
                "cover_name": "filetest",
                "region_name": "hunan",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "anhui",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "jiangxi",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "_default_region_4.0",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "zhejiang",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "hubei",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "jiangsu",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "fujian",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "shanghai",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "henan",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            },
            {
                "cover_name": "filetest",
                "region_name": "shandong",
                "end_time_mill_second": 1695796839166,
                "pass_req_count": 0,
                "intercept_req_count": 0,
                "remote_pass_req_count": 0,
                "remote_intercept_req_count": 0
            }
        ],
        "detect_fail_card_ips": {},
        "hunter_store_ip_count": 0,
        "last_minute_req_info": {
            "start_time_mill_second": 1695796740000,
            "pass_req_count": 0,
            "intercept_req_count": 0,
            "remote_pass_req_count": 0,
            "remote_intercept_req_count": 0
        },
        "last_minute_req_info_by_cover": {
            "start_time_mill_second": 1695796740000,
            "cover_map": {
                "filetest": {
                    "pass_req_count": 0,
                    "intercept_req_count": 0,
                    "remote_pass_req_count": 0,
                    "remote_intercept_req_count": 0
                }
            },
            "ip_map": null
        },
        "last_minute_pass_info": {
            "start_time_mill_second": 1695796740000,
            "cover_map": {
                "filetest": {
                    "pass_req_count": 0,
                    "intercept_req_count": 0,
                    "remote_pass_req_count": 0,
                    "remote_intercept_req_count": 0
                }
            },
            "ip_map": {}
        }
    }
}`
)

func Run() {
	//var result HunterRespBody
	//err := json.Unmarshal([]byte(respBody), &result)
	//if err != nil {
	//	fmt.Printf("uploadSNodeAbility failed to Unmarshal resp.err: %v", err)
	//	return
	//}

	fmt.Println(appendString("a", "b", "c"))
	var (
		viewList = []share.RegionName{
			"anhui",
			"fujian",
			"jiangsu",
			"jiangxi",
			"shandong",
			"shanghai",
			"zhejiang",
			"_default_region_4.0",
		}
		pros = []share.RegionName{"xibei",
			"xinan",
			"huazhong",
			"huabei",
			"_default_region_4.0"}
	)
	fmt.Println(sortRegionPriority(viewList, pros))
}

func appendString(strs ...string) string {
	result := ""
	for i, s := range strs {
		if i == 0 {
			result += s
			continue
		}
		result += "_" + s
	}
	return result
}

func sortRegionPriority(viewList, svgRegionPriority []share.RegionName) []share.RegionName {
	var (
		alreadyExistRegion = make(map[share.RegionName]struct{})
		result             = make([]share.RegionName, 0, len(viewList))
		existDefaultRegion = false
	)
	for _, sortedRegionName := range svgRegionPriority {
		for _, regionName := range viewList {
			if regionName == share.DefaultRegion {
				existDefaultRegion = true
				continue
			}
			if regionName == sortedRegionName {
				result = append(result, regionName)
				alreadyExistRegion[regionName] = struct{}{}
			}
		}
	}
	for _, regionName := range viewList {
		if _, exist := alreadyExistRegion[regionName]; !exist && regionName != share.DefaultRegion {
			result = append(result, regionName)
		}
	}
	if existDefaultRegion {
		result = append(result, share.DefaultRegion)
	}
	return result
}
