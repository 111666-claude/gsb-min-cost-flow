# min-cost-flow

最小费用最大流（连续最短路增广 + 势函数），只用 Go 标准库。

```
go test ./...
go run ./cmd/min-cost-flow --nodes 4 --edges '0-1:1:1,0-2:1:1,1-3:1:1,2-3:1:1' --source 0 --sink 3
go run ./cmd/min-cost-flow --nodes 4 --edges '0-1:3:1,1-3:3:1,0-2:3:2,2-3:3:2' --source 0 --sink 3
```

## 口径（README 为准）

- **要一直增广到不能再增广**：每轮沿残量网络的最短路增广一条路径，重复到汇点不可达为止；
  只做一轮就返回是错的。`--nodes 4 --edges '0-1:1:1,0-2:1:1,1-3:1:1,2-3:1:1'` 按口径是 `flow=2 cost=4`，
  现在只增广一次，给 `flow=1 cost=2`。
- **平行边要累加**：同一对 `from-to` 出现多次时容量相加、各自保留费用；现在用 map 赋值直接覆盖。
- **负费用要能处理**：残量网络会出现负费用边，要用 SPFA/Bellman-Ford 或先用势函数初始化再跑 Dijkstra。
- 不变量：结果满足容量约束与流量守恒；`flow` 是最大流，且在所有最大流里 `cost` 最小
  （与「小图上枚举所有可行流」的暴力实现对照，差异必须为 0）；同一输入重复求解结果相同。
- 规模：1000 个顶点、1 万条边 3 秒内算完，内存 O(V+E)；自环与零容量边忽略，越界端点报 `ErrInput`。

## 现在的行为

- `Solve` 只有一轮 Bellman-Ford 加一次增广：除了一条直链之外，`flow` 通常只是最大流的一部分。
- `capacity[[2]int{...}] = edge.Cap` 让平行边互相覆盖。
- 没有势函数，多轮增广时负费用残量边会让最短路失效。

## 计划新增的能力

```
--nodes 节点数
--edges 边表    逗号分隔的 from-to:cap:cost
--source / --sink
```

## 输出

```
flow=2 cost=4
```

## 目录

```
mcmf.go                 求解器
cmd/min-cost-flow       命令行入口
mcmf_test.go            go test 用例
```
