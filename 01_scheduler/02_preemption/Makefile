# 01_scheduler/02_preemption/Makefile
.PHONY: run view clean

run:
	go build -o preemption main.go
	./preemption

# 生成 trace 后，使用此命令打开可视化页面 (需要 Chrome 浏览器)
view:
	go tool trace trace.out

clean:
	rm -f preemption trace.out