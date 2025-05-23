package presenters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	C *gin.Context
}

func (ctx *Context) BindJSON(obj interface{}) error {
	return ctx.C.ShouldBindJSON(obj)
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.C.JSON(code, obj)
}

func (ctx *Context) Param(key string) string {
	return ctx.C.Param(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.C.Query(key)
}

func (ctx *Context) Status(code int) {
	ctx.C.Status(code)
}

func (ctx *Context) Request() *http.Request {
	return ctx.C.Request
}

func (ctx *Context) Value(key any) (string, bool) {
	val, ok := ctx.C.Value(key).(string)
	return val, ok
}

func (ctx *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	ctx.C.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}
