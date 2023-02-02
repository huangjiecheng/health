package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	//c := cron.New(cron.WithSeconds())
	//_, err := c.AddFunc("0 0 4 * * *", func() {
	//	r.calculateAvgSize()
	//})
	const lll = `{"name":"张三"}`
	a := &Adapter{}
	a.SerializeConfig(lll)

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
