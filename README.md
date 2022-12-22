# string-concat-performance

Run: ``` go test -gcflags=-N -bench=.```


* Method 1: Naive Appending with +=
* Method 2: Sprintf
* Method 3: strings.Join
* Method 4: strings.Builder
* Method 5: bytes.Buffer

Results:

```text
goos: linux
goarch: amd64
pkg: github.com/akhilesharora/string-concat-performance
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkNaivePlus-8                      308402            213744 ns/op
BenchmarkLongNaivePlus-8                   10000            633633 ns/op
BenchmarkConstNaivePlus-8                 183928            109711 ns/op
BenchmarkJoin-8                           246140            181815 ns/op
BenchmarkLongJoin-8                        10000            717604 ns/op
BenchmarkConstJoin-8                      207338            140726 ns/op
BenchmarkSprintf-8                        166102            244065 ns/op
BenchmarkLongSprintf-8                     10000           1119389 ns/op
BenchmarkConstSprintf-8                   192945            303452 ns/op
BenchmarkStringBuilder-8                71856100             17.31 ns/op
BenchmarkLongStringBuilder-8             2814878             463.9 ns/op
BenchmarkConstStringBuilder-8           69755790             18.62 ns/op
BenchmarkBytesBuffer-8                  125612434            10.35 ns/op
BenchmarkLongBytesBuffer-8               5283704             846.7 ns/op
BenchmarkConstBytesBuffer-8             114313764            10.12 ns/op
PASS
ok      github.com/akhilesharora/string-concat-performance      299.848s

```