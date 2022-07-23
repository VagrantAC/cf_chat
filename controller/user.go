package controller

import (
	"cf_chat/framework"
	"fmt"
	"time"
)

func UserLoginController(c *framework.Context) error {
  fmt.Println("+++++++++++++++")
  time.Sleep(time.Second * 5)
  fmt.Println("+++++++++++++++")
  c.SetOkStatus().Json("ok, UserLoginController")
  return nil
}
