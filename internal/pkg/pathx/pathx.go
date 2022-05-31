package pathx

import (
	"fmt"
)

func BuildMapping(method, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
