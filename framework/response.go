package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	Json(obj interface{}) IResponse

	Jsonp(obj interface{}) IResponse

	Xml(obj interface{}) IResponse

	Html(file string, obj interface{}) IResponse

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

	if _, err := ctx.responseWriter.Write([]byte(callback)); err != nil {
		return ctx
	}

	if _, err := ctx.responseWriter.Write([]byte("(")); err != nil {
		return ctx
	}

	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}

	if _, err := ctx.responseWriter.Write([]byte(")")); err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Context-Type", "application/html")
	ctx.responseWriter.Write(byt)
	return ctx
}

func (ctx *Context) Html(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}

	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Context-Type", "application/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Context-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string,
	secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Context-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}
