package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"health/router"
	"health/share"
	"strings"
	"sync"
	"time"
)

type AffinityInfo struct {
	CardIp       share.CardIp
	IpCard       share.IpCard
	UsedWeight   int64
	Priority     int64
	LastUsedTime int64
	MapDemo      map[string]string
}

type TrafficRespByBlb struct {
	Duration  float64                `json:"duration"`
	DomainMap map[string]interface{} `json:"domain_map"`
	Now       float64                `json:"now"`
}
type Source struct {
	Parent TrafficDetail `json:"parent"`
	Client TrafficDetail `json:"client"`
}

type TrafficDetail struct {
	UpTraffic int64 `json:"up_traffic"`
	UpBytes   int64 `json:"up_bytes"`
	Traffic   int64 `json:"traffic"`
	Bytes     int64 `json:"bytes"`
}

const (
	coverViewCacheDelimiter      = "_:_"
	minAffinityPriorityThreshold = int64(-15)
)

var (
	globalTraffic     sync.Map
	globalAffinityMgr = make(map[string][]AffinityInfo, 0)
	ip1               = AffinityInfo{
		CardIp:     "ip1",
		Priority:   0,
		UsedWeight: 1,
	}
	ip2 = AffinityInfo{
		CardIp:     "ip2",
		Priority:   0,
		UsedWeight: 0,
	}
	ip3 = AffinityInfo{
		CardIp:     "ip3",
		Priority:   0,
		UsedWeight: 0,
	}
	ip4 = AffinityInfo{
		CardIp:     "ip4",
		Priority:   0,
		UsedWeight: 0,
	}
)

type (
	StatisticResp struct {
		Status int            `json:"status"`
		Msg    string         `json:"msg"`
		Data   *StatisticUnit `json:"data"`
	}
	StatisticUnit struct {
		Outs       map[share.CardIp]OutIPUnit `json:"outs"`
		Ins        map[share.CardIp]InIPUnit  `json:"ins"`
		BackSrc    BackSrcUnit                `json:"back_src"`
		BackClient BackClientUnit             `json:"back_client"`
		Traffic    map[string]TrafficUnit     `json:"traffic"`
	}

	TrafficUnit struct {
		Source StatUnit `json:"source"`
		Parent StatUnit `json:"parent"`
		Client StatUnit `json:"client"`
	}

	StatUnit struct {
		CodeStat map[string]int64 `json:"codestat"`
		Speed    int64            `json:"speed"`
		Req      int64            `json:"req"`
		Resp     int64            `json:"resp"`
		AvgTime  float64          `json:"avg_time"`
		AvgSize  int64            `json:"avg_size"`
		Active   int64            `json:"active"`
		HitStat  HitStat          `json:"hitstat"`
	}

	HitStat struct {
		ParentHit     int64   `json:"parent-hit"`
		Hit           int64   `json:"hit"`
		MissRate      float64 `json:"miss_rate"`
		ParentHitRate float64 `json:"parenthit_rate"`
		HitRate       float64 `json:"hit_rate"`
		Miss          int64   `json:"miss"`
	}

	OutIPUnit struct {
		Success SuccessUnit `json:"success"`
		Fail    FailUnit    `json:"fail"`
		Speed   int64       `json:"speed"`
	}

	InIPUnit struct {
		Success SuccessUnit `json:"success"`
	}

	SuccessUnit struct {
		All  int64 `json:"all"`
		Hit  int64 `json:"hit"`
		Miss int64 `json:"miss"`
	}

	FailUnit struct {
		All      int64            `json:"all"`
		HttpCode map[string]int64 `json:"http_code"`
	}

	BackSrcUnit struct {
		UpstreamRequest  int64 `json:"upstream_request"`
		UpstreamResponse int64 `json:"upstream_response"`
		BwIn             int64 `json:"bandwidth_in"`
		Speed            int64 `json:"speed"`
	}

	BackClientUnit struct {
		Time           int64   `json:"time"`             //更新时间 秒级
		Active         int64   `json:"active"`           //当前请求数
		Speed          int64   `json:"speed"`            //软件下载速度 Bps
		LogSpeed       int64   `json:"log_speed"`        //日志下载速度 Bps
		AvgSize        int64   `json:"avg_size"`         //日志下载速度 Bps
		SlowRespNum    int64   `json:"slow_resp_num"`    // 1分钟内响应次数
		SlowSpeedRatio float64 `json:"slow_speed_ratio"` // 1分钟内慢速比
	}
)
type TrafficInfo struct {
	Traffic    int64 `json:"traffic"`
	HitCount   int64 `json:"hit_count"`
	MissCount  int64 `json:"miss_count"`
	TotalCount int64 `json:"total_count"`
}
type Adapter struct {
	Option Option
	Info   Info
}
type Info struct {
	CurrentRound   int64   `json:"current_round"`
	PacketLostRate float64 `json:"packet_lost_rate"`
}

type Option struct {
	Switch            bool
	Period            int
	SilentRoundNumber int
	MajorCover        share.CoverName
}

type bandwidthCell struct {
	Out    int64            `json:"out"`
	In     int64            `json:"in"`
	IPv4   share.CardIp     `json:"ip"`
	IPv6   share.CardIp     `json:"ipv6"`
	IpList []*bandwidthCell `json:"ipList"`
}
type IpBasic struct {
	Priority string
	SNode    string
	Ip       string
}

func (ib IpBasic) IsOk() bool {
	return ib.Ip != "" && ib.SNode != ""
}

func main() {
	//aaa := make(map[share.CoverName]map[share.CardIp]*IpBasic, 0)
	//funxxx(aaa)
	//newAffinityMgr := make(map[string][]AffinityInfo, 0) // 当前轮的亲和性列表
	//newAffinityMgr["cover1_:_view1_:_cache1"] = append(newAffinityMgr["cover1_:_view1_:_cache1"], ip1)
	//integrateAffinity(newAffinityMgr)
	c := cron.New(cron.WithSeconds())
	_, _ = c.AddFunc("*/5 * * * * *", func() {
		fmt.Printf("时间： %d, 准确: %s\n", time.Now().Unix(), time.Now().String())
		fmt.Printf("五分钟前的时间：%d\n", ((time.Now().Unix()/60)-5)*60)
	})
	c.Start()
	router.Init()
}

func funxxx(aaa map[share.CoverName]map[share.CardIp]*IpBasic) {
	ddd := make(map[share.CardIp]*IpBasic)
	ddd[share.CardIp("dsada")] = &IpBasic{
		Ip: "1.1.1.1",
	}
	aaa[share.CoverName("kkk")] = ddd

}

func integrateAffinity(newAffinityMgr map[string][]AffinityInfo) {
	var (
		raisePriorityMap = make(map[string][]AffinityInfo) // cover层、cover-view层有提升亲和性的列表
		now              = time.Now().Unix()
	)
	for coverViewCacheKey, infos := range newAffinityMgr {
		var (
			coverName, viewName, _ = splitCoverViewCache(coverViewCacheKey)
			coverViewKey           = appendCoverView(coverName, viewName)
		)
		// cover-view
		coverViewAffinityInfos, exist := raisePriorityMap[coverViewKey]
		if !exist {
			coverViewAffinityInfos = make([]AffinityInfo, 0)
		}
		// cover
		coverAffinityInfos, exist := raisePriorityMap[string(coverName)]
		if !exist {
			coverAffinityInfos = make([]AffinityInfo, 0)
		}
		for _, info := range infos {
			if info.UsedWeight == 0 {
				continue
			}
			// exp > 0 说明有亲和性
			info.UsedWeight = 0
			info.Priority = 1
			if _, ok := existAffinityCardIp(info.CardIp, coverAffinityInfos); !ok {
				coverAffinityInfos = append(coverAffinityInfos, info)
			}
			if _, ok := existAffinityCardIp(info.CardIp, coverViewAffinityInfos); !ok {
				coverViewAffinityInfos = append(coverViewAffinityInfos, info)
			}
		}
		raisePriorityMap[coverViewKey] = coverViewAffinityInfos
		raisePriorityMap[string(coverName)] = coverAffinityInfos
	}
	// 先将提升亲和性加入
	for key, raiseInfos := range raisePriorityMap {
		oldInfos, exist := globalAffinityMgr[key]
		if !exist {
			newAffinityMgr[key] = raiseInfos
			continue
		}
		for _, raiseInfo := range raiseInfos {
			if oldInfo, ok := existAffinityCardIp(raiseInfo.CardIp, oldInfos); ok {
				raiseInfo.Priority = oldInfo.Priority + 1
			}
			newAffinityMgr[key] = append(newAffinityMgr[key], raiseInfo)
		}
	}
	// 再将降低亲和性加入
	for key, oldInfos := range globalAffinityMgr {
		newAffinityInfos, exist := newAffinityMgr[key]
		if !exist {
			newAffinityInfos = make([]AffinityInfo, 0)
		}
		for _, oldInfo := range oldInfos {
			oldInfo.Priority--
			// 亲和性优先级 < 最低阈值 || 超过存活时间未使用 则去除亲和性
			if oldInfo.Priority < minAffinityPriorityThreshold || (now-int64(oldInfo.LastUsedTime)) > 300 {
				continue
			}
			if _, ok := existAffinityCardIp(oldInfo.CardIp, newAffinityInfos); !ok {
				newAffinityInfos = append(newAffinityInfos, oldInfo)
			}
		}
		newAffinityMgr[key] = newAffinityInfos
	}
}
func existAffinityCardIp(cardIp share.CardIp, infos []AffinityInfo) (AffinityInfo, bool) {
	for _, item := range infos {
		if item.CardIp == cardIp {
			return item, true
		}
	}
	return AffinityInfo{}, false
}
func case2() (newAffinityMgr map[string][]AffinityInfo) {
	newAffinityMgr["cover1_:_view1_:_cache1"] = []AffinityInfo{
		{
			CardIp:     "ip1",
			Priority:   2,
			UsedWeight: 10,
		},
		{
			CardIp:     "ip2",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover1_:_view1_:_cache2"] = []AffinityInfo{
		{
			CardIp:     "ip2",
			Priority:   0,
			UsedWeight: 0,
		},
		{
			CardIp:     "ip3",
			Priority:   1,
			UsedWeight: 10,
		},
		{
			CardIp:     "ip4",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover2_:_view1_:_cache3"] = []AffinityInfo{
		{
			CardIp:     "ip3",
			Priority:   2,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover2_:_view2_:_cache3"] = []AffinityInfo{
		{
			CardIp:     "ip4",
			Priority:   2,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover2_:_view2_:_cache4"] = []AffinityInfo{
		{
			CardIp:     "ip1",
			Priority:   2,
			UsedWeight: 10,
		},
	}
	return newAffinityMgr
}
func case1() map[string][]AffinityInfo {
	newAffinityMgr := make(map[string][]AffinityInfo, 0)
	newAffinityMgr["cover1_:_view1_:_cache1"] = []AffinityInfo{
		{
			CardIp:     "ip1",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover1_:_view1_:_cache2"] = []AffinityInfo{
		{
			CardIp:     "ip2",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover2_:_view1_:_cache3"] = []AffinityInfo{
		{
			CardIp:     "ip3",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover2_:_view2_:_cache3"] = []AffinityInfo{
		{
			CardIp:     "ip4",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	newAffinityMgr["cover2_:_view2_:_cache4"] = []AffinityInfo{
		{
			CardIp:     "ip1",
			Priority:   1,
			UsedWeight: 10,
		},
	}
	return newAffinityMgr
}

func splitCoverViewCache(key string) (share.CoverName, share.ViewName, share.CacheGroupName) {
	values := strings.Split(key, coverViewCacheDelimiter)
	return share.CoverName(values[0]), share.ViewName(values[1]), share.CacheGroupName(values[2])
}

func appendCoverView(coverName share.CoverName, viewName share.ViewName) string {
	return strings.Join([]string{string(coverName), string(viewName)}, coverViewCacheDelimiter)
}

func compareAndInsertAffinity(currentMap map[string][]AffinityInfo,
	coverViewMap map[string]map[share.CardIp]AffinityInfo,
	raiseMap map[string]map[share.CardIp]AffinityInfo) {
	for key, oldAffinityList := range globalAffinityMgr {
		// 已经处理的不再处理
		if _, ok := currentMap[key]; ok {
			continue
		}
		var (
			isRaise       = false                               // 当前key是否存在提升亲和性的列表
			temRaiseIpMap = make(map[share.CardIp]AffinityInfo) // 需要提升亲和性的列表
		)
		for k, item := range raiseMap {
			var (
				coverName, viewName, _ = splitCoverViewCache(k)
				coverViewKey           = appendCoverView(coverName, viewName)
			)
			if key == string(coverName) || key == coverViewKey {
				isRaise = true
				temRaiseIpMap = item
				break
			}
		}
		newAffinityIpInfoMap, exist := coverViewMap[key]
		for _, oldItem := range oldAffinityList {
			// 当前轮亲和性存在并且提升
			if exist && isRaise {
				// 当前cardIp亲和关系存在
				if newItem, ok := newAffinityIpInfoMap[oldItem.CardIp]; ok {
					// 当前cardIp在提升优先级列表中，亲和性+1
					if _, ok2 := temRaiseIpMap[oldItem.CardIp]; ok2 {
						currentMap[key] = append(currentMap[key], buildAffinityInfo(newItem, 0, oldItem.Priority+1))
						continue
					}
				}
			}
			// 亲和性优先级 < 最低阈值 || 超过存活时间未使用 则去除亲和性
			if oldItem.Priority < minAffinityPriorityThreshold || (time.Now().Unix()-int64(oldItem.LastUsedTime)) > 300 {
				continue
			}
			// 降低亲和性-1
			currentMap[key] = append(currentMap[key], buildAffinityInfo(oldItem, 0, oldItem.Priority-1))
		}
	}
}

func buildAffinityInfo(item AffinityInfo, usedWeight, priority int64) AffinityInfo {
	return AffinityInfo{
		CardIp:       item.CardIp,
		IpCard:       item.IpCard,
		UsedWeight:   usedWeight,
		Priority:     priority,
		LastUsedTime: item.LastUsedTime,
	}
}
