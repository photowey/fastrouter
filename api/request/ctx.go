package request

import (
	"github.com/photowey/fastrouter/internal/pkg/pathx"
	"github.com/valyala/fasthttp"
)

type Context struct {
	Ctx    *fasthttp.RequestCtx
	Method string
	Path   string
}

func (ctx Context) BuildMapping() string {
	return pathx.BuildMapping(ctx.Method, ctx.Path)
}
