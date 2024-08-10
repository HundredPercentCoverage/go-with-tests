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
- A method is like a function in a class
  - Method receivers are typically pointers, but watch out for pitfalls when calling the methods (sometimes you should declare things from structs with `&`)
- Struct definition `type Something struct { SomeField int }`
  - Struct declarations can use anonymous or named fields, e.g.
  ```go
  foo := Something{10}
  bar := Something{SomeField: 10}
  ```
- In formatted strings, using `v` as an example, `%v` will show the value in default format and `%#v` will show the Go syntax representation of the value ([link](https://pkg.go.dev/fmt#hdr-Printing))
- Interface definition `type Something interface { SomeMethod() int }`
- Use pointer receivers for methods, especially when you don't want the values to be copied, i.e. you want the original method instance to be changed
- Create your own errors, e.g. `var ErrSomeProblem = errors.New("some error")`
- Implement the `stringer` interface on things by making them implement the `String() string` method, then you can print them in format strings with the `%s` verb. This is useful when you create your own types, e.g.:

```go
type MyNumber int

func (n MyNumber) String() string {
  return fmt.Sprintf("%d is my number", n)
}
```

- Consider creating assertion helper functions outside the tests, so anyone reading the tests can just see the basic assertions first then drill into the helpers later if necessary.
- Helper functions can use the `t testing.TB` parameter to make them more flexible, i.e. can accept benchmark tests as well as unit tests

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

- Can also be run in a _table-driven_ way

```go
  // An array of structs, each one a test case, is our table
  areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 10.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	for _, areaTest := range areaTests {
    // Using t.Run() will list each test individually in the test output
    // so we can see which one fails if it does
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				// %#v will print out the struct with its fields and values
				// so if it fails we get detail
				t.Errorf("%#v got %g, wanted %g", areaTest.shape, got, areaTest.want)
			}
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
