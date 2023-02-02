package config

import (
	"encoding/json"
	"fmt"
	"health/enum"
	"health/share"
)

type (
	multiHunterResp struct {
		Code int           `json:"code"`
		Msg  string        `json:"msg"`
		Data *TargetUnitV3 `json:"data"`
	}
	TargetUnitV3 struct {
		FeatureSwitch *FeatureSwitch  `json:"feature_switch"`
		CoverList     []*CoverUnit    `json:"cover_list"`
		SNodeInfo     *SNodeCfgInfoV3 `json:"s_node_info"`
	}
	CoverUnit struct {
		Name     share.CoverName `json:"name"`
		ViewList []*ViewUnit     `json:"view_list"`
	}
	ViewUnit struct {
		Name      share.ViewName `json:"name"`
		Weight    float64        `json:"weight"`
		CacheList []*CacheUnitV3 `json:"cache_list"`
	}
	FeatureSwitch struct {
		QualityDecisionV2     bool `json:"quality_decision_v2"`
		ArriveRateDecision    bool `json:"arrive_rate_decision"`
		StutterDecision       bool `json:"stutter_decision"`
		PacketLostDecision    bool `json:"packet_lost_decision"`
		DownloadSpeedDecision bool `json:"download_speed_decision"`
	}
	CacheUnitV3 struct {
		CacheGroupName share.CacheGroupName `json:"cache_group_name"`
		HunterUrlV4    string               `json:"hunter_url_v4"`
		HunterUrlV6    string               `json:"hunter_url_v6"`
		MasterHost     CacheHostUnitV3      `json:"master_host"`
	}
	CacheHostUnitV3 struct {
		Hostname string `json:"hostname"`
		IPV4     string `json:"ip_v4"`
		IPV6     string `json:"ip_v6"`
	}
	SNodeCfgInfoV3 struct {
		SvgName     share.SvgName   `json:"svg_name"`
		Isp         string          `json:"isp"`
		Province    string          `json:"province"`
		Region      string          `json:"region"`
		Stype       string          `json:"s_type"`
		BuildLine   int64           `json:"build_line"`
		PlanLine    int64           `json:"plan_line"`
		IpType      enum.IpTypeEnum `json:"ip_type"`
		MachineType string          `json:"machine_type"`
	}
)

func ResolveConf() {
	result := &multiHunterResp{}
	input := "{\n\t\"code\": 0,\n\t\"msg\": \"success\",\n\t\"data\": {\n\t\t\"feature_switch\": {\n\t\t\t\"quality_decision_v2\": true,\n\t\t\t\"arrive_rate_decision\": true,\n\t\t\t\"stutter_decision\": false,\n\t\t\t\"packet_lost_decision\": false\n\t\t},\n\t\t\"s_node_info\": {\n\t\t\t\"svg_name\": \"lt-beijing-fog_docker_fogcdn-svg-1\",\n\t\t\t\"isp\": \"lt\",\n\t\t\t\"province\": \"anhui\",\n\t\t\t\"region\": \"huadong\",\n\t\t\t\"s_type\": \"1.6\",\n\t\t\t\"machine_type\": \"1.5\",\n\t\t\t\"build_line\": 2500000000,\n\t\t\t\"plan_line\": 2250000000,\n\t\t\t\"ip_type\": 400001\n\t\t},\n\t\t\"cover_list\": [\n\t\t\t{\n\t\t\t\t\"name\": \"filetest\",\n\t\t\t\t\"view_list\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"name\": \"default\",\n\t\t\t\t\t\t\"weight\": 0,\n\t\t\t\t\t\t\"cache_list\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"cache_group_name\": \"lt-shandong-zaozhuang-12-cache-5\",\n\t\t\t\t\t\t\t\t\"hunter_url_v4\": \"http://124.132.138.78:7080\",\n\t\t\t\t\t\t\t\t\"hunter_url_v6\": \"http://[2408:8719:4000:b::1:2f]:7080\",\n\t\t\t\t\t\t\t\t\"master_host\": {\n\t\t\t\t\t\t\t\t\t\"hostname\": \"lt-shandong-zaozhuang-12-124-132-138-78\",\n\t\t\t\t\t\t\t\t\t\"ip_v4\": \"124.132.138.78\",\n\t\t\t\t\t\t\t\t\t\"ip_v6\": \"[2408:8719:4000:b::1:2f]\"\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t]\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t}\n\t\t]\n\t}\n}"
	json.Unmarshal([]byte(input), result)
	// 最终决策策略
	input2 := CacheHostUnitV3{
		IPV4:     "1",
		IPV6:     "1",
		Hostname: "1",
	}
	fmt.Println(8%4 == 0)
	fmt.Println(444%4 == 0)
	switch input2.IPV4 {
	case input2.IPV6:
		fmt.Println(1111)
	case input2.Hostname:
		fmt.Println(222)
	case input2.IPV4:
		fmt.Println(333)
	}
}
