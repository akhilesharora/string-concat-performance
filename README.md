# string-concat-performance

### Usage:
```  
go test -run=^$ -bench=. -benchmem 
```

* Method 1: Naive Appending with +=
* Method 2: Sprintf
* Method 3: strings.Join
* Method 4: strings.Builder
* Method 5: bytes.Buffer

### Results:

```text
goos: linux
goarch: amd64
pkg: github.com/akhilesharora/string-concat-performance
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkNaivePlus-8                      330184            264491 ns/op          994605 B/op          1 allocs/op
BenchmarkLongNaivePlus-8                   10000            616223 ns/op         2229302 B/op          1 allocs/op
BenchmarkConstNaivePlus-8                 297709            214269 ns/op          897172 B/op          1 allocs/op
BenchmarkJoin-8                           233115            173120 ns/op          703376 B/op          1 allocs/op
BenchmarkLongJoin-8                        10000            694925 ns/op         2229308 B/op          1 allocs/op
BenchmarkConstJoin-8                      243464            206440 ns/op          734427 B/op          1 allocs/op
BenchmarkSprintf-8                        177138            248982 ns/op         1247406 B/op          5 allocs/op
BenchmarkLongSprintf-8                     10000           1042027 ns/op         4620556 B/op          6 allocs/op
BenchmarkConstSprintf-8                   164761            272869 ns/op         1160515 B/op          4 allocs/op
BenchmarkStringBuilder-8                109715338               10.19 ns/op           33 B/op          0 allocs/op
BenchmarkLongStringBuilder-8             2941603               376.2 ns/op          2437 B/op          0 allocs/op
BenchmarkConstStringBuilder-8           146921577                7.235 ns/op          31 B/op          0 allocs/op
BenchmarkBytesBuffer-8                  99538117                11.81 ns/op           23 B/op          0 allocs/op
BenchmarkLongBytesBuffer-8               5806888               375.4 ns/op          1285 B/op          0 allocs/op
BenchmarkConstBytesBuffer-8             100000000               10.29 ns/op           23 B/op          0 allocs/op
PASS
ok      github.com/akhilesharora/string-concat-performance      364.623s

```

### Machine:
``` bash 
$ go env
GO111MODULE="auto"
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.18.4"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build2064074376=/tmp/go-build -gno-record-gcc-switches"
```
