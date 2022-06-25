package interceptor

import (
	"github.com/photowey/fastrouter/api/request"
	"github.com/photowey/fastrouter/ordered"
)

type Interceptor interface {
	ordered.PriorityOrdered
	Name() string
	PreHandle(rctx *request.Context) bool
	PostHandle(rctx *request.Context)
}
