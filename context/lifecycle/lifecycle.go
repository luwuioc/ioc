package lifecycle

type Lifecycle interface {
	Start()
	Stop()
	IsRunning() bool
}
