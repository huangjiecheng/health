package ttl_map

import (
	"health/util/ticker"
	"sync"
	"time"
)

const (
	defaultCleanTick = 10 * 60 // 默认清理过期数据间隔时间（秒）
	defaultTtl       = 60      // 默认key的存活时间（秒）
)

type TTLMap struct {
	mux sync.RWMutex
	m   map[interface{}]*data
	ttl int64
}

type data struct {
	value  interface{}
	expire int64
}

func (tm *TTLMap) Get(key interface{}) (interface{}, bool) {
	tm.mux.RLock()
	defer tm.mux.RUnlock()
	now := time.Now().Unix()
	if i, exist := tm.m[key]; exist && i.expire > now {
		return i.value, true
	}
	return nil, false
}

func (tm *TTLMap) Set(key, val interface{}, ttl ...int64) {
	tm.mux.Lock()
	defer tm.mux.Unlock()
	var (
		now    = time.Now().Unix()
		expire = int64(0)
	)
	if len(ttl) > 0 {
		if ttl[0] > 0 {
			expire = now + ttl[0]
		}
	} else {
		expire = now + defaultTtl
	}
	tm.m[key] = &data{value: val, expire: expire}
}

func (tm *TTLMap) clean() {
	tm.mux.Lock()
	defer tm.mux.Unlock()
	now := time.Now().Unix()
	for k, v := range tm.m {
		// 如果之前设置的 ttl <= 0 则认为永久保留
		if v.expire <= 0 {
			continue
		}
		if v.expire < now {
			delete(tm.m, k)
		}
	}
}

func New(interval ...int) *TTLMap {
	tm := &TTLMap{
		m: make(map[interface{}]*data),
	}
	if len(interval) > 0 && interval[0] > 0 {
		ticker.Ticker(interval[0], tm.clean)
		return tm
	}
	ticker.Ticker(defaultCleanTick, tm.clean)
	return tm
}
