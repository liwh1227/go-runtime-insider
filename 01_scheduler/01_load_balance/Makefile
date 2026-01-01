# 01_scheduler/01_load_balance/Makefile
.PHONY: run clean

# GODEBUG=schedtrace=1000: 每 1000ms 打印一次调度器状态
run: build
	@echo "--- Running with Scheduler Tracing ---"
	GODEBUG=schedtrace=1000 ./load_balance

build:
	go build -o load_balance main.go

clean:
	rm -f load_balance