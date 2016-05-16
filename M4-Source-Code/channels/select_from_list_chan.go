//From http://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement
//tests :http://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement
// Process1 calls `fn` for each value received from any of the `chans`
// channels. The arguments to `fn` are the index of the channel the
// value came from and the string value. Process1 returns once all the
// channels are closed.
func Process1(chans []<-chan string, fn func(int, string)) {
    // Setup
    type item struct {
        int    // index of which channel this came from
        string // the actual string item
    }
    merged := make(chan item)
    var wg sync.WaitGroup
    wg.Add(len(chans))
    for i, c := range chans {
        go func(i int, c <-chan string) {
            // Reads and buffers a single item from `c` before
            // we even know if we can write to `merged`.
            //
            // Go doesn't provide a way to do something like:
            //     merged <- (<-c)
            // atomically, where we delay the read from `c`
            // until we can write to `merged`. The read from
            // `c` will always happen first (blocking as
            // required) and then we block on `merged` (with
            // either the above or the below syntax making
            // no difference).
            for s := range c {
                merged <- item{i, s}
            }
            // If/when this input channel is closed we just stop
            // writing to the merged channel and via the WaitGroup
            // let it be known there is one fewer channel active.
            wg.Done()
        }(i, c)
    }
    // One extra goroutine to watch for all the merging goroutines to
    // be finished and then close the merged channel.
    go func() {
        wg.Wait()
        close(merged)
    }()

    // "select-like" loop
    for i := range merged {
        // Process each value
        fn(i.int, i.string)
    }
}


// Reflection case 
// Process1 (which is still insignificant if their is a significant
// delay between incoming values or if `fn` runs for a significant
// time).
func Process2(chans []<-chan string, fn func(int, string)) {
    // Setup
    cases := make([]reflect.SelectCase, len(chans))
    // `ids` maps the index within cases to the original `chans` index.
    ids := make([]int, len(chans))
    for i, c := range chans {
        cases[i] = reflect.SelectCase{
            Dir:  reflect.SelectRecv,
            Chan: reflect.ValueOf(c),
        }
        ids[i] = i
    }

    // Select loop
    for len(cases) > 0 {
        // A difference here from the merging goroutines is
        // that `v` is the only value "in-flight" that any of
        // the workers have sent. All other workers are blocked
        // trying to send the single value they have calculated
        // where-as the goroutine version reads/buffers a single
        // extra value from each worker.
        i, v, ok := reflect.Select(cases)
        if !ok {
            // Channel cases[i] has been closed, remove it
            // from our slice of cases and update our ids
            // mapping as well.
            cases = append(cases[:i], cases[i+1:]...)
            ids = append(ids[:i], ids[i+1:]...)
            continue
        }

        // Process each value
        fn(ids[i], v.String())
    }
}
