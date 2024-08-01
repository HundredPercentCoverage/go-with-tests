# go-with-tests

Working through [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests).

## Notes

### Commands

`go test`
`go test -v` for verbose
`go test -cover` to generate coverage
`go test -bench=.` for benchmark testing

### TIL Items

- Go function return types can be named
- Add comments to functions to document them
- Can also document with testable examples (prefix test with `Example`)
- `godoc` must be installed separately so it can be updated more often
  - `go install -v golang.org/x/tools/cmd/godoc@latest` (you might need to get your GOPATH variable set correctly)
- Arrays & Slices
  - Arrays have fixed **capacity** e.g. `[5]int{1, 2, 3, 4, 5}`. A function that takes a fixed-capacity array has to take the same capacity in its arguments e.g. `func Foo(arr [5]int) int {}`
  - Capacity is not the same as size
  - Slices do not have fixed size e.g. `nums := []int{}`, so the function taking a slice as an argument becomes `func Foo(arr []int) int {}`
  - Slices of specific capacity can be initialised with the `make()` function e.g. `sums := make([]int, lengthOfNumbers)` (`make` has more arguments)
  - Append to slices with the `append()` function
- A _variadic function_ takes a variable number of arguments of a given type, e.g. `func SomeFunc(nums ...int) {}` will take an indefinite number of int args (but is also equivalent to `func SomeFunc(nums []int) {}`)

#### Test Types

##### Unit tests

- `t *testing.T` for unit tests
- Run with `go test -v` (`-v` for verbose is optional)
- Group multiple tests in one function with `t.Run()` e.g.

```go
func TestSomeFunc(t *testing.T) {
  t.Run("first thing", func(t *testing.T) {
    // One test against SomeFunc
  })

  t.Run("other thing", func(t *testing.T) {
    // Another test against SomeFunc
  })
}
```

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
