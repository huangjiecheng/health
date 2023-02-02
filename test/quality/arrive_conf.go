package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ArriveRateConfig struct {
	Switch           bool                         `json:"switch"`
	DecisionInterval int                          `json:"decision_interval"`
	CoverMap         map[string]*RangeLimitArrive `json:"cover_map"`
}

// RangeLimitArrive 到达率控线策略配置
type RangeLimitArrive struct {
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
type SyncConfigResponse struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data map[string]ConfigInfo `json:"data"`
}

type ConfigInfo struct {
	Id         int64  `json:"id"`
	Alias      string `json:"alias"`
	Content    string `json:"content"`
	ContentMd5 string `json:"content_md5"`
}

func main() {
	inputStr := "{\n\t\"code\": 0,\n\t\"msg\": \"success\",\n\t\"data\": {\n\t\t\"arrive_rate\": {\n\t\t\t\"id\": 24,\n\t\t\t\"alias\": \"default\",\n\t\t\t\"content\": \"{\\n    \\\"switch\\\":true,\\n    \\\"decision_interval\\\":60,\\n    \\\"cover_map\\\":{\\n        \\\"default\\\":{\\n            \\\"arrive_rate_list\\\":[\\n                0.6,\\n                0.65,\\n                0.7,\\n                0.75,\\n                0.8\\n            ],\\n            \\\"lt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"gt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"lt_avg_bw_plan_line_rate_list\\\":[\\n                0.6,\\n                0.7,\\n                0.8,\\n                0.9,\\n                1\\n            ],\\n            \\\"gt_avg_bw_plan_line_rate_list\\\":[\\n                0.5,\\n                0.6,\\n                0.7,\\n                0.8,\\n                0.9\\n            ],\\n            \\\"sardine_pass_count\\\":1000,\\n            \\\"silent_rate\\\":0.7,\\n            \\\"recovery_rate\\\":0.8,\\n            \\\"recovery_adjust_line_rate\\\":1.05\\n        },\\n        \\\"filetest\\\":{\\n            \\\"arrive_rate_list\\\":[\\n                0.93,\\n                0.94,\\n                0.95,\\n                0.96,\\n                0.97\\n            ],\\n            \\\"lt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"gt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"lt_avg_bw_plan_line_rate_list\\\":[\\n                0.6,\\n                0.7,\\n                0.8,\\n                0.9,\\n                1\\n            ],\\n            \\\"gt_avg_bw_plan_line_rate_list\\\":[\\n                0.5,\\n                0.6,\\n                0.7,\\n                0.8,\\n                0.9\\n            ],\\n            \\\"sardine_pass_count\\\":1000,\\n            \\\"silent_rate\\\":0.92,\\n            \\\"recovery_rate\\\":0.98,\\n            \\\"recovery_adjust_line_rate\\\":1.05\\n        },\\n        \\\"zttapk-s\\\":{\\n            \\\"arrive_rate_list\\\":[\\n                0.6,\\n                0.65,\\n                0.7,\\n                0.75,\\n                0.8\\n            ],\\n            \\\"lt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"gt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"lt_avg_bw_plan_line_rate_list\\\":[\\n                0.6,\\n                0.7,\\n                0.8,\\n                0.9,\\n                1\\n            ],\\n            \\\"gt_avg_bw_plan_line_rate_list\\\":[\\n                0.5,\\n                0.6,\\n                0.7,\\n                0.8,\\n                0.9\\n            ],\\n            \\\"sardine_pass_count\\\":1000,\\n            \\\"silent_rate\\\":0.7,\\n            \\\"recovery_rate\\\":0.8,\\n            \\\"recovery_adjust_line_rate\\\":1.05\\n        },\\n        \\\"zttvi-p\\\":{\\n            \\\"arrive_rate_list\\\":[\\n                0.7,\\n                0.75,\\n                0.8,\\n                0.85\\n            ],\\n            \\\"lt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"gt_avg_bw_adjust_line_rate_list\\\":[\\n                0.85,\\n                0.9,\\n                0.95,\\n                1\\n            ],\\n            \\\"lt_avg_bw_plan_line_rate_list\\\":[\\n                0.6,\\n                0.8,\\n                0.9,\\n                1\\n            ],\\n            \\\"gt_avg_bw_plan_line_rate_list\\\":[\\n                0.5,\\n                0.7,\\n                0.8,\\n                0.9\\n            ],\\n            \\\"sardine_pass_count\\\":2000,\\n            \\\"silent_rate\\\":0.8,\\n            \\\"recovery_rate\\\":0.85,\\n            \\\"recovery_adjust_line_rate\\\":1.05\\n        }\\n    }\\n}\",\n\t\t\t\"content_md5\": \"182b8941eaab1d7eaa308ffbf00d6c84\"\n\t\t},\n\t\t\"download_speed\": {\n\t\t\t\"id\": 26,\n\t\t\t\"alias\": \"default\",\n\t\t\t\"content\": \"{\\n    \\\"switch\\\":true,\\n    \\\"sync_interval\\\":4,\\n    \\\"speed_limit_min_rate\\\":0.01,\\n    \\\"speed_limit_freeze_time\\\":120,\\n    \\\"cover_map\\\":{\\n        \\\"default\\\":[\\n            {\\n                \\\"speed\\\":1000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":2000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            }\\n        ],\\n        \\\"faiqiyi\\\":[\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":0.1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":0.4\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":0.6\\n            },\\n            {\\n                \\\"speed\\\":10000000,\\n                \\\"rate\\\":0.7\\n            },\\n            {\\n                \\\"speed\\\":16000000,\\n                \\\"rate\\\":0.8\\n            },\\n            {\\n                \\\"speed\\\":20000000,\\n                \\\"rate\\\":0.9\\n            }\\n        ],\\n        \\\"filetest\\\":[\\n            {\\n                \\\"speed\\\":1000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":2000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            }\\n        ],\\n        \\\"pf001\\\":[\\n            {\\n                \\\"speed\\\":1000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":2000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            }\\n        ],\\n        \\\"pf003\\\":[\\n            {\\n                \\\"speed\\\":1000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":2000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            }\\n        ],\\n        \\\"zqctgdl\\\":[\\n            {\\n                \\\"speed\\\":1000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":2000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            }\\n        ],\\n        \\\"zttapk-s\\\":[\\n            {\\n                \\\"speed\\\":1000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":2000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            }\\n        ],\\n        \\\"zttvi-p\\\":[\\n            {\\n                \\\"speed\\\":3000000,\\n                \\\"rate\\\":0.05\\n            },\\n            {\\n                \\\"speed\\\":4000000,\\n                \\\"rate\\\":0.85\\n            },\\n            {\\n                \\\"speed\\\":5000000,\\n                \\\"rate\\\":0.9\\n            },\\n            {\\n                \\\"speed\\\":6000000,\\n                \\\"rate\\\":0.95\\n            },\\n            {\\n                \\\"speed\\\":7000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":8000000,\\n                \\\"rate\\\":1\\n            },\\n            {\\n                \\\"speed\\\":9000000,\\n                \\\"rate\\\":1\\n            }\\n        ]\\n    }\\n}\",\n\t\t\t\"content_md5\": \"f8c04a080b12c3c18107ea484005d188\"\n\t\t},\n\t\t\"loss\": {\n\t\t\t\"id\": 13,\n\t\t\t\"alias\": \"default\",\n\t\t\t\"content\": \"{\\n\\\"loss\\\":\\\"loss_config\\\"\\n}\\n\",\n\t\t\t\"content_md5\": \"12ba1602eb6d9cf57bf237813e008f18\"\n\t\t},\n\t\t\"packet_lost\": {\n\t\t\t\"id\": 25,\n\t\t\t\"alias\": \"default\",\n\t\t\t\"content\": \"{\\n    \\\"switch\\\":true,\\n    \\\"period\\\":4,\\n    \\\"silent_round_number\\\":4,\\n    \\\"cover_map\\\":{\\n        \\\"default\\\":{\\n            \\\"adjust_line_rate_list\\\":[\\n                1.05,\\n                1.04,\\n                1.03,\\n                1.02,\\n                1,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.9,\\n                0.9\\n            ],\\n            \\\"plan_line_rate_list\\\":[\\n                1,\\n                1,\\n                1,\\n                1,\\n                1,\\n                0.95,\\n                0.9,\\n                0.85,\\n                0.8,\\n                0.75,\\n                0.7\\n            ],\\n            \\\"pkg_lost_rate_list\\\":[\\n                0.01,\\n                0.02,\\n                0.03,\\n                0.04,\\n                0.05,\\n                0.07,\\n                0.09,\\n                0.11,\\n                0.15,\\n                0.2,\\n                10\\n            ]\\n        },\\n        \\\"zttapk-s\\\":{\\n            \\\"adjust_line_rate_list\\\":[\\n                1.05,\\n                1.04,\\n                1.03,\\n                1.02,\\n                1,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.9,\\n                0.9\\n            ],\\n            \\\"plan_line_rate_list\\\":[\\n                1,\\n                1,\\n                1,\\n                1,\\n                1,\\n                0.95,\\n                0.9,\\n                0.85,\\n                0.8,\\n                0.75,\\n                0.7\\n            ],\\n            \\\"pkg_lost_rate_list\\\":[\\n                0.01,\\n                0.02,\\n                0.03,\\n                0.04,\\n                0.05,\\n                0.07,\\n                0.09,\\n                0.11,\\n                0.15,\\n                0.2,\\n                10\\n            ]\\n        },\\n        \\\"zttvi-p\\\":{\\n            \\\"adjust_line_rate_list\\\":[\\n                1.05,\\n                1.02,\\n                1,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.95,\\n                0.9,\\n                0.9,\\n                0.9\\n            ],\\n            \\\"plan_line_rate_list\\\":[\\n                1,\\n                1,\\n                1,\\n                0.97,\\n                0.94,\\n                0.9,\\n                0.85,\\n                0.8,\\n                0.75,\\n                0.7,\\n                0.65,\\n                0.6\\n            ],\\n            \\\"pkg_lost_rate_list\\\":[\\n                0.01,\\n                0.02,\\n                0.025,\\n                0.03,\\n                0.04,\\n                0.05,\\n                0.07,\\n                0.09,\\n                0.11,\\n                0.15,\\n                0.2,\\n                10\\n            ]\\n        }\\n    }\\n}\",\n\t\t\t\"content_md5\": \"ea71b3139385f57172efd4bc92dd8769\"\n\t\t},\n\t\t\"stutter\": {\n\t\t\t\"id\": 23,\n\t\t\t\"alias\": \"default\",\n\t\t\t\"content\": \"{\\n    \\\"switch\\\":true,\\n    \\\"decision_interval\\\":60,\\n    \\\"cover_map\\\":{\\n        \\\"default\\\":{\\n            \\\"snode_a\\\":[\\n                {\\n                    \\\"lower_limit\\\":0,\\n                    \\\"upper_limit\\\":0.05,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.06,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.05,\\n                    \\\"upper_limit\\\":0.07,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.04,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.07,\\n                    \\\"upper_limit\\\":0.09,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.02,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.11,\\n                    \\\"upper_limit\\\":0.13,\\n                    \\\"plan_line_control_rate\\\":0.95,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.13,\\n                    \\\"upper_limit\\\":0.15,\\n                    \\\"plan_line_control_rate\\\":0.9,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.15,\\n                    \\\"upper_limit\\\":0.2,\\n                    \\\"plan_line_control_rate\\\":0.85,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.2,\\n                    \\\"upper_limit\\\":0.25,\\n                    \\\"plan_line_control_rate\\\":0.8,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.25,\\n                    \\\"upper_limit\\\":0.3,\\n                    \\\"plan_line_control_rate\\\":0.75,\\n                    \\\"adjust_line_control_rate\\\":0.85,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.3,\\n                    \\\"upper_limit\\\":1,\\n                    \\\"plan_line_control_rate\\\":0.7,\\n                    \\\"adjust_line_control_rate\\\":0.85,\\n                    \\\"do_decision_when_over_line\\\":true\\n                }\\n            ],\\n            \\\"snode_b\\\":[\\n                {\\n                    \\\"lower_limit\\\":0,\\n                    \\\"upper_limit\\\":0.06,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.06,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.06,\\n                    \\\"upper_limit\\\":0.08,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.04,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.08,\\n                    \\\"upper_limit\\\":0.1,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.02,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.12,\\n                    \\\"upper_limit\\\":0.14,\\n                    \\\"plan_line_control_rate\\\":0.95,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.14,\\n                    \\\"upper_limit\\\":0.16,\\n                    \\\"plan_line_control_rate\\\":0.9,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.16,\\n                    \\\"upper_limit\\\":0.19,\\n                    \\\"plan_line_control_rate\\\":0.85,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.19,\\n                    \\\"upper_limit\\\":0.22,\\n                    \\\"plan_line_control_rate\\\":0.8,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.22,\\n                    \\\"upper_limit\\\":0.25,\\n                    \\\"plan_line_control_rate\\\":0.75,\\n                    \\\"adjust_line_control_rate\\\":0.85,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.25,\\n                    \\\"upper_limit\\\":1,\\n                    \\\"plan_line_control_rate\\\":0.7,\\n                    \\\"adjust_line_control_rate\\\":0.85,\\n                    \\\"do_decision_when_over_line\\\":true\\n                }\\n            ]\\n        },\\n        \\\"zttapk-s\\\":{\\n            \\\"snode_a\\\":[\\n                {\\n                    \\\"lower_limit\\\":0,\\n                    \\\"upper_limit\\\":0.05,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.06,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.05,\\n                    \\\"upper_limit\\\":0.07,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.04,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.07,\\n                    \\\"upper_limit\\\":0.09,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.02,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.11,\\n                    \\\"upper_limit\\\":0.13,\\n                    \\\"plan_line_control_rate\\\":0.95,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.13,\\n                    \\\"upper_limit\\\":0.15,\\n                    \\\"plan_line_control_rate\\\":0.9,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.15,\\n                    \\\"upper_limit\\\":0.17,\\n                    \\\"plan_line_control_rate\\\":0.85,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.17,\\n                    \\\"upper_limit\\\":0.2,\\n                    \\\"plan_line_control_rate\\\":0.8,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.2,\\n                    \\\"upper_limit\\\":1,\\n                    \\\"plan_line_control_rate\\\":0.75,\\n                    \\\"adjust_line_control_rate\\\":0.85,\\n                    \\\"do_decision_when_over_line\\\":true\\n                }\\n            ],\\n            \\\"snode_b\\\":[\\n                {\\n                    \\\"lower_limit\\\":0,\\n                    \\\"upper_limit\\\":0.01,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.06,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.01,\\n                    \\\"upper_limit\\\":0.03,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.04,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.03,\\n                    \\\"upper_limit\\\":0.05,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.02,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.07,\\n                    \\\"upper_limit\\\":0.09,\\n                    \\\"plan_line_control_rate\\\":0.95,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.09,\\n                    \\\"upper_limit\\\":0.11,\\n                    \\\"plan_line_control_rate\\\":0.9,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.11,\\n                    \\\"upper_limit\\\":0.13,\\n                    \\\"plan_line_control_rate\\\":0.85,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.13,\\n                    \\\"upper_limit\\\":0.15,\\n                    \\\"plan_line_control_rate\\\":0.8,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.15,\\n                    \\\"upper_limit\\\":1,\\n                    \\\"plan_line_control_rate\\\":0.75,\\n                    \\\"adjust_line_control_rate\\\":0.85,\\n                    \\\"do_decision_when_over_line\\\":true\\n                }\\n            ]\\n        },\\n        \\\"zttvi-p\\\":{\\n            \\\"snode_a\\\":[\\n                {\\n                    \\\"lower_limit\\\":0,\\n                    \\\"upper_limit\\\":0.01,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.06,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.01,\\n                    \\\"upper_limit\\\":0.03,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.04,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.03,\\n                    \\\"upper_limit\\\":0.05,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.02,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.07,\\n                    \\\"upper_limit\\\":0.09,\\n                    \\\"plan_line_control_rate\\\":0.95,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.09,\\n                    \\\"upper_limit\\\":0.11,\\n                    \\\"plan_line_control_rate\\\":0.9,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.11,\\n                    \\\"upper_limit\\\":0.13,\\n                    \\\"plan_line_control_rate\\\":0.85,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.13,\\n                    \\\"upper_limit\\\":0.15,\\n                    \\\"plan_line_control_rate\\\":0.8,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.15,\\n                    \\\"upper_limit\\\":1,\\n                    \\\"plan_line_control_rate\\\":0.75,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                }\\n            ],\\n            \\\"snode_b\\\":[\\n                {\\n                    \\\"lower_limit\\\":0,\\n                    \\\"upper_limit\\\":0.01,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.06,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.01,\\n                    \\\"upper_limit\\\":0.02,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.04,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.02,\\n                    \\\"upper_limit\\\":0.03,\\n                    \\\"plan_line_control_rate\\\":1,\\n                    \\\"adjust_line_control_rate\\\":1.02,\\n                    \\\"do_decision_when_over_line\\\":false\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.04,\\n                    \\\"upper_limit\\\":0.05,\\n                    \\\"plan_line_control_rate\\\":0.95,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.05,\\n                    \\\"upper_limit\\\":0.06,\\n                    \\\"plan_line_control_rate\\\":0.9,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.06,\\n                    \\\"upper_limit\\\":0.07,\\n                    \\\"plan_line_control_rate\\\":0.85,\\n                    \\\"adjust_line_control_rate\\\":0.95,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.07,\\n                    \\\"upper_limit\\\":0.08,\\n                    \\\"plan_line_control_rate\\\":0.8,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                },\\n                {\\n                    \\\"lower_limit\\\":0.08,\\n                    \\\"upper_limit\\\":1,\\n                    \\\"plan_line_control_rate\\\":0.75,\\n                    \\\"adjust_line_control_rate\\\":0.9,\\n                    \\\"do_decision_when_over_line\\\":true\\n                }\\n            ]\\n        }\\n    }\\n}\",\n\t\t\t\"content_md5\": \"4db173498bdd008742384dedaebb91e2\"\n\t\t}\n\t}\n}"
	tmpppp := &SyncConfigResponse{}
	err := json.Unmarshal([]byte(inputStr), tmpppp)
	fmt.Println(err)
	haa := float64(10.5)
	hbb := int64(6)
	hcc := haa / float64(hbb)
	coverName := strings.ReplaceAll("download_zttapk-s", "download_", "")
	fmt.Println(coverName, hcc)
	finalMap := make(map[string]*RangeLimitArrive)
	key, conf := arriveZttvi_p()
	finalMap[key] = conf
	key, conf = arrivezttapk_s()
	finalMap[key] = conf
	key, conf = arriveDefault()
	finalMap[key] = conf
	key, conf = arrivefiletest()
	finalMap[key] = conf
	input := &ArriveRateConfig{
		Switch:           true,
		DecisionInterval: 60,
		CoverMap:         finalMap,
	}
	data, err := json.Marshal(input)
	fmt.Println(string(data), err)
}

func arriveZttvi_p() (string, *RangeLimitArrive) {
	cover := "zttvi-p"
	return cover, &RangeLimitArrive{
		//arrive_rate = 0.7
		//arrive_rate = 0.75
		//arrive_rate = 0.8
		//arrive_rate = 0.85
		ArriveRateList: []float64{0.7, 0.75, 0.8, 0.85},
		//#小于平均带宽的调度线
		//lt_adjust_line_rate = 0.85
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.95
		//lt_adjust_line_rate = 1
		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.95, 1},
		//#大于平均带宽的调度线
		//gt_adjust_line_rate = 0.85
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.95
		//gt_adjust_line_rate = 1
		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.95, 1},
		//#小于平均带宽的规划线
		//lt_plan_line_rate = 0.6
		//lt_plan_line_rate = 0.8
		//lt_plan_line_rate = 0.9
		//lt_plan_line_rate = 1
		LtAvgBwPlanLineRateList: []float64{0.6, 0.8, 0.9, 1},
		//#大于平均带宽的规划线
		//gt_plan_line_rate = 0.5
		//gt_plan_line_rate = 0.7
		//gt_plan_line_rate = 0.8
		//gt_plan_line_rate = 0.9
		GtAvgBwPlanLineRateList: []float64{0.5, 0.7, 0.8, 0.9},
		//#sardine放行总数
		//sardine_pass_count = 2000
		//#静默
		//silent_rate = 0.8
		//#恢复
		//recovery_rate = 0.85
		//#恢复后的调度线
		//recovery_adjust_line_rate = 1.05
		SardinePassCount:       2000,
		SilentRate:             0.8,
		RecoveryRate:           0.85,
		RecoveryAdjustLineRate: 1.05,
	}
}
func arriveDefault() (string, *RangeLimitArrive) {
	cover := "default"
	return cover, &RangeLimitArrive{
		//#到达率
		//arrive_rate = 0.6
		//arrive_rate = 0.65
		//arrive_rate = 0.7
		//arrive_rate = 0.75
		//arrive_rate = 0.8
		ArriveRateList: []float64{0.6, 0.65, 0.7, 0.75, 0.8},
		//#小于平均带宽的调度线
		//lt_adjust_line_rate = 0.85
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.95
		//lt_adjust_line_rate = 1
		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
		//#大于平均带宽的调度线
		//gt_adjust_line_rate = 0.85
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.95
		//gt_adjust_line_rate = 1
		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
		//#小于平均带宽的规划线
		//lt_plan_line_rate = 0.6
		//lt_plan_line_rate = 0.7
		//lt_plan_line_rate = 0.8
		//lt_plan_line_rate = 0.9
		//lt_plan_line_rate = 1
		LtAvgBwPlanLineRateList: []float64{0.6, 0.7, 0.8, 0.9, 1},
		//#大于平均带宽的规划线
		//gt_plan_line_rate = 0.5
		//gt_plan_line_rate = 0.6
		//gt_plan_line_rate = 0.7
		//gt_plan_line_rate = 0.8
		//gt_plan_line_rate = 0.9
		GtAvgBwPlanLineRateList: []float64{0.5, 0.6, 0.7, 0.8, 0.9},
		//#sardine放行总数
		//sardine_pass_count = 1000
		//#静默
		//silent_rate = 0.7
		//#恢复
		//recovery_rate = 0.8
		//#恢复后的调度线
		//recovery_adjust_line_rate = 1.05
		SardinePassCount:       1000,
		SilentRate:             0.7,
		RecoveryRate:           0.8,
		RecoveryAdjustLineRate: 1.05,
	}
}

func arrivefiletest() (string, *RangeLimitArrive) {
	cover := "filetest"
	return cover, &RangeLimitArrive{
		//#到达率
		//arrive_rate = 0.93
		//arrive_rate = 0.94
		//arrive_rate = 0.95
		//arrive_rate = 0.96
		//arrive_rate = 0.97
		ArriveRateList: []float64{0.93, 0.94, 0.95, 0.96, 0.97},
		//#小于平均带宽的调度线
		//lt_adjust_line_rate = 0.85
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.95
		//lt_adjust_line_rate = 1
		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
		//#大于平均带宽的调度线
		//gt_adjust_line_rate = 0.85
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.95
		//gt_adjust_line_rate = 1
		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
		//#小于平均带宽的规划线
		//lt_plan_line_rate = 0.6
		//lt_plan_line_rate = 0.7
		//lt_plan_line_rate = 0.8
		//lt_plan_line_rate = 0.9
		//lt_plan_line_rate = 1
		LtAvgBwPlanLineRateList: []float64{0.6, 0.7, 0.8, 0.9, 1},
		//#大于平均带宽的规划线
		//gt_plan_line_rate = 0.5
		//gt_plan_line_rate = 0.6
		//gt_plan_line_rate = 0.7
		//gt_plan_line_rate = 0.8
		//gt_plan_line_rate = 0.9
		GtAvgBwPlanLineRateList: []float64{0.5, 0.6, 0.7, 0.8, 0.9},
		//#sardine放行总数
		//sardine_pass_count = 1000
		//#静默
		//silent_rate = 0.92
		//#恢复
		//recovery_rate = 0.98
		//#恢复后的调度线
		//recovery_adjust_line_rate = 1.05
		SardinePassCount:       1000,
		SilentRate:             0.92,
		RecoveryRate:           0.98,
		RecoveryAdjustLineRate: 1.05,
	}
}

func arrivezttapk_s() (string, *RangeLimitArrive) {

	//#小于平均带宽的调度线
	//lt_adjust_line_rate = 0.85
	//lt_adjust_line_rate = 0.9
	//lt_adjust_line_rate = 0.9
	//lt_adjust_line_rate = 0.95
	//lt_adjust_line_rate = 1
	//#大于平均带宽的调度线
	//gt_adjust_line_rate = 0.85
	//gt_adjust_line_rate = 0.9
	//gt_adjust_line_rate = 0.9
	//gt_adjust_line_rate = 0.95
	//gt_adjust_line_rate = 1
	//#sardine放行总数
	//sardine_pass_count = 1000
	//#静默
	//silent_rate = 0.7
	//#恢复
	//recovery_rate = 0.8
	//#恢复后的调度线
	//recovery_adjust_line_rate = 1.05
	cover := "zttapk-s"
	return cover, &RangeLimitArrive{
		//#到达率
		//arrive_rate = 0.6
		//arrive_rate = 0.65
		//arrive_rate = 0.7
		//arrive_rate = 0.75
		//arrive_rate = 0.8
		ArriveRateList: []float64{0.6, 0.65, 0.7, 0.75, 0.8},
		//#小于平均带宽的调度线
		//lt_adjust_line_rate = 0.85
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.9
		//lt_adjust_line_rate = 0.95
		//lt_adjust_line_rate = 1
		LtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
		//#大于平均带宽的调度线
		//gt_adjust_line_rate = 0.85
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.9
		//gt_adjust_line_rate = 0.95
		//gt_adjust_line_rate = 1
		GtAvgBwAdjustLineRateList: []float64{0.85, 0.9, 0.9, 0.95, 1},
		//#小于平均带宽的规划线
		//lt_plan_line_rate = 0.6
		//lt_plan_line_rate = 0.7
		//lt_plan_line_rate = 0.8
		//lt_plan_line_rate = 0.9
		//lt_plan_line_rate = 1
		LtAvgBwPlanLineRateList: []float64{0.6, 0.7, 0.8, 0.9, 1},
		//#大于平均带宽的规划线
		//gt_plan_line_rate = 0.5
		//gt_plan_line_rate = 0.6
		//gt_plan_line_rate = 0.7
		//gt_plan_line_rate = 0.8
		//gt_plan_line_rate = 0.9
		GtAvgBwPlanLineRateList: []float64{0.5, 0.6, 0.7, 0.8, 0.9},
		//#sardine放行总数
		//sardine_pass_count = 1000
		//#静默
		//silent_rate = 0.7
		//#恢复
		//recovery_rate = 0.8
		//#恢复后的调度线
		//recovery_adjust_line_rate = 1.05
		SardinePassCount:       1000,
		SilentRate:             0.7,
		RecoveryRate:           0.8,
		RecoveryAdjustLineRate: 1.05,
	}
}
