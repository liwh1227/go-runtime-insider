// 01_scheduler/03_stealing_hack/main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 开启 4 个 P
	runtime.GOMAXPROCS(4)
	fmt.Println("🕵️  Runtime Hacking Experiment: Work Stealing Tracer")
	fmt.Println("🔥 Generating imbalance work to trigger stealing...")

	var wg sync.WaitGroup
	wg.Add(20)

	// 制造 20 个短任务，它们会很快执行完，
	// 导致 P 的本地队列经常空，需要去偷别的 P。
	for i := 0; i < 20; i++ {
		go func(id int) {
			defer wg.Done()
			// 模拟一点工作量
			start := time.Now()
			for time.Since(start) < 100*time.Millisecond {
				_ = 1 * 1
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ Done.")
}
