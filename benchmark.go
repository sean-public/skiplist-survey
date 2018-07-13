package main

import (
    "fmt"
    "reflect"
    "runtime"
    "sync"
    "time"
)

const (
    // The range of values for n passed to the individual benchmarks
    start, end, step int = 100000, 3000000, 100000
)

var (
    wg             sync.WaitGroup
    testByteString = []byte(fmt.Sprint("test value"))
    test10Kilobyte = make([]byte, 10240)
)

// timeTrack will print out the number of nanoseconds since the start time divided by n
// Useful for printing out how long each iteration took in a benchmark
func timeTrack(start time.Time, n int) {
    loopNS := time.Since(start).Nanoseconds() / int64(n)
    fmt.Print(loopNS)
}

// iterations is used to print out the CSV header with iteration counts
func iterations(n int) {
    fmt.Print(n)
}

// funcName returns just the function name of a string, given any function at all
func funcName(f func(int)) string {
    return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()[5:]
}

// runIterations executes the tests in a loop with the given parameters
func runIterations(name string, start, end, step int, f func(int)) {
    fmt.Print(name, ",")
    for i := start; i <= end; i += step {
        f(i)
        fmt.Print(",")
    }
    fmt.Println()
}

func main() {
    // first, print the CSV header with iteration counts
    runIterations("iterations", start, end, step, iterations)

    allFunctions := append(seanFunctions, zhenjlFunctions...)
    allFunctions = append(allFunctions, mtchavezFunctions...)
    allFunctions = append(allFunctions, huanduFunctions...)
    allFunctions = append(allFunctions, colFunctions...)
    allFunctions = append(allFunctions, ryszardFunctions...)
    allFunctions = append(allFunctions, mtFunctions...)

    for _, f := range allFunctions {
        runIterations(funcName(f), start, end, step, f)
    }
}
