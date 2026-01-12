package debug

import (
	"github.com/goforj/godump"
	"sync/atomic"
)

var enabled atomic.Bool

func Set(v bool) {
	enabled.Store(v)
}

func Enabled() bool {
	return enabled.Load()
}
func Dump(v ...any) {
	if !enabled.Load() {
		return
	}
	godump.Dump(v...)
}
