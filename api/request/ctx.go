package request

import (
	"github.com/photowey/fastrouter/internal/pkg/pathx"
	"github.com/photowey/fastrouter/pkg/collection"
	"github.com/photowey/fastrouter/pkg/headermap"
	"github.com/photowey/fastrouter/pkg/jsonx"
	"github.com/photowey/fastrouter/pkg/nanoid"
	"github.com/valyala/fasthttp"
)

const (
	XForwardedFor = "X-Forwarded-For"
	ContentType   = "Content-Type"

	emptyString          = ""
	defaultHeaderContent = "-"
)

const (
	IdLength = 32
)

type Context struct {
	Ip             string
	RequestId      string
	ContentType    string
	Path           string
	Method         string
	UserAgent      string
	HeaderMap      headermap.HeaderMap
	BodyMap        collection.StringMap
	HandleResponse bool
	Header         headermap.HeaderMap
	StatusCode     int
	Body           []byte
}

func (ctx Context) BuildMapping() string {
	return pathx.BuildMapping(ctx.Method, ctx.Path)
}

func (ctx *Context) Response(body []byte) {
	ctx.HandleResponse = true
	ctx.Body = body
}

func (ctx *Context) ResponseJSON(body any) {
	ctx.HandleResponse = true
	ctx.Body = jsonx.Bytes(body)
}

func NewRequestContext(ctx *fasthttp.RequestCtx) *Context {
	return &Context{
		Ip:          parseIp(ctx),
		RequestId:   buildRequestId(),
		ContentType: parseContentType(ctx),
		Path:        parsePath(ctx),
		Method:      parseMethod(ctx),
		UserAgent:   parseUserAgent(ctx),
		HeaderMap:   headermap.NewHeaderMap(),
		BodyMap:     collection.NewStringMap(),
		StatusCode:  fasthttp.StatusOK,
	}
}

func parsePath(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Path())
}

func parseMethod(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Method())
}

func parseUserAgent(ctx *fasthttp.RequestCtx) string {
	return string(ctx.UserAgent())
}

func parseIp(ctx *fasthttp.RequestCtx) string {
	ip := string(ctx.Request.Header.Peek(XForwardedFor))
	if ip == emptyString {
		ip = ctx.RemoteIP().String()
		if ip == emptyString {
			return defaultHeaderContent
		}
	}

	return ip
}

func buildRequestId() string {
	return nanoid.MustNew(IdLength)
}

func parseContentType(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Request.Header.Peek(ContentType))
}
