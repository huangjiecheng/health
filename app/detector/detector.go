package detector

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

type PingResult struct {
	SourceIP   string
	TargetIP   string
	PacketLoss float64
	Stat       *ping.Statistics
}

func pingWorker(sourceIP, targetIP string, results chan<- PingResult, wg *sync.WaitGroup) {
	pinger, err := ping.NewPinger(targetIP)
	if err != nil {
		log.Printf("Error creating pinger for %s from %s: %s\n", targetIP, sourceIP, err)
		return
	}

	pinger.Count = 10
	pinger.Timeout = 2 * time.Second
	pinger.Interval = 100 * time.Millisecond

	// 设置 ICMP 包大小为最小值 (8 bytes)
	pinger.Size = 24

	err = pinger.Run()
	if err != nil {
		fmt.Printf("Error pinging %s from %s: %s\n", targetIP, sourceIP, err)
		return
	}

	stats := pinger.Statistics()

	result := PingResult{
		SourceIP: sourceIP,
		TargetIP: targetIP,
		Stat:     stats,
	}

	results <- result
}

func pingTargets(sourceIPs, targetIPs []string, results chan<- PingResult, numWorkers int, wg *sync.WaitGroup) {
	workerPool := make(chan struct{}, numWorkers)

	for _, sourceIP := range sourceIPs {
		for _, targetIP := range targetIPs {
			wg.Add(1)
			workerPool <- struct{}{} // 占用一个协程池的工作协程

			go func(sourceIP, targetIP string) {
				defer func() {
					<-workerPool // 释放协程池的工作协程
					wg.Done()
				}()
				pingWorker(sourceIP, targetIP, results, wg)
			}(sourceIP, targetIP)
		}
	}
}

func Start() {

	sourceIPs := []string{"172.20.124.10"}                                                                                                                                                          // 你的源 IP 列表
	targetIPs := []string{"172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32", "172.20.124.32"} // 你的目标 IP 列表

	results := make(chan PingResult)
	numWorkers := 3 // 设置协程池中的工作协程数量

	// 创建全局的 WaitGroup
	var wg sync.WaitGroup

	go pingTargets(sourceIPs, targetIPs, results, numWorkers, &wg)

	for item := range results {
		fmt.Println(fmt.Sprintf("sIp: %s tIp: %s res: %+v", item.SourceIP, item.TargetIP, item.Stat))
	}

	wg.Wait()
}
