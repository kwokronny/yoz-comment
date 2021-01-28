package helper

import (
	"github.com/importcjj/sensitive"
)

var Filter sensitive.Filter

func init() {
	filter := sensitive.New()
	filter.LoadWordDict("path/to/dict")
}
