Suffix Array Induced Sorting(SA-IS) written in golang.
SA-IS: suffix array construction algorithm.

Two Efficient Algorithms for Linear Suffix Array Construction. Ge Nong, Sen Zhang and Wai Hong Chan [pdf](https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/ge-nong/Two%20Efficient%20Algorithms%20for%20Linear%20Time%20Suffix%20Array%20Construction.pdf)

## Installation

```
$ go get github.com/tobi-c/go-sais/suffixarray
```


## Benchmark

* CPU: Intel Core i5 1.8 GHz 
* Memory: 8G
* Go version: 1.6

```
$ go test -bench . github.com/tobi-c/go-sais/suffixarray
BenchmarkNewIndexRandom-4              2         778500400 ns/op
BenchmarkNewIndexRepeat-4             10         162400650 ns/op
BenchmarkSaveRestore-4                30          45003666 ns/op          92.88 MB/s
```

```
$ go test -bench . index/suffixarray
BenchmarkNewIndexRandom-4              1        1448997500 ns/op
BenchmarkNewIndexRepeat-4              1        2484993800 ns/op
BenchmarkSaveRestore-4                50          45519330 ns/op          91.82 MB/s
```


## License
BSD License

(suffixarray.go and suffixarray_test.go copied from https://github.com/golang/go/tree/master/src/index/suffixarray)

