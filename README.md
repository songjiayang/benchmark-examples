# benchmark-examples

Benchmarks in Go.

## Unmarshal vs Decoder

[json_test.go](https://github.com/songjiayang/benchmark-examples/blob/master/json_test.go) benchmark result is:

```
go test -bench=^BenchmarkJson.*

goos: darwin
goarch: amd64
pkg: benchmark
BenchmarkJsonUnmarshal-8   	  213195	      4962 ns/op
BenchmarkJsonDecoder-8     	  261657	      4592 ns/op
PASS
ok  	benchmark	2.375s
```

As result using `Decoder` is speeder almost 8% and with lower memory, details to see [Unmarshal vs Decoder](https://stackoverflow.com/questions/21197239/decoding-json-using-json-unmarshal-vs-json-newdecoder-decode). 

