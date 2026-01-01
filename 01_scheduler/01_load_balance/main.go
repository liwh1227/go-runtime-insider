// 01_scheduler/01_load_balance/main.go
package main

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 1. 限制 P 的数量，制造“僧多粥少”的局面，强制调度器工作
	// 设置为 2，方便我们在日志中观察 P0 和 P1 的负载
	runtime.GOMAXPROCS(2)

	// 2. 启动大量的 Goroutine (超过 P 的数量)
	// 模拟高并发的交易验签请求
	const numGoroutines = 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	fmt.Println("🚀 Experiment Start: Generating CPU Load...")
	fmt.Printf("🔧 GOMAXPROCS: %d, Goroutines: %d\n", runtime.GOMAXPROCS(0), numGoroutines)
	fmt.Println("👀 Please watch the SCHED traces below...")

	// 启动工作协程
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			// 模拟计算密集型任务 (如: 区块哈希计算)
			// 这种任务不会主动让出 CPU (除非被抢占)
			start := time.Now()
			for time.Since(start) < 2*time.Second {
				// 疯狂做哈希运算，燃烧 CPU
				data := fmt.Sprintf("block-data-%d-%d", id, time.Now().UnixNano())
				sha256.Sum256([]byte(data))
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Experiment Finished.")
}
