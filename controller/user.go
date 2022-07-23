package controller

import "cf_chat/framework"

func UserLoginController(c *framework.Context) error {
  c.Json(200, "ok, UserLoginController")
  return nil
}
