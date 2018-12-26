- Create a cancellation channel on which no values are ever sent/ on which values never are sent
- Define a utility function "cancelled" that checks or polls the cancellation state

```go
var done = make(chan struct{})
func cancelled() bool {
    select {
        case <- done:
            return true
        default:
            return false
    }
}
```

- Close the done channel
- Making our goroutines respond to the cancellation adding in the main goroutine a third case to the select statement that tries to receive from the done channel. The function returns if this case is ever selected, but before it must first drain others goroutines.

```go
    for {
        select {
            case <- done:
                for range othergoroutine {
                    // Do nothing
                }

                return
        }
    }
```
