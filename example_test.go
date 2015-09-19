package work_test

import (
        "fmt"
        "math/rand"
        "sort"
        "time"

        "github.com/v2e4lisp/work"
)

func Example() {
        // Worker function
        addOne := func(i interface{}) interface{} {
                j := i.(int) + 1

                rand.Seed(time.Now().UnixNano())
                t := rand.Intn(300)
                time.Sleep(time.Duration(t) * time.Millisecond)

                return j
        }

        // handle two jobs at the same time
        in, out := work.Start(addOne, 2)
        var results []int

        go func() {
                // send jobs to input channel
                for i := 1; i < 10; i++ {
                        in <- i
                }
                // We close the inbound channel to inform the workers that
                // there is no job left.
                close(in)
        }()

        // collecting the results
        for ret := range out {
                results = append(results, ret.(int))
        }

        sort.Ints(results)
        fmt.Println(results)
        // Output:
        // [2 3 4 5 6 7 8 9 10]
}
