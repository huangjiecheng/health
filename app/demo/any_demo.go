package demo

import (
	"fmt"
	"health/share"
	"health/util/ticker"
	"math"
	"regexp"
	"time"
)

type AnyDemo struct {
	IspInfoList []*IspInfo `json:"isp_info_list"`
}
type IspInfo struct {
	CoverList []CoverUnit `json:"cover_list"`
	Name      string      `json:"name"`
}

type CoverUnit struct {
	Name          share.CoverName   `json:"name"`
	AvgSize       int64             `json:"avg_size"`        // 赋值字段非接口返回：本机计算的平均文件大小
	SmoothAvgSize int64             `json:"smooth_avg_size"` // 赋值字段非接口返回：本机计算的平滑5分钟平均文件大小
	Ceil          float64           `json:"ceil"`            // 赋值字段非接口返回：向上取整阈值，本机获取配置得到
	Tag           map[string]string `json:"mmm"`
	Time          int64             `json:"time"`
}

func RunAnyDemo() {
	startTime := time.Now()

	ccc := &CoverUnit{}
	for k, v := range ccc.Tag {
		fmt.Println(k, v)
	}
	// 记录结束时间
	endTime := time.Now()
	// 计算执行时间
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("huangjiecheng endTime.After(startTime)：%t", endTime.After(startTime))
	fmt.Printf("huangjiecheng 执行耗时：%s", elapsedTime)
}

func RunAnyDemo2() {
	ticker.Ticker(1, func() {
		fmt.Println("xxxxxx========== %s", time.Now().String())
	})
	aaa := math.Sqrt(0)
	fmt.Println(aaa)

	//var (
	//	expireTime = time.Now().Add(5 * time.Second)
	//)
	//fmt.Printf("time.Now(): %s\n", time.Now().String())
	//
	//fmt.Printf("expireTime: %s\n", expireTime.String())
	//
	//for i := 0; i < 100; i++ {
	//	now := time.Now()
	//	fmt.Printf("now: %s", now)
	//
	//	if expireTime.Before(now) {
	//		fmt.Printf("Before\n")
	//	}
	//	if expireTime.After(now) {
	//		fmt.Printf("After\n")
	//	}
	//	time.Sleep(500 * time.Millisecond)
	//}
	//a := "111"
	//switch a {
	//case "11": // 整机到达率
	//	fmt.Println("0000000000")
	//case "111": // 覆盖到达率
	//	fmt.Println("111111111")
	//
	//case "222": // ip到达率
	//	fmt.Println("2222222")
	//default:
	//	fmt.Println("333333333")
	//
	//}
	//var (
	//	//bwGap                = int64(20000000)
	//	ipWeightDecreaseRate = int64(5)
	//	ipWeightBaseRate     = int64(1000)
	//	//adjustLine           = int64(100000000)
	//)
	//dTimes := int64(2)
	//res := ipWeightDecreaseRate / ipWeightBaseRate * dTimes
	//
	//fmt.Println(res)
}

// 判断时间范围是否合法
func isValidTimeRange(startTimeStr, endTimeStr string) bool {
	var (
		startTime, startOk = isValidTimeFormat(startTimeStr)
		endTime, endOk     = isValidTimeFormat(endTimeStr)
	)
	return startOk && endOk && startTime.Before(endTime)
}

// 判断时间格式是否合法
func isValidTimeFormat(timeStr string) (time.Time, bool) {
	if len(timeStr) == 0 {
		return time.Time{}, false
	}
	// 定义 hh:mm 格式的正则表达式
	timeRegex := regexp.MustCompile(`^([01][0-9]|2[0-3]):[0-5][0-9]$`)

	// 使用正则表达式匹配时间格式
	if !timeRegex.MatchString(timeStr) {
		return time.Time{}, false
	}

	// 将字符串解析为时间
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, false
	}

	return t, true
}

// 判断当前时间是否在指定时间范围内
func isCurrentTimeInRange(startTimeStr, endTimeStr string) bool {
	if !isValidTimeRange(startTimeStr, endTimeStr) {
		return false
	}
	// 获取当前时间
	currentTime := time.Now()
	startTime, _ := isValidTimeFormat(startTimeStr)
	endTime, _ := isValidTimeFormat(endTimeStr)

	// 提取小时和分钟部分，忽略日期
	currentTime = time.Date(0, 1, 1, currentTime.Hour(), currentTime.Minute(), 0, 0, time.UTC)
	startTime = time.Date(0, 1, 1, startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
	endTime = time.Date(0, 1, 1, endTime.Hour(), endTime.Minute(), 0, 0, time.UTC)

	// 判断当前时间是否在范围内
	return currentTime.After(startTime) && currentTime.Before(endTime)
}
