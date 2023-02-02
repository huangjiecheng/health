package util

import (
	"container/list"
)

type LimitQueue struct {
	capacity  int
	queueList *list.List
}

// Set 设置元素
func (l *LimitQueue) Set(value interface{}) {
	if l.queueList == nil {
		return
	}

	l.queueList.PushFront(value)
	if l.queueList.Len() > l.capacity {
		lastElement := l.queueList.Back()
		if lastElement == nil {
			return
		}
		l.queueList.Remove(lastElement)
	}
}

// RangeLimitQueue 遍历元素
func (l *LimitQueue) RangeLimitQueue(f func(value interface{})) {
	for e := l.queueList.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}

// NewLimitQueue 创建一个LimitQueue对象
func NewLimitQueue(size int) *LimitQueue {
	limitQueue := LimitQueue{capacity: size}
	limitQueue.queueList = list.New()
	return &limitQueue
}
