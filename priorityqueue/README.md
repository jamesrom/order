# priorityqueue
A fast, generic, thread safe priority queue for Go 1.18+.

## Features

 - Minimal API
 - Generic initializers – anything that can be ordered can be put into the queue
 - Built-in thread safety
 - Speed – see benchmarks

## TODO

 - [ ] Documentation
 - [x] Benchmarks
 - [ ] Concurrency tests

## Under the hood
The priority queue implements a binary heap (using `container/heap` from the standard library).

### Benchmarks
A priority queue is populated to a given size, we then measure the performance of Push/Pop/Peek.
<p align="center">

![image](https://user-images.githubusercontent.com/539129/154957192-eefdfd1b-bcd8-4af5-9156-43b16cb614b7.png)

</p>

Two benchmarks for Push are measured: PushMin pushes the lowest priority item into the queue, which represents the _best-case_ swim-up scenario. Conversely, PushMax pushes the highest priority item into the queue, which represents the _worst-case_ swim-up scenario.

```sh
goos: windows
goarch: amd64
pkg: github.com/jamesrom/priorityqueue
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkPushMaxOntoSize10-8         	 8987006	       142.5 ns/op	      51 B/op	       0 allocs/op
BenchmarkPushMaxOntoSize100-8        	 9489698	       135.7 ns/op	      49 B/op	       0 allocs/op
BenchmarkPushMaxOntoSize1000-8       	 9486435	       137.9 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize10000-8      	 9446527	       136.0 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize100000-8     	 9374347	       136.4 ns/op	      49 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize1000000-8    	 9056944	       142.9 ns/op	      57 B/op	       1 allocs/op
BenchmarkPushMaxOntoSize10000000-8   	 8540882	       150.5 ns/op	      40 B/op	       1 allocs/op
BenchmarkPushMinOntoSize10-8         	25805064	        45.41 ns/op	      54 B/op	       1 allocs/op
BenchmarkPushMinOntoSize100-8        	26073862	        45.74 ns/op	      54 B/op	       1 allocs/op
BenchmarkPushMinOntoSize1000-8       	26361690	        44.61 ns/op	      53 B/op	       1 allocs/op
BenchmarkPushMinOntoSize10000-8      	26388182	        41.99 ns/op	      53 B/op	       1 allocs/op
BenchmarkPushMinOntoSize100000-8     	26091663	        42.56 ns/op	      53 B/op	       1 allocs/op
BenchmarkPushMinOntoSize1000000-8    	27603910	        41.52 ns/op	      50 B/op	       1 allocs/op
BenchmarkPushMinOntoSize10000000-8   	32006144	        46.13 ns/op	      51 B/op	       1 allocs/op
BenchmarkPopFromSize10-8             	 6760711	       202.0 ns/op	       7 B/op	       0 allocs/op
BenchmarkPopFromSize100-8            	 6468810	       206.4 ns/op	       7 B/op	       0 allocs/op
BenchmarkPopFromSize1000-8           	 6283765	       204.2 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize10000-8          	 6299400	       218.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize100000-8         	 6451550	       215.1 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize1000000-8        	 5741742	       216.7 ns/op	       8 B/op	       1 allocs/op
BenchmarkPopFromSize10000000-8       	 5393007	       223.3 ns/op	       8 B/op	       1 allocs/op
BenchmarkPeekFromSize10-8            	92274331	        12.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize100-8           	92274331	        12.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize1000-8          	92299881	        13.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize10000-8         	92272202	        12.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize100000-8        	92318343	        12.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize1000000-8       	92499093	        12.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFromSize10000000-8      	96001536	        12.92 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 69.0% of statements
ok  	github.com/jamesrom/priorityqueue	75.666s
```
