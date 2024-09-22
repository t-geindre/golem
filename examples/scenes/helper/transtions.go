package helper

import (
	"fmt"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/pkg/golem"
	"strconv"
	"time"
)

type Transition struct {
	Func     component.TransitionFunc
	Duration time.Duration
}

type TransitionLoader struct {
	TransMap    map[string]component.TransitionFunc
	Transitions map[string]Transition
}

func NewTransitionLoader() *TransitionLoader {
	return &TransitionLoader{
		TransMap: map[string]component.TransitionFunc{
			"none":  TransitionNone,
			"fade":  TransitionFade,
			"scale": TransitionScale,
		},
		Transitions: make(map[string]Transition),
	}
}

func (tl *TransitionLoader) LoadXML(node *Node) error {
	for _, tNode := range node.Children {
		if tNode.GetName() != "transition" {
			return fmt.Errorf("invalid node node: \"%s\", \"transition\" expected", tNode.GetName())
		}

		name := tNode.GetAttr("name")
		if name == "" {
			return fmt.Errorf("missing attribute \"name\" in node \"transition\"")
		}

		if _, ok := tl.Transitions[name]; ok {
			return fmt.Errorf("duplicated transition name: \"%s\"", name)
		}

		tType := tNode.GetAttr("type")
		if tType == "" {
			return fmt.Errorf("missing attribute \"type\" in node \"transition\"")
		}
		trans, ok := tl.TransMap[tType]
		if !ok {
			return fmt.Errorf("unknown transition type: \"%s\"", tType)
		}

		dAttr := tNode.GetAttr("duration")
		duration, err := strconv.Atoi(dAttr)
		if err != nil {
			if len(dAttr) == 0 {
				duration = 0
			} else {
				return fmt.Errorf("invalid attribute \"duration\" in node \"transition\": %w", err)
			}
		}

		tl.Transitions[name] = Transition{
			Func:     trans,
			Duration: time.Duration(duration) * time.Millisecond,
		}
	}

	return nil
}

func (tl *TransitionLoader) ApplyTransition(scene *entity.Scene, name string) error {
	if t, ok := tl.Transitions[name]; ok {
		scene.Transition = component.NewTransition(t.Func, t.Duration)
		return nil
	}

	if len(name) == 0 {
		return fmt.Errorf("missing transition name")
	}

	return fmt.Errorf("unknown transition name: \"%s\"", name)
}

func TransitionFade(entity golem.Entity, v float64) {
	op := component.GetOpacity(entity)
	if op != nil {
		op.Value = float32(v)
	}
}

func TransitionScale(entity golem.Entity, v float64) {
	scale := component.GetScale(entity)
	if scale != nil {
		scale.Value = v
	}
}

func TransitionNone(_ golem.Entity, _ float64) {
	// Do nothing
}
