package component

//go:generate golem component Lifecycle
type Lifecycle struct {
	SetUp    func()
	TearDown func()
}

func NewLifecycle() *Lifecycle {
	return &Lifecycle{}
}
