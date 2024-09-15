package component

//go:generate golem component Scene
type Scene struct {
	Name string
}

func NewScene(name string) *Scene {
	return &Scene{
		Name: name,
	}
}
