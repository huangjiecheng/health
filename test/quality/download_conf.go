package main

import (
	"encoding/json"
	"fmt"
)

type (
	DpConfig struct {
		Switch               bool                     `json:"switch"`
		SyncInterval         int                      `json:"sync_interval"`
		SpeedLimitMinRate    float64                  `json:"Speed_limit_min_rate"`
		SpeedLimitFreezeTime int64                    `json:"Speed_limit_freeze_time"`
		CoverMap             map[string][]*CoverInfo2 `json:"cover_map"`
	}

	CoverInfo2 struct {
		Speed int64   `json:"Speed"`
		Rate  float64 `json:"Rate"`
	}
)

func main() {
	finalMap := make(map[string][]*CoverInfo2)
	key, conf := defaultDp()
	finalMap[key] = conf
	key, conf = dpzttapk_s()
	finalMap[key] = conf
	key, conf = dpzttvi_p()
	finalMap[key] = conf

	key, conf = dppf001()
	finalMap[key] = conf
	key, conf = dppf003()
	finalMap[key] = conf
	key, conf = dpfaiqiyi()
	finalMap[key] = conf
	key, conf = dpzqctgdl()
	finalMap[key] = conf
	finalMap["filetest"] = conf
	input := &DpConfig{
		Switch:               true,
		SyncInterval:         4,
		SpeedLimitMinRate:    0.01,
		SpeedLimitFreezeTime: 120,
		CoverMap:             finalMap,
	}
	data, err := json.Marshal(input)
	data2, _ := json.Marshal(finalMap)
	fmt.Println(string(data), err)
	fmt.Println(string(data2), err)
	demosuccess()
}

func demosuccess() {
	input := "{\n  \"switch\": true,\n  \"sync_interval\": 4,\n  \"Speed_limit_min_rate\": 0.01,\n  \"Speed_limit_freeze_time\": 120,\n  \"cover_map\": {\n    \"default\": [\n      {\n        \"Speed\": 7,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 4,\n        \"Rate\": 0.95\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 2,\n        \"Rate\": 0.85\n      },\n      {\n        \"Speed\": 1,\n        \"Rate\": 0.05\n      }\n    ],\n    \"faiqiyi\": [\n      {\n        \"Speed\": 20,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 16,\n        \"Rate\": 0.8\n      },\n      {\n        \"Speed\": 10,\n        \"Rate\": 0.7\n      },\n      {\n        \"Speed\": 7,\n        \"Rate\": 0.6\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 0.4\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 0.1\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.05\n      }\n    ],\n    \"filetest\": [\n      {\n        \"Speed\": 7,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 4,\n        \"Rate\": 0.95\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 2,\n        \"Rate\": 0.85\n      },\n      {\n        \"Speed\": 1,\n        \"Rate\": 0.05\n      }\n    ],\n    \"pf001\": [\n      {\n        \"Speed\": 7,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 4,\n        \"Rate\": 0.95\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 2,\n        \"Rate\": 0.85\n      },\n      {\n        \"Speed\": 1,\n        \"Rate\": 0.05\n      }\n    ],\n    \"zqctgdl\": [\n      {\n        \"Speed\": 7,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 4,\n        \"Rate\": 0.95\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 2,\n        \"Rate\": 0.85\n      },\n      {\n        \"Speed\": 1,\n        \"Rate\": 0.05\n      }\n    ],\n    \"zttapk-s\": [\n      {\n        \"Speed\": 7,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 4,\n        \"Rate\": 0.95\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 2,\n        \"Rate\": 0.85\n      },\n      {\n        \"Speed\": 1,\n        \"Rate\": 0.05\n      }\n    ],\n    \"zttvi-p\": [\n      {\n        \"Speed\": 9,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 8,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 7,\n        \"Rate\": 1\n      },\n      {\n        \"Speed\": 6,\n        \"Rate\": 0.95\n      },\n      {\n        \"Speed\": 5,\n        \"Rate\": 0.9\n      },\n      {\n        \"Speed\": 4,\n        \"Rate\": 0.85\n      },\n      {\n        \"Speed\": 3,\n        \"Rate\": 0.05\n      }\n    ]\n  }\n}"
	result := &DpConfig{}
	json.Unmarshal([]byte(input), result)
	fmt.Println(result)
}

func defaultDp() (string, []*CoverInfo2) {
	cover := "default"
	result := []*CoverInfo2{
		{
			Speed: 7,
			Rate:  1.0,
		},
		{
			Speed: 6,
			Rate:  1.0,
		},
		{
			Speed: 5,
			Rate:  1.0,
		},
		{
			Speed: 4,
			Rate:  0.95,
		},
		{
			Speed: 3,
			Rate:  0.90,
		},
		{
			Speed: 2,
			Rate:  0.85,
		},
		{
			Speed: 1,
			Rate:  0.05,
		},
	}
	return cover, result
}

func dpzttapk_s() (string, []*CoverInfo2) {
	cover := "zttapk-s"

	result := []*CoverInfo2{
		{
			Speed: 7,
			Rate:  1.0,
		},
		{
			Speed: 6,
			Rate:  1.0,
		},
		{
			Speed: 5,
			Rate:  1.0,
		},
		{
			Speed: 4,
			Rate:  0.95,
		},
		{
			Speed: 3,
			Rate:  0.90,
		},
		{
			Speed: 2,
			Rate:  0.85,
		},
		{
			Speed: 1,
			Rate:  0.05,
		},
	}
	return cover, result
}
func dpzttvi_p() (string, []*CoverInfo2) {
	cover := "zttvi-p"

	result := []*CoverInfo2{
		{
			Speed: 9,
			Rate:  1.0,
		},
		{
			Speed: 8,
			Rate:  1.0,
		},
		{
			Speed: 7,
			Rate:  1.0,
		},
		{
			Speed: 6,
			Rate:  0.95,
		},
		{
			Speed: 5,
			Rate:  0.90,
		},
		{
			Speed: 4,
			Rate:  0.85,
		},
		{
			Speed: 3,
			Rate:  0.05,
		},
	}
	return cover, result
}

func dppf001() (string, []*CoverInfo2) {
	cover := "pf001"

	result := []*CoverInfo2{
		{
			Speed: 7,
			Rate:  1.0,
		},
		{
			Speed: 6,
			Rate:  1.0,
		},
		{
			Speed: 5,
			Rate:  1.0,
		},
		{
			Speed: 4,
			Rate:  0.95,
		},
		{
			Speed: 3,
			Rate:  0.90,
		},
		{
			Speed: 2,
			Rate:  0.85,
		},
		{
			Speed: 1,
			Rate:  0.05,
		},
	}
	return cover, result
}
func dppf003() (string, []*CoverInfo2) {
	cover := "pf001"

	result := []*CoverInfo2{
		{
			Speed: 7,
			Rate:  1.0,
		},
		{
			Speed: 6,
			Rate:  1.0,
		},
		{
			Speed: 5,
			Rate:  1.0,
		},
		{
			Speed: 4,
			Rate:  0.95,
		},
		{
			Speed: 3,
			Rate:  0.90,
		},
		{
			Speed: 2,
			Rate:  0.85,
		},
		{
			Speed: 1,
			Rate:  0.05,
		},
	}
	return cover, result
}
func dpfaiqiyi() (string, []*CoverInfo2) {
	cover := "faiqiyi"

	result := []*CoverInfo2{
		{
			Speed: 20,
			Rate:  0.9,
		},
		{
			Speed: 16,
			Rate:  0.8,
		},
		{
			Speed: 10,
			Rate:  0.7,
		},
		{
			Speed: 7,
			Rate:  0.6,
		},
		{
			Speed: 6,
			Rate:  0.4,
		},
		{
			Speed: 5,
			Rate:  0.1,
		},
		{
			Speed: 3,
			Rate:  0.05,
		},
	}
	return cover, result
}
func dpzqctgdl() (string, []*CoverInfo2) {
	cover := "zqctgdl"

	result := []*CoverInfo2{
		{
			Speed: 7,
			Rate:  1.0,
		},
		{
			Speed: 6,
			Rate:  1.0,
		},
		{
			Speed: 5,
			Rate:  1.0,
		},
		{
			Speed: 4,
			Rate:  0.95,
		},
		{
			Speed: 3,
			Rate:  0.90,
		},
		{
			Speed: 2,
			Rate:  0.85,
		},
		{
			Speed: 1,
			Rate:  0.05,
		},
	}
	return cover, result
}
