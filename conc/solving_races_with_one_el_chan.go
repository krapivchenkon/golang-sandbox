// Blocked Goroutines and Resource Leaks


// Rob Pike talked about a number of fundamental concurrency patterns 
// in his "Go Concurrency Patterns" presentation at Google I/O in 2012. 
// Fetching the first result from a number of targets is one of them.

func First(query string, replicas ...Search) Result {  
    c := make(chan Result)
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
// The function starts a goroutines for each search replica. Each goroutine 
// sends its search result to the result channel. The first value from the
// result channel is returned.

// What about the results from the other goroutines? What about the goroutines 
// themselves?

// The result channel in the First() function is unbuffered. This means that 
// only the first goroutine returns. All other goroutines are stuck trying to 
// send their results. This means if you have more than one replica each call 
// will leak resources.

// To avoid the leaks you need to make sure all goroutines exit. One potential 
// solution is to use a buffered result channel big enough to hold all results.

func First(query string, replicas ...Search) Result {  
    c := make(chan Result,len(replicas))
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
// Another potential solution is to use a select statement with a default case 
// and a buffered result channel that can hold one value. The default case 
// ensures that the goroutines don't get stuck even when the result channel 
// can't receive messages.

func First(query string, replicas ...Search) Result {  
    c := make(chan Result,1)
    searchReplica := func(i int) { 
        select {
        case c <- replicas[i](query):
        default:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
// You can also use a special cancellation channel to interrupt the workers.

func First(query string, replicas ...Search) Result {  
    c := make(chan Result)
    done := make(chan struct{})
    defer close(done)
    searchReplica := func(i int) { 
        select {
        case c <- replicas[i](query):
        case <- done:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }

    return <-c
}
// Why did the presentation contain these bugs? Rob Pike simply didn't want to 
// comlicate the slides. It makes sense, but it can be a problem for new Go 
// developers who would use the code as is without thinking that it might have 
// problems.


// Using Pointer Receiver Methods On Value Instances

// level: advanced
// It's OK to call a pointer receiver method on a value as long as the value 
// is addressable. In other words, you don't need to have a value receiver 
// version of the method in some cases.