# Performance lessons

## Profiling with pprof

Run the profiling benchmark long enough to collect useful CPU and allocation data:

```shell
go test ./lesson/performance -run '^$' -bench '^BenchmarkProfileWorkload$' -benchtime=5s -cpuprofile=cpu.out -memprofile=mem.out
go tool pprof -top cpu.out
go tool pprof -top -alloc_space mem.out
```

`ProfileWorkload` is synthetic; profile real workloads before making optimization decisions.

## Escape analysis

Ask the compiler to report escape decisions for this package:

```shell
go build -gcflags=-m=2 ./lesson/performance
```

Look for `LocalRecordValue` and `NewEscapeRecord`. Compiler diagnostics can change with Go versions and call-site context; an escape is not automatically a problem, and avoiding one is not automatically faster.
