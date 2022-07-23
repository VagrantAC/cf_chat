package route

import (
	"cf_chat/framework"
	"context"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
  finish := make(chan struct{}, 1)
  panicChan := make(chan interface{}, 1)

  durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(time.Second))
  defer cancel()

  go func() {
    defer func() {
      if p := recover(); p != nil {
        panicChan <- p
      }
    }()

    time.Sleep(time.Second)
    c.Json(200, "ok")
    
    finish <- struct{}{}
  }()

  select {
  case p := <-panicChan:
    c.WriterMux().Lock()
    defer c.WriterMux().Unlock()
    log.Println(p)
    c.Json(500, "panic")
  	break
  case <- finish:
  	fmt.Println("finish")
  	break
  case <- durationCtx.Done():
  	c.WriterMux().Lock()
  	defer c.WriterMux().Unlock()
    c.Json(500, "time out")
    c.SetHasTimeout()
  	break
  }
  return nil
}
