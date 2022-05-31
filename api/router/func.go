package router

import (
	"github.com/photowey/fastrouter/internal/pkg/pathx"
)

func BuildMapping(method, path string) string {
	return pathx.BuildMapping(method, path)
}
