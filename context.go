package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Create a context with a timeout of 2 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    // Call a function that takes a context
    result := longRunningOperation(ctx)
    fmt.Println(result)
}

func longRunningOperation(ctx context.Context) string {
    select {
    case <-time.After(3 * time.Second):
        return "Operation completed"
    case <-ctx.Done():
        return "Operation canceled: " + ctx.Err().Error()
    }
}