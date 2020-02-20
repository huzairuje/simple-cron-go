package main

import (
    "fmt"
    "time"
)

const INTERVAL_PERIOD time.Duration = 1 * time.Minute

// const just init value
// const HOUR_TO_TICK int = 18
// const MINUTE_TO_TICK int = 8
// const SECOND_TO_TICK int = 03

type jobTicker struct {
    t *time.Timer
}

func getNextTickDuration() time.Duration {
    now := time.Now()
    nextTick := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
    if nextTick.Before(now) {
        nextTick = nextTick.Add(INTERVAL_PERIOD)
    }
    return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
    fmt.Println("new tick here")
    return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
    fmt.Println("next tick here")
    jt.t.Reset(getNextTickDuration())
}

func main() {
    jt := NewJobTicker()
    for {
        <-jt.t.C
        fmt.Println(time.Now(), "- just ticked")
        jt.updateJobTicker()
    }
}