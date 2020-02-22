package path

import (
	"fmt"
	"os"
)

func BasePath(r interface{}) string {
	p, _ := os.Getwd()

	if r != nil {
		return fmt.Sprint(p, r)
	}

	return p
}
