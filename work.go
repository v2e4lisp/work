package work

// Worker function that hanndles jobs
type Worker func(input interface{}) (output interface{})

// Start working.
//
// w is the Worker function handling jobs.
// n is the number of workers running at the same time.
//
// Jobs should be send to inbound channel. When all the jobs are sent,
// close the inbound channel to tell workers that there is no job left.
// Worker function's return value will be sent this channel. When all
// the jobs are handled, it will get closed.
func Start(w Worker, n int) (inbound chan<- interface{}, outbound <-chan interface{}) {
        in, out := make(chan interface{}), make(chan interface{})

        go func() {
                outstd := make(chan struct{}, n)
                for job := range in {
                        outstd <- struct{}{}
                        go func(job interface{}) {
                                out <- w(job)
                                <-outstd
                        }(job)
                }

                for i := 0; i < n; i++ {
                        outstd <- struct{}{}
                }
                close(outstd)
                close(out)
        }()

        return in, out
}
