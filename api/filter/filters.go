package filter

import (
	"github.com/photowey/fastrouter/api/request"
	"github.com/photowey/fastrouter/ordered"
)

type Filter interface {
	ordered.PriorityOrdered
	Name() string
	Init()
	DoFilter(rctx *request.Context) error
	Destroy()
}
