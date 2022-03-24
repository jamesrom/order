# package `atomicbit`

Performance experiments storing and loading a single bit atomically.

## Benchmarks

### Build tag `atomicbit_int32`
```
> go test ./atomicbit/... -bench . -tags=atomicbit_int32
goos: windows
goarch: amd64
pkg: github.com/jamesrom/order/atomicbit
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNewTrue-8      1000000000               0.2180 ns/op
BenchmarkNewFalse-8     1000000000               0.2180 ns/op
BenchmarkFlip-8         305710051                3.922 ns/op
BenchmarkGet-8          1000000000               0.2735 ns/op
BenchmarkSetFalse-8     219972272                5.446 ns/op
BenchmarkSetTrue-8      220373896                5.457 ns/op
PASS
ok      github.com/jamesrom/order/atomicbit     6.061s
```

### Build tag `atomicbit_int64`
```
> go test ./atomicbit/... -bench . -tags=atomicbit_int64
goos: windows
goarch: amd64
pkg: github.com/jamesrom/order/atomicbit
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNewTrue-8      1000000000               0.2189 ns/op
BenchmarkNewFalse-8     1000000000               0.2200 ns/op
BenchmarkFlip-8         305710519                3.920 ns/op
BenchmarkGet-8          1000000000               0.2720 ns/op
BenchmarkSetFalse-8     220183122                5.448 ns/op
BenchmarkSetTrue-8      305709117                3.919 ns/op
PASS
ok      github.com/jamesrom/order/atomicbit     5.906s
```

### Build tag `atomicbit_uint32`
```
> go test ./atomicbit/... -bench .
goos: windows
goarch: amd64
pkg: github.com/jamesrom/order/atomicbit
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNewTrue-8      1000000000               0.2180 ns/op
BenchmarkNewFalse-8     1000000000               0.2180 ns/op
BenchmarkFlip-8         306102926                3.928 ns/op
BenchmarkGet-8          1000000000               0.2755 ns/op
BenchmarkSetFalse-8     220171083                5.443 ns/op
BenchmarkSetTrue-8      220575949                5.438 ns/op
PASS
ok      github.com/jamesrom/order/atomicbit     6.061s
```

### Build tag `atomicbit_uint64`
```
> go test ./atomicbit/... -bench . -tags=atomicbit_uint64
goos: windows
goarch: amd64
pkg: github.com/jamesrom/order/atomicbit
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNewTrue-8      1000000000               0.2180 ns/op
BenchmarkNewFalse-8     1000000000               0.2175 ns/op
BenchmarkFlip-8         306098709                3.919 ns/op
BenchmarkGet-8          1000000000               0.2725 ns/op
BenchmarkSetFalse-8     220373248                5.461 ns/op
BenchmarkSetTrue-8      304183357                3.938 ns/op
PASS
ok      github.com/jamesrom/order/atomicbit     5.907s
```

### Build tag `atomicbit_uintptr`
```
> go test ./atomicbit/... -bench . -tags=atomicbit_uintptr
goos: windows
goarch: amd64
pkg: github.com/jamesrom/order/atomicbit
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNewTrue-8      1000000000               0.2180 ns/op
BenchmarkNewFalse-8     1000000000               0.2187 ns/op
BenchmarkFlip-8         305318494                3.942 ns/op
BenchmarkGet-8          1000000000               0.2725 ns/op
BenchmarkSetFalse-8     220385634                5.495 ns/op
BenchmarkSetTrue-8      305733262                3.925 ns/op
PASS
ok      github.com/jamesrom/order/atomicbit     5.924s
```


## TODO

 - [ ] Documentation
 - [x] Benchmarks
   - [x] `int32`
   - [x] `int64`
   - [x] `uint32`
   - [x] `uint64`
   - [x] `uintptr`
 - [ ] Concurrency tests
