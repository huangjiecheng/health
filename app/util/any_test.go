package util

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func Test_asdas(t *testing.T) {
	c := cron.New()
	aaa, err := c.AddFunc("* * * * *", func() {
		printLog()
	})
	if err != nil {
		fmt.Errorf("ssss %v", aaa)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(500)
	}
}

func printLog() {
	fmt.Printf("时间： %d\n", time.Now().Unix())

}
