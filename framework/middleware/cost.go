package middleware

import (
	"cf_chat/framework"
	"log"
	"time"
)

func Cost() framework.ControllerHandler {
  return func(c *framework.Context) error {
    start := time.Now()

    c.Next()

    end := time.Now()
    cost := end.Sub(start)
    log.Printf("api uri: %v, cost: %.10fs", c.GetRequest().RequestURI, cost.Seconds())
    return nil
  }
}
