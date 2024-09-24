package helper

import (
	"fmt"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/pkg/golem"
	"strconv"
	"strings"
	"time"
)

type Transition struct {
	Func     component.TransitionFunc
	Duration time.Duration
}

type TransitionLoader struct {
	TransMap    map[string]component.TransitionFunc
	Transitions map[string]Transition
	Default     string
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

		tTypes := strings.Split(tNode.GetAttr("type"), ",")
		if len(tTypes) == 0 {
			return fmt.Errorf("missing attribute \"type\" in node \"transition\"")
		}

		transFs := make([]component.TransitionFunc, 0, len(tTypes))
		for _, tType := range tTypes {
			trans, ok := tl.TransMap[strings.TrimSpace(tType)]
			if !ok {
				return fmt.Errorf("unknown transition type: \"%s\"", tType)
			}
			transFs = append(transFs, trans)
		}

		var trans component.TransitionFunc
		if len(transFs) == 1 {
			trans = transFs[0]
		} else {
			trans = TransitionMulti(transFs)
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

		if tNode.GetAttr("default") == "true" {
			if tl.Default != "" {
				return fmt.Errorf("only one transition can be default")
			}
			tl.Default = name
		}
	}

	return nil
}

func (tl *TransitionLoader) ApplyTransition(scene *entity.Scene, name string) error {
	if len(name) == 0 {
		if tl.Default == "" {
			return fmt.Errorf("missing slide transition and no default transition defined")
		}
		name = tl.Default
	}

	if t, ok := tl.Transitions[name]; ok {
		scene.Transition = component.NewTransition(t.Func, t.Duration)
		return nil
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

func TransitionMulti(ts []component.TransitionFunc) component.TransitionFunc {
	return func(e golem.Entity, v float64) {
		for _, t := range ts {
			t(e, v)
		}
	}
}

func TransitionNone(_ golem.Entity, _ float64) {
	// Do nothing
}
