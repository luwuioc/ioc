package lifecycle

import (
	"github.com/luwuioc/ioc/context"
)

const (
	DefaultPhase = (-1) ^ (-1 << 63)
)

var (
	_ Lifecycle      = (SmartLifecycle)(nil)
	_ context.Phased = (SmartLifecycle)(nil)
)

type SmartLifecycle interface {
	Lifecycle
	context.Phased
}
