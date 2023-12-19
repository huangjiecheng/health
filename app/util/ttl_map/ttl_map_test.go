package ttl_map

import (
	"testing"
)

// 注意：在这里导入你的 TTL Map 包

func BenchmarkTTLMap_Set(b *testing.B) {
	tm := New()
	keys := make([]int, b.N)
	values := make([]int, b.N)

	// 初始化测试数据
	for i := 0; i < b.N; i++ {
		keys[i] = i
		values[i] = i * 2
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idx := b.N - 1
			key := keys[idx]
			val := values[idx]
			tm.Set(key, val, 60) // 使用适当的 TTL 值
		}
	})
}

func BenchmarkTTLMap_Get(b *testing.B) {
	tm := New()
	keys := make([]int, b.N)
	values := make([]int, b.N)

	// 初始化测试数据
	for i := 0; i < b.N; i++ {
		keys[i] = i
		values[i] = i * 2
		tm.Set(keys[i], values[i], 60) // 使用适当的 TTL 值
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			idx := b.N - 1
			key := keys[idx]
			val, _ := tm.Get(key)
			_ = val
		}
	})
}
