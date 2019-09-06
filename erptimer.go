package main

import (
    "time"
)

type ErpTimer struct {
    outChan chan int64
}

func erpTimer(outChan chan int64) *ErpTimer {
  return &ErpTimer{
      outChan: outChan,
  }
}

func (tmr *ErpTimer) run() {
    ticker := time.NewTicker(time.Second)
    for {
        t := <- ticker.C
        tmr.outChan <- t.Unix()
    }
}