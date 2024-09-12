package component

//go:generate golem life.go
type Life struct {
	Max, Current int
}

func NewLife(max int) *Life {
	return &Life{
		Max:     max,
		Current: max,
	}
}
