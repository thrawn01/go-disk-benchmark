### To setup the Cassandra tables for test
```shell
$ go test ./benchmark
```

### To Run the benchmarks
```shell
$ go test -bench=.
```

### Results
**Block** is obviously the fastest since it doesn't need to open any new files. I
included it in the test as a control for the other tests.

**File** and **Pog** are the second fastest. I believe this is because they don't
sync to the underlying disk like **Pebble** and **Bolt** do.

**Cassandra** has good performance as it also doesn't really need to sync to disk.
it's write reliability comes from sending the data to replicas.

NOTE: I commented out the 25MB test for Cassandra as my local database isn't tuned
for 25MB and it errors for some reason. I know 25MB blocks are possible as we do
this in production, so I just disabled the test here.

```
goos: darwin
goarch: arm64
pkg: github.com/mailgun/geo-quest/benchmark
BenchmarkBlock_100KB
BenchmarkBlock_100KB-10               65038         20760 ns/op
BenchmarkPog_100KB
BenchmarkPog_100KB-10                 18448         75664 ns/op
BenchmarkFile_100KB
BenchmarkFile_100KB-10                12748         93184 ns/op
BenchmarkCassandra_100KB
BenchmarkCassandra_100KB-10             951       1607128 ns/op
BenchmarkPebble_100KB
BenchmarkPebble_100KB-10                199       5815951 ns/op
BenchmarkBolt_100KB
BenchmarkBolt_100KB-10                  124      10283756 ns/op
BenchmarkBlock_10MB
BenchmarkBlock_10MB-10                  601       2364541 ns/op
BenchmarkFile_10MB
BenchmarkFile_10MB-10                   520       2890556 ns/op
BenchmarkPog_10MB
BenchmarkPog_10MB-10                    313       4339664 ns/op
BenchmarkCassandra_10MB
BenchmarkCassandra_10MB-10               31      43400421 ns/op
BenchmarkPebble_10MB
BenchmarkPebble_10MB-10                  32      39782395 ns/op
BenchmarkBolt_10MB
BenchmarkBolt_10MB-10                    31      40666505 ns/op
BenchmarkBlock_25MB
BenchmarkBlock_25MB-10                  157       8347790 ns/op
BenchmarkFile_25MB
BenchmarkFile_25MB-10                   218       5226370 ns/op
BenchmarkPog_25MB
BenchmarkPog_25MB-10                    138       9541951 ns/op
BenchmarkPebble_25MB
BenchmarkPebble_25MB-10                  30      45803864 ns/op
BenchmarkBolt_25MB
BenchmarkBolt_25MB-10                    15     101294828 ns/op
PASS
```
