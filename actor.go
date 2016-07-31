// Author: RuiTao XU <ruitao.xu@alibaba-inc.com>
// Description:
//	 Actor impl with the following attrbutes:
//		-- Async startup
//		-- Sync shutdown

package main

import (
    "fmt"
    "time"
)

type Actor struct {
    mQuit   chan bool
    mTicker *time.Ticker
}

func NewActor(intervalInMs int) *Actor {
    return &Actor{
        mTicker: time.NewTicker(time.Duration(intervalInMs) * time.Millisecond),
    }
}

// Startup may fail, so take error as return value
func (this *Actor) Startup() error {
    go this.loop()
    this.mQuit = make(chan bool)
    return nil
}

// do NOTHING when Startup failed
func (this *Actor) Shutdown() {
    if this.mQuit != nil {
        this.mQuit <- true
        this.mQuit = nil
    }
}

func (this *Actor) loop() {
    for {
        select {
        case <-this.mQuit:
            return
        case <-this.mTicker.C:
            this.workflow()
        }
    }
}

func (this *Actor) workflow() {
    fmt.Println(time.Now(), "workflow begin")
    time.Sleep(time.Duration(4) * time.Second)
    fmt.Println(time.Now(), "workflow end")
}

func main() {
    actor := NewActor(7000)
    fmt.Println(time.Now(), "startup")
    actor.Startup()
    time.Sleep(time.Duration(30) * time.Second)
    fmt.Println(time.Now(), "to shutdown")
    actor.Shutdown()
    fmt.Println(time.Now(), "shutdown")
}
