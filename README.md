# Go Runtime Insider: The Comprehensive Verification Lab

[![Go Version](https://img.shields.io/badge/Go-1.14%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active_Research-orange.svg)](https://github.com/liwh1227/go-runtime-insider)

<div align="center">
  <a href="#cn">中文说明</a> | <a href="#en">English Readme</a>
</div>

---

<div id="cn"></div>

# Go Runtime 内核透视：确证驱动的深度实验室

> **"Don't trust the docs, verify with data."**
> **"不迷信文档，只相信实证。"**

## 📖 项目愿景 (Vision)

**Go Runtime Insider** 是一个致力于解构 Go 语言运行时（Runtime）内部机制的实验性项目。大多数工程师停留在“背诵八股文”的阶段，而本项目旨在通过**科学实验**和**数据观测**，建立对 Runtime 行为的 **确证能力 (Verification Capability)**。

本项目将 Go Runtime 视为一个待解剖的复杂系统，通过 **黑盒遥测 (Telemetry)**、**灰盒可视化 (Visualization)** 和 **白盒源码注入 (Source Injection)** 三种手段，深入探索以下核心领域：

1.  **调度器 (Scheduler)**: GMP 模型、Work-stealing、异步抢占。
2.  **内存子系统 (Memory Subsystem)**: TCMalloc 分配器、三色标记 GC、混合写屏障。
3.  **核心结构 (Core Structures)**: Swiss Map (Go 1.24+)、Hashed Wheel Timer、Channel 内部机制。

## 🔬 实验模块 (Modules)

本项目按 Runtime 核心组件划分为不同模块，每个模块包含若干个循序渐进的实验。

### 1. 🧠 计算与调度 (Scheduling & GMP)
> *深入理解 Goroutine 的生老病死与 CPU 资源的流转。*

* **[EXP-01] 负载均衡验证**:
    * 利用 `GODEBUG=schedtrace=1000` 观测全局队列与本地队列的压力分布。
    * 区分 CPU 饱和与锁竞争导致的性能瓶颈。
* **[EXP-02] 异步抢占可视化**:
    * 使用 `go tool trace` 捕捉死循环协程被 `SIGURG` 信号强制切分的微秒级证据 (Go 1.14+)。
* **[EXP-03] 窃取路径追踪 (Runtime Hacking)**:
    * **[Hardcore]** 修改 `runtime/proc.go`，注入探针打印 Work-stealing 的精确拓扑（P1 偷了 P2），验证 CPU 亲和性。

### 2. 💾 内存与垃圾回收 (Memory & GC)
> *透视对象在 Heap 上的分配路径与回收机制。*

* **[EXP-04] 分配器行为确证**:
    * 验证 Tiny 对象 (<16B) 与 Small 对象在 `mcache` 中的分配逻辑。
    * 使用 `go build -gcflags="-m"` 结合 Benchmark 验证逃逸分析对性能的影响。
* **[EXP-05] GC Pacing 与压力测试**:
    * 通过 `GODEBUG=gctrace=1` 分析 `p` (Pacing ratio) 的变化。
    * 实测 `GOGC` 与 `GOMEMLIMIT` 在容器环境下的 OOM 防护效果。
* **[EXP-06] 写屏障观测**:
    * **[Hardcore]** 深入混合写屏障 (Hybrid Write Barrier) 源码，验证对象颜色在并发标记阶段的流转。

### 3. 🏗️ 数据结构与并发 (Data Structures)
> *剖析 Go 核心数据结构的演进与设计权衡。*

* **[EXP-07] Swiss Map 性能革命**:
    * 对比 Go 1.24 新版 Swiss Map 与旧版 Map 在高并发读写下的 CPU Cache Miss 率 (使用 `perf`)。
    * 验证 SIMD 指令加速探测的实际收益。
* **[EXP-08] 定时器时间轮 (Timing Wheel)**:
    * 验证 Go 1.14+ 引入的 Netpoller 集成定时器如何消除原本的堆锁竞争。

## 🛠️ 方法论 (Methodology)

本项目严格遵循 **"观察-假设-验证"** 的闭环：

1.  **Level 1: 黑盒 (Blackbox)**
    * 不修改代码，仅通过环境变量 (`GODEBUG`) 和标准 Metrics 观测 Runtime 外部表现。
2.  **Level 2: 灰盒 (Greybox)**
    * 使用 Profiling 工具 (`pprof`, `trace`, `perf`) 查看函数调用栈和系统事件。
3.  **Level 3: 白盒 (Whitebox)**
    * 修改并重新编译 Go Runtime 源码，注入自定义日志和钩子，获取第一手内核数据。

## 🚀 快速开始 (Quick Start)

### 环境准备
```bash
# 推荐使用 Go 1.20+ (部分实验需要 Go 1.24RC)
go version

# 克隆仓库
git clone [https://github.com/yourname/go-runtime-insider.git](https://github.com/yourname/go-runtime-insider.git)
cd go-runtime-insider
```

# 👨‍💻 作者 (Author)
Go Runtime Researcher & System Architect 专注于分布式系统架构与 Go 语言底层原理。致力于通过实验数据揭示软件系统的物理定律。

<div id="en"></div>

# Go Runtime Insider: The Comprehensive Verification Lab

**"Moving beyond theory to empirical evidence."**

## 📖 Vision
**Go Runtime Insider** is an experimental project dedicated to deconstructing the internal mechanics of the Go Runtime. While many engineers stop at memorizing concepts, this project aims to build **Verification Capability** through **scientific experiments** and **empirical observation**.

We treat the Go Runtime as a complex system to be dissected using three methodologies: **Blackbox Telemetry**, **Greybox Visualization**, and **Whitebox Source Injection**, covering the following core areas:

1. **Scheduler**: GMP model, Work-stealing, Asynchronous Preemption.
2. **Memory Subsystem**: TCMalloc allocator, Tri-color GC, Hybrid Write Barrier.
3. **Core Structures**: Swiss Map (Go 1.24+), Hashed Wheel Timer, Channel internals.

## 🔬 Modules
The project is organized into modules corresponding to Go Runtime components.

### 1. 🧠 Scheduling & GMP
_Understanding the lifecycle of Goroutines and CPU resource flow._

+ **[EXP-01] Load Balance Verification**:
    - Observe pressure distribution on Global/Local queues using `GODEBUG=schedtrace=1000`.
    - Distinguish between CPU saturation and lock contention.
+ **[EXP-02] Preemption Visualization**:
    - Capture microsecond-level evidence of tight-loop goroutines being sliced by `SIGURG` signals using `go tool trace` (Go 1.14+).
+ **[EXP-03] Stealing Path Trace (Runtime Hacking)**:
    - **[Hardcore]** Modify `runtime/proc.go` to log the precise topology of Work-stealing, verifying CPU affinity.

### 2. 💾 Memory & GC
_Visualizing object allocation paths and reclamation mechanics._

+ **[EXP-04] Allocator Behavior**:
    - Verify allocation logic for Tiny (<16B) vs. Small objects in `mcache`.
    - Benchmark the impact of Escape Analysis using `go build -gcflags="-m"`.
+ **[EXP-05] GC Pacing & Stress Test**:
    - Analyze `p` (Pacing ratio) dynamics via `GODEBUG=gctrace=1`.
    - Test `GOGC` vs. `GOMEMLIMIT` protection in containerized environments.
+ **[EXP-06] Write Barrier Observation**:
    - **[Hardcore]** Dive into Hybrid Write Barrier source code to verify object coloring flows during concurrent marking.

### 3. 🏗️ Data Structures
_Dissecting design trade-offs of core Go structures._

+ **[EXP-07] Swiss Map Revolution**:
    - Compare CPU Cache Miss rates of Go 1.24 Swiss Map vs. Old Map using `perf`.
    - Verify the actual gains from SIMD probing.
+ **[EXP-08] Timing Wheel**:
    - Verify how Netpoller-integrated timers (Go 1.14+) eliminate heap lock contention.

## 🛠️ Methodology
We strictly follow the **"Observe - Hypothesize - Verify"** loop:

1. **Level 1: Blackbox**
    - Observe external behavior using Env Vars (`GODEBUG`) and Metrics without code modification.
2. **Level 2: Greybox**
    - Inspect call stacks and system events using Profiling tools (`pprof`, `trace`, `perf`).
3. **Level 3: Whitebox**
    - Modify and recompile Go Runtime source code to inject custom logs and hooks for first-hand kernel data.

## 🚀 Quick Start
### Prerequisites
```plain
# Go 1.20+ Recommended (Go 1.24RC for some experiments)
go version

# Clone Repo
git clone [https://github.com/yourname/go-runtime-insider.git](https://github.com/yourname/go-runtime-insider.git)
cd go-runtime-insider
```

## 👨‍💻 Author
**Go Runtime Researcher & System Architect** Focusing on distributed system architecture and Go language internals. Dedicated to revealing the physical laws of software systems through empirical data.

