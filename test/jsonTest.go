package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type AlarmInfo struct {
	Endpoint string                 `json:"endpoint"` // 主机名 如果传空，监控组件会取当前机器的主机名
	Step     int64                  `json:"step"`     // 采集时间步长 如果为空，默认取 60 单位 秒
	Value    int64                  `json:"value"`    // 默认value 当 fields 只有一个字段的时候，推荐只使用该值
	Tags     map[string]interface{} `json:"tags"`     // 用于标识报警的唯一性
	Fields   map[string]interface{} `json:"fields"`   // 主要数据
	Time     int64                  `json:"time"`     // 秒级时间戳 如果为空, 默认取当前时间
	Name     string                 `json:"name"`     // 监控数据表名, 必传
}

const (
	// 监控指标名
	TABLE_S_CLIENT_IP_QUALITY   = "s_client_ip_quality"
	TABLE_S_CLIENT_NODE_QUALITY = "s_client_node_quality"
	TABLE_S_CLIENT_IP_AFFINITY  = "s_client_ip_affinity"
	TABLE_S_CLIENT_PROCESS      = "s_client_process"

	// 标签参数名称
	TAG_HTTP_PORT  = "http_port"
	TAG_HTTPS_PORT = "https_port"
	TAG_SNODE      = "snode"
	TAG_IP         = "ip"
	TAG_IP_CARD    = "ip_card"
	TAG_CARD_IP    = "card_ip"
	TAG_CACHE      = "cache"

	// 字段参数名称
	FIELD_IS_IPV6                        = "is_v6"
	FIELD_EXPECT_WEIGHT                  = "expect_weight"
	FIELD_ADJUST_LINE                    = "adjust_line"
	FIELD_BUILD_LINE                     = "build_line"
	FIELD_PLAN                           = "plan"
	FIELD_BW                             = "bw"
	FIELD_AVG_LOST                       = "avg_lost"
	FIELD_AVG_RTT                        = "avg_rtt"
	FIELD_ALIVE_STATUS                   = "alive_status"
	FIELD_ADJUST_CONTROL_RATE            = "adjust_control_rate"
	FIELD_PLAN_CONTROL_RATE              = "plan_control_rate"
	FIELD_CONTROL_SUCCESS                = "control_success"
	FIELD_ADJUST_LINE_BY_LOST            = "adjust_line_by_lost"
	FIELD_ADJUST_LINE_BY_ARRIVE          = "adjust_line_by_arrive"
	FIELD_ADJUST_CONTROL_RATE_BY_ARRIVE  = "adjust_control_rate_by_arrive"
	FIELD_PLAN_CONTROL_RATE_BY_ARRIVE    = "plan_control_rate_by_arrive"
	FIELD_CONTROL_SUCCESS_BY_ARRIVE      = "control_success_by_arrive"
	FIELD_ADJUST_LINE_BY_STUTTER         = "adjust_line_by_stutter"
	FIELD_ADJUST_CONTROL_RATE_BY_STUTTER = "adjust_control_rate_by_stutter"
	FIELD_PLAN_CONTROL_RATE_BY_STUTTER   = "plan_control_rate_by_stutter"
	FIELD_CONTROL_SUCCESS_BY_STUTTER     = "control_success_by_stutter"
	FIELD_USED_EXPECT_WEIGHT             = "used_expect_weight"
	FIELD_AFFINITY_PRIORITY              = "affinity_priority"
	FIELD_LIMIT                          = "limit"
	FIELD_DECISION_BW                    = "decision_bw"
	FIELD_CURRENT_BW                     = "current_bw"
	FIELD_CURRENT_SPEED                  = "current_speed"
	FIELD_SMOOTH_SPEED                   = "smooth_speed"
	FIELD_LIMIT_LEVEL                    = "limit_level"
	FIELD_LIMIT_RATE                     = "limit_rate"
	FIELD_AVG_REQ_SIZE                   = "avg_req_size"
	FIELD_DECISION_QPS                   = "decision_qps"
	FIELD_CACHE_RATE                     = "cache_rate"
	FIELD_OPEN_V2_QUALITY                = "open_v2_quality"
	FIELD_MEM_PERCENT                    = "mem_percent"
	FIELD_MEM_INUSE                      = "mem_inuse"
	FIELD_GOROUTINE_NUM                  = "goroutine_num"
	FIELD_THREAD_NUM                     = "thread_num"
)

func Hostname() string {
	return "hostname"
}

type HDemo struct {
	HttpPort  uint32
	HttpsPort uint32
}

type AdjustLineByQuality struct {
	Stutter    int64
	ArriveRate int64
	PacketLost int64
}

// ReportV2 上报内容
type ReportV2 struct {
	Hostname    string       `json:"hostname"`
	TTL         int64        `json:"ttl"`
	Version     int64        `json:"version"`
	SNodeInfo   SNodeInfoV3  `json:"s_node_info"`
	SNodeIpList []*IPAssign  `json:"s_node_ip_list"`
	CoverList   []*CoverUnit `json:"cover_list"`
}

// SNodeInfoV3 节点信息
type SNodeInfoV3 struct {
	CurrentSpeed int64 `json:"current_speed"` // 当前下载速度
	SmoothSpeed  int64 `json:"smooth_speed"`  // 平滑下载速度
	AvgReqSize   int64 `json:"avg_req_size"`  // 平均请求文件大小
	//LimitLevel       int     `json:"limit_level"`        // 监控所需
	//LimitRate        float64 `json:"limit_rate"`         // 监控所需
	DetectTimeSecond int64 `json:"detect_time_second"` // 探测结果时间
	DecisionQps      int64 `json:"decision_qps"`       // 决策qps
	CacheRate        int64 `json:"cache_rate"`         // cache组分配比例
	//OpenV2Quality    bool    `json:"open_v2_quality"`    // 是否开启2.0质量决策
	AvgBwLastMinute int64   `json:"avg_bw_last_minute"` // 上一分钟平均带宽
	PacketLostRate  float64 `json:"packet_lost_rate"`   // 丢包率
	ArriveRate      float64 `json:"arrive_rate"`        // 到达率
	Stutter         float64 `json:"stutter"`            // 卡顿率
}

type CoverUnit struct {
	Name     string      `json:"name"`
	ViewList []*ViewUnit `json:"view_list"`
}

type ViewUnit struct {
	Name   string      `json:"name"`
	IpList []*IPAssign `json:"ip_list"`
}

type IPAssign struct {
	IpCard            string `json:"ip_card"`
	Ip                string `json:"ip"`
	CardIp            string `json:"card_ip"`
	TotalExpectWeight int64  `json:"total_expect_weight"`
	UsedExpectWeight  int64  `json:"used_expect_weight"`
}

var (
	lastAdjustLineByQualityMap = make(map[string]AdjustLineByQuality)
)

const lllkkk = `{
    "switch":true,
    "period":4,
    "silent_round_number":4,
    "cover_priority":"zttvi-p,zttvi-p2,zttapk-s,default",
    "cover_map":{
        "default":[
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.15
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.1
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":0.7,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.7,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.15
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.1
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":0.7,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.7,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.025,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.025,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.025,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.03,
                "upper_limit":0.04,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.03,
                "upper_limit":0.04,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.04,
                "upper_limit":0.05,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1
            },
            {
                "lower_limit":0.05,
                "upper_limit":0.07,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.95,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.07,
                "upper_limit":0.09,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.9,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.09,
                "upper_limit":0.11,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.85,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.11,
                "upper_limit":0.15,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.8,
                "adjust_line_control_rate":0.9
            },
            {
                "lower_limit":0.15,
                "upper_limit":0.2,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.75,
                "adjust_line_control_rate":0.9
            },
            {
                "lower_limit":0.2,
                "upper_limit":10,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.7,
                "adjust_line_control_rate":0.9
            }
        ],
        "zttapk-s":[
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.15
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.1
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":0.7,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.7,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.15
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.1
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":0.7,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.7,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.025,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.025,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.025,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.03,
                "upper_limit":0.04,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.03,
                "upper_limit":0.04,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.04,
                "upper_limit":0.05,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1
            },
            {
                "lower_limit":0.05,
                "upper_limit":0.07,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.95,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.07,
                "upper_limit":0.09,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.9,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.09,
                "upper_limit":0.11,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.85,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.11,
                "upper_limit":0.15,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.8,
                "adjust_line_control_rate":0.9
            },
            {
                "lower_limit":0.15,
                "upper_limit":0.2,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.75,
                "adjust_line_control_rate":0.9
            },
            {
                "lower_limit":0.2,
                "upper_limit":10,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.7,
                "adjust_line_control_rate":0.9
            }
        ],
        "zttvi-p":[
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.15
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.1
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":0.7,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.7,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.03
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.03,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.03,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.03,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.03
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.03,
                "upper_limit":0.035,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1
            },
            {
                "lower_limit":0.035,
                "upper_limit":0.04,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.97,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.04,
                "upper_limit":0.05,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.94,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.05,
                "upper_limit":0.07,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.9,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.07,
                "upper_limit":0.09,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.85,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.09,
                "upper_limit":0.11,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.8,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.11,
                "upper_limit":0.15,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.75,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.15,
                "upper_limit":0.2,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.7,
                "adjust_line_control_rate":0.9
            },
            {
                "lower_limit":0.2,
                "upper_limit":10,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.6,
                "adjust_line_control_rate":0.9
            }
        ],
        "zttvi-p2":[
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.15
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.1
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":0.7,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0,
                "upper_limit":0.01,
                "gt_plan_ratio":0.7,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.03
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.01,
                "upper_limit":0.02,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.03,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.08
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.03,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.02,
                "upper_limit":0.03,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0,
                "le_plan_ratio":0.2,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.05
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.2,
                "le_plan_ratio":0.5,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.03
            },
            {
                "lower_limit":0.025,
                "upper_limit":0.03,
                "gt_plan_ratio":0.5,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1.02
            },
            {
                "lower_limit":0.03,
                "upper_limit":0.035,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":1,
                "adjust_line_control_rate":1
            },
            {
                "lower_limit":0.035,
                "upper_limit":0.04,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.97,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.04,
                "upper_limit":0.05,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.94,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.05,
                "upper_limit":0.07,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.9,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.07,
                "upper_limit":0.09,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.85,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.09,
                "upper_limit":0.11,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.8,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.11,
                "upper_limit":0.15,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.75,
                "adjust_line_control_rate":0.95
            },
            {
                "lower_limit":0.15,
                "upper_limit":0.2,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.7,
                "adjust_line_control_rate":0.9
            },
            {
                "lower_limit":0.2,
                "upper_limit":10,
                "gt_plan_ratio":0,
                "le_plan_ratio":10,
                "plan_line_control_rate":0.6,
                "adjust_line_control_rate":0.9
            }
        ]
    }
}`

type Config struct {
	Switch            bool                     `json:"switch"`
	Period            int                      `json:"period"`
	SilentRoundNumber int                      `json:"silent_round_number"`
	CoverPriority     string                   `json:"cover_priority"`
	CoverMap          map[string][]*RangeLimit `json:"cover_map"`
}

type RangeLimit struct {
	LowerLimit            float64 `json:"lower_limit"`              // 丢包率下限
	UpperLimit            float64 `json:"upper_limit"`              // 丢包率上限
	GtPlanRatio           float64 `json:"gt_plan_ratio"`            // 上一次调度线大于规划线的比例
	LePlanRatio           float64 `json:"le_plan_ratio"`            // 上一次调度线小于等于规划线的比例
	PlanLineControlRate   float64 `json:"plan_line_control_rate"`   // 规划线的控线比例
	AdjustLineControlRate float64 `json:"adjust_line_control_rate"` // 调度线的控线比例
}

func main() {
	tmmmm := &Config{}
	err1 := json.Unmarshal([]byte(lllkkk), tmmmm)
	fmt.Println(err1)
	ddd1 := 0
	ccc1 := 1 / ddd1
	fmt.Println(ccc1)
	iPAbilityMap := make(map[string]*IPAssign)

	iPAbilityMap["111"] = &IPAssign{
		Ip:               "1.1.1.1",
		UsedExpectWeight: 110,
	}
	kkk := iPAbilityMap["111"]
	kkk.Ip = "2.2.2.2"
	result := make(map[string]*ReportV2, 0)
	affinityInfoMap := make(map[string]IPAssign, 0)
	hahaha(affinityInfoMap)
	var keys []string
	keys = append(keys, "b")
	keys = append(keys, "a")
	keys = append(keys, "d")
	keys = append(keys, "c")
	reportInfo := buildReportV2Info(666)
	lastAdjustLineByQualityMap["aaa"] = AdjustLineByQuality{
		Stutter:    int64(123),
		ArriveRate: int64(123),
		PacketLost: int64(123),
	}
	lastAdjustLineByQualityMap["bbb"] = AdjustLineByQuality{
		Stutter:    int64(456),
		ArriveRate: int64(456),
		PacketLost: int64(456),
	}
	for k, _ := range lastAdjustLineByQualityMap {
		item := reportInfo
		item.TTL--
		modify(item)
		result[k] = &item
	}
	delete(result, "aaa")
	tempAdjustLineByQualityMap := make(map[string]AdjustLineByQuality)
	lastAdjustLineByQualityMap["aaa"] = AdjustLineByQuality{
		Stutter:    int64(123),
		ArriveRate: int64(123),
		PacketLost: int64(123),
	}
	bbb := &AdjustLineByQuality{
		Stutter:    int64(123),
		ArriveRate: int64(123),
		PacketLost: int64(123),
	}
	ccc := *bbb
	ccc.ArriveRate = 2022
	ddd := &ccc
	fmt.Println(ddd)
	testFunc(tempAdjustLineByQualityMap)
	demo := HDemo{}
	aaa := AlarmInfo{
		Tags: map[string]interface{}{
			TAG_SNODE:      Hostname(),
			TAG_HTTP_PORT:  strconv.Itoa(int(demo.HttpPort)),
			TAG_HTTPS_PORT: strconv.Itoa(int(demo.HttpsPort)),
		},
		Fields: map[string]interface{}{
			FIELD_IS_IPV6: 1,
		},
		Name:     TABLE_S_CLIENT_IP_QUALITY,
		Endpoint: Hostname(),
	}
	data, err := json.Marshal(aaa)
	fmt.Println(data, err)
}

func hahaha(infoMap map[string]IPAssign) {
	infoMap["dasds"] = IPAssign{
		Ip: "1.1.1.1",
	}
}

func modify(item ReportV2) {
	item.TTL = 999

}

func testFunc(qualityMap map[string]AdjustLineByQuality) {
	qualityMap["aaa"] = lastAdjustLineByQualityMap["aaa"]
}

func buildReportV2Info(sNodeAvgBwLastMinute int64) ReportV2 {
	result := ReportV2{
		Hostname: Hostname(),
		TTL:      300,
		Version:  time.Now().UnixNano(),
		SNodeInfo: SNodeInfoV3{
			AvgBwLastMinute: sNodeAvgBwLastMinute,
		},
	}
	return result
}
