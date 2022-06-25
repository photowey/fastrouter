package request

import (
	"github.com/photowey/fastrouter/ordered"
)

type Handler interface {
	ordered.PriorityOrdered
	Method() string
	Path() string
	Supports(rctx *Context) bool
	Handle(rctx *Context)
}
