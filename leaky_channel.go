package main

import "fmt"

func main() {
    var queue chan int
    queue = make(chan int, 3)

    for i := 0; i < 5; i++ {
        select {
        case queue <- i:
            fmt.Println("Enqueue:", i)
        default:
            fmt.Println("Queue is full, drop:", i)
        }
    }
}
