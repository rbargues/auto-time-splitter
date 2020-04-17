package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(33333333 * time.Nanosecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(10 * time.Second)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}