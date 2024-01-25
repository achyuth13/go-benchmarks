#Benchadelic Markesh benchmarked passing by value and reference.

To run,
cd passing-by

#With Garbage Collector:
-----
go test -bench=. -count 1 -> With Garbage Collector

goos: darwin
goarch: arm64
pkg: passing-by
BenchmarkPBR-8          1000000000               0.3158 ns/op
BenchmarkPBV-8          1000000000               0.3205 ns/op
PASS
ok      passing-by      0.708s


#Without Garbage Collector
go test -bench=. -count 1 -gcflags=-N
----
goos: darwin
goarch: arm64
pkg: passing-by
BenchmarkPBR-8          573872079                2.085 ns/op
BenchmarkPBV-8            921590              1392 ns/op
PASS
ok      passing-by      2.710s

Go's garbage collector optimizes very well!
