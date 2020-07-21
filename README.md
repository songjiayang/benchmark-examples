# benchmark-examples

Benchmarks in Go.

## Unmarshal vs Decoder

[json_test.go](https://github.com/songjiayang/benchmark-examples/blob/master/json_test.go) benchmark result is:

```bash
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

## len(str) vs str == ""

benchmark for empty string check, result is:

```bash
go test -bench=BenchmarkStringEmptyCheck

goos: darwin
goarch: amd64
pkg: benchmark
BenchmarkStringEmptyCheckWithDefault-8   	1000000000	         0.617 ns/op
BenchmarkStringEmptyCheckWithLen-8       	1000000000	         0.311 ns/op
PASS
ok  	benchmark	1.059s
```

As result using `len(str) == 0` is speeder than `str == ""` in go1.14.6.

But in go1.10.8 they are same, I have created a [performance issue](https://github.com/golang/go/issues/40330) about `str == ""`.