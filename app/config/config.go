package config

import (
	"encoding/json"
	"fmt"
	"health/common"
	"health/pkg/quality/arrive_rate"
	"health/pkg/quality/download_speed"
	"health/pkg/quality/packet_lost"
	"health/pkg/quality/stutter"
)

const (
	LOG_PATH = "/opt/cloud/happy/logs"
	LOG_FILE = "happy.log"
)

func LoadConfig() {
	tmp1 := &stutter.Config{}
	serialization([]byte(common.Stutter), tmp1)
	fmt.Println(tmp1)
	tmp2 := &arrive_rate.Config{}
	serialization([]byte(common.ArriveRate), tmp2)
	fmt.Println(tmp2)
	tmp3 := &packet_lost.Config{}
	serialization([]byte(common.PacketLost), tmp3)
	fmt.Println(tmp3)
	tmp4 := &download_speed.Config{}
	serialization([]byte(common.DownloadSpeed), tmp4)
	fmt.Println(tmp4)
}

func serialization(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		v = nil
		return
	}
	return
}
