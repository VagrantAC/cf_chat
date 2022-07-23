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
    c.SetOkStatus().Json("ok")
    
    finish <- struct{}{}
  }()

  select {
  case p := <-panicChan:
    c.WriterMux().Lock()
    defer c.WriterMux().Unlock()
    log.Println(p)
    c.SetStatus(500).Json("panic")
  	break
  case <- finish:
  	fmt.Println("finish")
  	break
  case <- durationCtx.Done():
  	c.WriterMux().Lock()
  	defer c.WriterMux().Unlock()
    c.SetStatus(500).Json("time out")
    c.SetHasTimeout()
  	break
  }
  return nil
}
