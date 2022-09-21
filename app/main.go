package main

import (
	"fmt"
	"health/router"
	"math/rand"
	"net"
	"reflect"
	"sort"
	"strconv"
	"time"
)

var (
	ip2one    = make(map[string]int64)
	ip2second = make(map[string]int64)
)

type AffinityState struct {
	ipCard string
	ttl    int64
	ipMap  map[string][]string
}

type cacheName string
type cardIp string

var cacheMap = make(map[cacheName][]cardIp)

var affinityMgr = make(map[string][]*AffinityState)

var affinityMgr2 = make(map[string][]AffinityState)

func main() {
	//testFunc()
	//testFunc2()
	//testFunc3()
	//testFunc4()
	//testFunc5()
	//testFunc6()
	fmt.Printf("时间1111：%s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("时间2222：%s\n", time.Now().UTC())

	router.Init()
}

func testFunc6() {
	temp := make([]*AffinityState, 0, 3)

	aa := &AffinityState{
		ipCard: "9.9.9.9",
		ttl:    3,
	}
	bb := &AffinityState{
		ipCard: "2.2.2.2",
		ttl:    3,
	}
	cc := &AffinityState{
		ipCard: "5.5.5.5",
		ttl:    3,
	}
	temp = append(temp, aa, bb, cc)

	ipStr := "127.0.0.1"
	ip := net.ParseIP(ipStr)
	if ipv4 := ip.To4(); ipv4 != nil {
		// 就是ipv6
	}
	//tf1(temp)
	//tf2(temp)

	for _, cacheGroupInfo := range temp {
		ee := fmt.Sprintf("aaaa%+v", *cacheGroupInfo)
		fmt.Printf(ee)
	}
	ff := fmt.Sprintf("kkkkk %+v", temp)
	fmt.Printf(ff)

	//fmt.Printf("%+v, %s", temp, ee)

}

func tf1(temp []AffinityState) {
	for _, info := range temp {
		info.ttl = info.ttl - 1
	}
}
func tf2(temp []AffinityState) {
	for _, info := range temp {
		info.ttl = info.ttl + 10
	}
}

func testFunc5() {
	temp := make([]*AffinityState, 0, 3)

	aa := &AffinityState{
		ipCard: "9.9.9.9",
		ttl:    3,
	}
	bb := &AffinityState{
		ipCard: "2.2.2.2",
		ttl:    3,
	}
	cc := &AffinityState{
		ipCard: "5.5.5.5",
		ttl:    3,
	}
	temp = append(temp, aa, bb, cc)

	cardIpToAbilityMap := make(map[string]*AffinityState, 3)
	for _, ipAb := range temp {
		cardIpToAbilityMap[ipAb.ipCard] = ipAb
	}
	affinityCardIps := []string{"2.2.2.2", "9.9.9.9"}
	for _, ip := range affinityCardIps {
		ipAb, ok := cardIpToAbilityMap[ip]
		if !ok {
			continue
		}
		ipAb.ttl = 200
	}
	fmt.Printf("%+v", cardIpToAbilityMap)

}
func testFunc4() {
	temp := make(map[string][]AffinityState)
	testMapSave(temp)
	testMapSave2(temp)

	aa := AffinityState{
		ipCard: "9.9.9.9",
	}
	testObjSave3(aa)
	affinityMgr2 = temp
	fmt.Printf("hhhhhhhhh %#v", temp)

}

func testObjSave3(aa AffinityState) {
	aa.ipCard = "0"
}

func testMapSave(temp map[string][]AffinityState) {
	temp["aa"] = append(temp["aa"], AffinityState{
		ipCard: "1.2.3",
	})
}

func testMapSave2(temp map[string][]AffinityState) {
	temp["aa"] = append(temp["aa"], AffinityState{
		ipCard: "2.2.2.2",
	})
	temp["bb"] = append(temp["bb"], AffinityState{
		ipCard: "66666",
	})
}

func testFunc3() (result map[string]*AffinityState) {
	obj1 := &AffinityState{
		ipCard: "1.1.1.1",
		ttl:    100,
	}
	obj2 := &AffinityState{
		ipCard: "2.2.2.2",
		ttl:    200,
	}
	obj3 := &AffinityState{
		ipCard: "3.3.3.3",
		ttl:    300,
	}
	obj4 := &AffinityState{
		ipCard: "4.4.4.4",
		ttl:    400,
	}
	affinityMgr["adsdas"] = make([]*AffinityState, 0)

	affinityMgr["aaa"] = []*AffinityState{
		obj1,
		obj4,
		obj3,
		obj2,
	}
	affinityMgr["bbb"] = []*AffinityState{
		obj1,
		obj3,
		obj2,
		obj4,
	}
	sort.Slice(affinityMgr["aaa"], func(i, j int) bool {
		return affinityMgr["aaa"][i].ipCard > affinityMgr["aaa"][j].ipCard
	})
	for name := range affinityMgr {
		sort.Slice(affinityMgr[name], func(i, j int) bool {
			return affinityMgr[name][i].ipCard < affinityMgr[name][j].ipCard
		})
	}
	affinities := make([]*AffinityState, 0)
	for k, v := range affinityMgr {

		for _, obj := range v {
			if k == "aaa" {
				o := obj
				affinities = append(affinities, o)
			}
		}
	}
	for _, item := range affinityMgr["aaa"] {
		if item.ipMap == nil {
			item.ipMap = make(map[string][]string)
		}
		item.ipMap["aaa"] = append(item.ipMap["aaa"], "hahah")
	}

	modifyFunc(result)

	aaa := fmt.Sprintf("输出map： %#v\n", affinities[0])
	fmt.Printf(aaa)
	fmt.Printf("输出map： %#v\n", affinities)

	cacheMap["aa"] = append(cacheMap["aa"], "123")
	cacheMap["bb"] = append(cacheMap["bb"], "456")

	ccc := fmt.Sprintf("输出map2： %#v\n", affinityMgr)
	fmt.Printf(ccc)

	if len(affinityMgr) == 0 {
		fmt.Printf("affinityMgr 长度为0")
	}

	if affinityStates, exist := affinityMgr["aaa"]; exist {
		for _, state := range affinityStates {
			fmt.Printf(state.ipCard)
		}
	}

	var iPAbilitys []AffinityState
	for _, ipAbList := range affinityMgr["aaa"] {

		iPAbilitys = append(iPAbilitys, *ipAbList)
		ipAbList.ttl = 666
	}
	funxxx(iPAbilitys)
	fmt.Printf("666")
	fmt.Printf("%v\n", result)
	return

}

func Infof(format string, v ...interface{}) {
}

func modifyFunc(result map[string]*AffinityState) {
	result = make(map[string]*AffinityState)

	result["aaa"] = &AffinityState{ipCard: "111"}
	result["bbb"] = &AffinityState{ipCard: "222"}
}

func funxxx(abilitys []AffinityState) (result map[string][]*AffinityState) {
	abilitys[0].ipCard = "hhhhhxxxx"

	testArr := make(map[string][]*AffinityState)
	hhh := make([]AffinityState, 0)
	for _, ipAb := range abilitys {
		newIPAbilities := ipAb
		newIPAbilities.ttl = 112233
		testArr["aaa"] = append(testArr["aaa"], &newIPAbilities)
		hhh = append(hhh, newIPAbilities)

	}
	delete(testArr, "aaa")
	hhh[0].ttl = 6
	return
}

func testFunc2() {
	adjust := int64(1000)
	bandwidth := int64(200)
	dValue := adjust - bandwidth
	lastDValue := int64(0)
	floatingPercentage, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", float64(dValue)/float64(lastDValue)), 64)
	if dValue*lastDValue > 0 && (1-0.01) <= floatingPercentage && floatingPercentage <= (1+0.01) {
		fmt.Printf("222")

	}
	fmt.Printf("%f", 0.0211111111)

}
func testFunc() {
	bw := int64(100)
	adjustLine := int64(1000)
	dValue := adjustLine - bw
	ipWeightIncreaseRate := 5
	ipWeightBaseRate := 1000
	dTimes := dValue*5/adjustLine + 1
	weight := dValue * (int64(ipWeightIncreaseRate)) / int64(ipWeightBaseRate) * dTimes
	fmt.Printf(string(weight))

	testMapFunc()
	fmt.Printf("111")
	//go func() {
	//	eventsTick := time.NewTicker(time.Duration(1) * time.Second)
	//	defer eventsTick.Stop()
	//
	//	for {
	//		select {
	//		case <-eventsTick.C:
	//			//timerFunc()
	//			//testMapFunc()
	//		}
	//	}
	//}()
}

func testMapFunc() {
	ip2one["aaa"] = 111
	ip2one["bbb"] = 222
	ip2second["ccc"] = 333
	ip2second["ddd"] = 444
	tempOne := make(map[string]int64)
	tempTwo := make(map[string]int64)
	for i := 0; i < 10; i++ {
		tempOne[getRangString(5)] = int64(rand.Intn(10))
		tempTwo[getRangString(5)] = int64(rand.Intn(10))
	}
	ip2one = tempOne
	ip2second = tempTwo
}

func timerFunc() {
	for i := 0; i < 200; i++ {
		randStr := getRangString(5)
		randStr2 := getRangString(10)
		one := ip2one[randStr]
		two := ip2second[randStr2]
		one++
		one++
		one = one + rand.Int63n(5)
		two++
		two++
		two++
		ip2one[randStr] = one
		ip2second[randStr2] = two
		fmt.Printf(randStr + ":" + string(one) + ":" + string(ip2one[randStr]))
		fmt.Printf(randStr2 + ":" + string(two) + ":" + string(ip2second[randStr2]))
	}

}

func getRangString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Contains(array interface{}, val interface{}) bool {
	targetValue := reflect.ValueOf(array)
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == val {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(val)).IsValid() {
			return true
		}
	}
	return false
}
