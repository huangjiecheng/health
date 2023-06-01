package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Stutter struct {
	Num  int64  `json:"num"`
	Name string `json:"name"`
}

type Decision interface {
	Modify()
	Print()
}

type Factory struct {
	sync.RWMutex
	decisions []Decision
}

type DemoDecision struct {
	Abc int64 `json:"abc"`
}

func (d *DemoDecision) Modify() {
	d.Abc = 666
}
func (d *DemoDecision) Print() {
	fmt.Println(d.Abc)
}

const input = `{"num":123,"name":"zhangsan"}`

var (
	decisionRound  = int64(-1) // 当前决策轮数
	qualityFactory = Factory{
		decisions: make([]Decision, 0),
	}
)

type Adapter struct {
	Config *Config `json:"config"`
}

type Config struct {
	Name string `json:"name"`
	Num  int64  `json:"num"`
}

func (s *Stutter) Validate() {

}

type IpAb struct {
	DownloadSpeedDecision decisionResult `json:"download_speed_decision"` // 下载速度控线结果
}

type decisionResult struct {
	AdjustLine            int64   `json:"adjust_line"`              // 控线结果
	AdjustLineControlRate float64 `json:"adjust_line_control_rate"` // 调度线的控线比例
	PlanLineControlRate   float64 `json:"plan_line_control_rate"`   // 规划线的控线比例
	Success               bool    `json:"success"`                  // 控线是否成功
}

type DecisionResultHjc struct {
	AdjustLine            int64   `json:"adjust_line"`              // 控线结果
	AdjustLineControlRate float64 `json:"adjust_line_control_rate"` // 调度线的控线比例
	PlanLineControlRate   float64 `json:"plan_line_control_rate"`   // 规划线的控线比例
	Success               bool    `json:"success"`                  // 控线是否成功
}

func main() {

	demo := &DecisionResultHjc{
		AdjustLine: 111,
		Success:    true,
	}

	tmp := *demo
	demo1 := &tmp
	demo1.AdjustLine = 222
	tmp2 := *demo
	demo2 := &tmp2
	demo2.AdjustLine = 222
	fmt.Println("1111")
	//c := cron.New(cron.WithSeconds())
	//_, err := c.AddFunc("0 0 4 * * *", func() {
	//	r.calculateAvgSize()
	//})aa
	aaa := "a,b,g,h,c,d,e"
	bbb := strings.Split(aaa, ",")
	for _, s := range bbb {
		fmt.Println(fmt.Sprintf("1111111: %s", s))
	}
	decisionResults := make([]decisionResult, 0, 5)
	decisionResults = append(decisionResults, decisionResult{
		AdjustLine: 333,
	})
	decisionResults = append(decisionResults, decisionResult{
		AdjustLine: 555,
	})
	decisionResults = append(decisionResults, decisionResult{
		AdjustLine: 444,
	})
	decisionResults = append(decisionResults, decisionResult{
		AdjustLine: 111,
	})
	decisionResults = append(decisionResults, decisionResult{
		AdjustLine: 222,
	})
	sort.Slice(decisionResults, func(i, j int) bool {
		return decisionResults[i].AdjustLine < decisionResults[j].AdjustLine
	})
	for _, result := range decisionResults {
		fmt.Println(result.AdjustLine)
	}
	var globalTraffic sync.Map
	globalTraffic.Store("123", &Config{Num: 3})

	v, ok := globalTraffic.Load("123")
	if !ok {
		v = &Config{}
		globalTraffic.Store("123", v)
	}
	stat := v.(*Config)
	stat.Num += 5
	v2, ok2 := globalTraffic.Load("123")
	if ok2 {
		fmt.Println(v2)
	}
	ipAb := &IpAb{}
	ipAb.DownloadSpeedDecision.AdjustLine = 222
	func111(ipAb)
	ipAb.DownloadSpeedDecision.AdjustLine = 333

}

func func111(ab *IpAb) {
	ab.DownloadSpeedDecision = decisionResult{
		AdjustLine: 111,
	}
}

// RegisterDecision 注册新的质量策略
func RegisterDecision(decision Decision) {
	qualityFactory.Lock()
	qualityFactory.decisions = append(qualityFactory.decisions, decision)
	qualityFactory.Unlock()
}

func (a *Adapter) SerializeConfig(content string) error {
	a.Config = &Config{}
	err := json.Unmarshal([]byte(content), a.Config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal. err: %v", err)
	}
	return nil
}
