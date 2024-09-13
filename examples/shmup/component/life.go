package component

//go:generate golem component Life
type Life struct {
	Max, Current int
}

func NewLife(max int) *Life {
	return &Life{
		Max:     max,
		Current: max,
	}
}
