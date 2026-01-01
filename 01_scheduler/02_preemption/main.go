// 01_scheduler/02_preemption/main.go
package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	// 1. 设置为单核心！
	// 如果抢占失效，这个核心将被 asyncLoop 占死，main goroutine 永远无法继续执行。
	runtime.GOMAXPROCS(1)

	// 2. 开启 Trace
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("🔥 Experiment Start: Testing Asynchronous Preemption")
	fmt.Println("🔧 GOMAXPROCS=1. Expecting 'Async Loop' to be sliced...")

	var wg sync.WaitGroup
	wg.Add(1)

	// 3. 启动“恶霸”协程 (Async Loop)
	// 这是一个纯计算循环，没有任何函数调用，没有 Gosched()。
	// 在 Go 1.13 及以前，这会造成 Deadlock 或 Hang。
	go func() {
		defer wg.Done()
		fmt.Println("😈 Bully Goroutine started. Trying to hog the CPU...")

		// 纯计算死循环，运行 2 秒
		start := time.Now()
		for time.Since(start) < 2*time.Second {
			// 没有任何 IO，没有任何 runtime 调用
			// 纯粹的汇编指令循环
			_ = 1 + 1
		}
		fmt.Println("😈 Bully Goroutine finished.")
	}()

	// 4. 主协程等待
	// 如果抢占生效，主协程应该能有机会打印日志或执行 Wait
	time.Sleep(500 * time.Millisecond)
	fmt.Println("😇 Main Goroutine is running! (Preemption works!)")

	wg.Wait()
	fmt.Println("✅ Experiment Finished.")
}
