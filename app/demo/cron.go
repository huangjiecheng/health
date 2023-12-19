package demo

import (
	"fmt"
	"health/util/ticker"
	"math/rand"
	"time"
)

func Cron() {
	ticker.Ticker(2, func() {
		start := time.Now()
		ms := rand.Intn(300)
		time.Sleep(time.Duration(ms) * time.Millisecond)
		end := time.Now()
		fmt.Println(fmt.Sprintf("now: %s \tsleep: %d \tcost: %s", time.Now(), ms, end.Sub(start)))
	})
}
