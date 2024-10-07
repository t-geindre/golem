package component

//go:generate golem component Lifecycle
type Lifecycle struct {
	SetUp    func()
	TearDown func()
}

func NewLifecycle(setup, teardown func()) *Lifecycle {
	return &Lifecycle{
		SetUp:    setup,
		TearDown: teardown,
	}
}
