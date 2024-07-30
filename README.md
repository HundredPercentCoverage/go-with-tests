# go-with-tests

Working through [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests).

## Notes

### Commands

`go test`
`go test -v`
`go test -cover`

### TIL Items

- Go function return types can be named
- Add comments to functions to document them
- Can also document with testable examples (prefix test with `Example`)
- `godoc` must be installed separately so it can be updated more often
  - `go install -v golang.org/x/tools/cmd/godoc@latest` (you might need to get your GOPATH variable set correctly)

#### Test Types

##### Unit tests

- `t *testing.T` for unit tests
- Run with `go test -v` (`-v` for verbose is optional)

##### Benchmark Tests

- `b *testing.B` for benchmarks
- Run with `go test -bench=.`
- This will also run unit tests unless you skip them with e.g. `-run=XXX`, where `XXX` is a regex that doesn't match anything so no unit tests should run.
- Tests are run `b.N` times, where N is a number determined by the test runner, which is the number of times it takes until the runner is satisfied with the stability of the results.
- Example output for benchmarking a given function:

```
BenchmarkRepeat-8        9194344               130.8 ns/op  // 130.8 nanoseconds avg.
PASS                                                        // Result of unit tests
ok      gowithtests/iteration   1.497s
```
