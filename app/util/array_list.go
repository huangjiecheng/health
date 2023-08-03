package util

import "fmt"

type ArrayList struct {
	index   int64   // 当前索引
	number  int64   //当前数量
	length  int64   // 长度
	history []int64 // 历史列表
}

func NewArrayList(length int64) *ArrayList {
	return &ArrayList{
		index:   0,
		number:  0,
		length:  length,
		history: make([]int64, length, length),
	}
}

func (a *ArrayList) Append(i int64) {
	fmt.Println(fmt.Sprintf("0000000000000000000:【%d】", a.index))
	// 赋值
	a.history[a.index] = i
	// 索引下标移位，并判断长度防止溢出
	if a.index++; a.index == a.length {
		a.index -= a.length
	}
	// 更新数量
	if a.number < a.length {
		a.number++
	}
}

func (a *ArrayList) RangeAll() (all, number int64) {
	for i := range a.history {
		all += a.history[i]
	}
	number = a.number
	return
}

func (a *ArrayList) GetNumber() int64 {
	return a.number
}
