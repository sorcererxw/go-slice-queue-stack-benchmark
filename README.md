# Go Slice Queue & Stack Benchmark

Compare the performance of queue and stack based on slice.

```shell
goos: darwin
goarch: amd64
pkg: github.com/sorcererxw/go-slice-queue-stack-benchmark
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
Benchmark
Benchmark/queue
Benchmark/queue-4         	   70870	     17735 ns/op
Benchmark/stack
Benchmark/stack-4         	  148848	      7655 ns/op
```