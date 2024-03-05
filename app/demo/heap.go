package demo

import (
	"fmt"
	"health/util/ticker"
	"math/rand"
	"time"
)

type Xi struct {
	name string
	age  int
}

func shuffleIPs(ips []string) {
	rand.Seed(time.Now().UnixNano())
	// 打乱数组
	rand.Shuffle(len(ips), func(i, j int) {
		ips[i], ips[j] = ips[j], ips[i]
	})
}

func Heap() {
	// 假设有 10 个 IP
	ips := []string{"ip1", "ip2", "ip3", "ip4", "ip5", "ip6", "ip7", "ip8", "ip9", "ip10"}

	xx := Xi{
		name: "aaa",
	}
	list := kkk(xx)
	fmt.Println(list)
	// 轮数
	round := 0

	fmt.Println(time.Now().After(time.Now().Add(-1 * time.Hour)))
	ticker.Ticker(2, func() {

		// 获取异常 IP
		exceptionIPs := ips[:11]

		// 打印当前时间和异常 IP
		fmt.Printf("Round %v, Exception IPs: %v\n", round, exceptionIPs)

		// 将 IP 打散
		shuffleIPs(ips)

		round++
	})
	time.Sleep(1 * time.Hour)
}

func kkk(xx Xi) []Xi {
	res := make([]Xi, 0)
	for i := 0; i < 5; i++ {
		tmp := xx
		tmp.age += i
		res = append(res, xx)

	}
	return res
}
