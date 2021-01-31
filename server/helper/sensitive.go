package helper

import (
	"github.com/importcjj/sensitive"
)

// SensitiveValid 敏感字验证
func SensitiveValid(content string) (bool, string) {
	filter := sensitive.New()
	filter.LoadWordDict(Config.SensitivePath)
	return filter.Validate(content)
}
