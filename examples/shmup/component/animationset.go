package component

//go:generate golem component AnimationSet
type AnimationSet struct {
	Animations map[string]*Animation
	Current    string
	Next       string
	Default    string
}

func NewAnimationSet(set map[string]*Animation, def string) *AnimationSet {
	return &AnimationSet{
		Animations: set,
		Next:       def,
		Default:    def,
	}
}
