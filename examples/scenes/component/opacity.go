package component

//go:generate golem component Opacity
type Opacity struct {
	Value float32
}

func NewOpacity(value float32) *Opacity {
	return &Opacity{
		Value: value,
	}
}
