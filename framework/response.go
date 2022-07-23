package framework

import (
  "html/template"
)

type IResponse interface {
  Json(obj interface{}) IResponse

  Jsonp(obj interface{}) IResponse

  Xml(obj interface{}) IResponse

  Html(obj interface{}) IResponse

  Text(format string, values ...interface{}) IResponse
  
  Redirect(path string) IResponse

  SetHeader(key string, val string) IResponse

  SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

  SetStatus(code int) IResponse

  SetOkStatus() IResponse
}

func (ctx *Context) Jsonp(obj interface{}) IResponse {
  callbackFunc, _ := ctx.QueryString("callback", "callback_function")
  ctx.SetHeader("Context-Type", "application/javascript")
  callback := template.JSEscapeString(callbackFunc)
  
  _, err := ctx.responseWriter.Write([]byte("("))
  if err != nil {
    return ctx
  }
}
