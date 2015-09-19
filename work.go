package work

type Worker func(input interface{}) (output interface{})

func Start(w Worker, n int) (
        in chan interface{},
        out chan interface{},
        exit chan struct{},
        exited chan struct{}) {

        in, out = make(chan interface{}), make(chan interface{})
        exit, exited = make(chan struct{}), make(chan struct{})

        go func() {
                outstd := make(chan struct{}, n)
                for {
                        select {
                        case job := <-in:
                                outstd <- struct{}{}
                                go func(job interface{}) {
                                        out <- w(job)
                                        <-outstd
                                }(job)
                        case <-exit:
                                goto EXIT
                        }
                }
        EXIT:
                for i := 0; i < n; i++ {
                        outstd <- struct{}{}
                }
                close(outstd)
                close(exited)
        }()

        return
}
