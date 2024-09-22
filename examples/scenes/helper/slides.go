package helper

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"image/color"
	"math"
	"strconv"
)

type SlideLoader struct {
	layer              golem.LayerID
	layoutX, layoutY   float64
	displayX, displayY float64
	scale              float64
	slides             []golem.Entity
	transLoader        *TransitionLoader
	stylesLoader       *StyleLoader
	backgroundColor    color.Color
}

func NewSlideLoader(l golem.LayerID) *SlideLoader {
	dX, dY := ebiten.Monitor().Size()

	return &SlideLoader{
		layer:        l,
		displayX:     float64(dX),
		displayY:     float64(dY),
		transLoader:  NewTransitionLoader(),
		stylesLoader: NewStyleLoader(),
	}
}

func (sl *SlideLoader) GetBackgroundColor() color.Color {
	return sl.backgroundColor
}

func (sl *SlideLoader) LoadXML(node *Node) error {
	if node.GetName() != "slideshow" {
		return fmt.Errorf("invalid node node: \"%s\", \"slideshow\" expected", node.GetName())
	}

	err := sl.LoadXMLLayout(node)
	if err != nil {
		return err
	}
	sl.stylesLoader.SetScale(sl.scale)

	stylesNode, err := node.GetChild("styles")
	if err != nil {
		return err
	}

	err = sl.stylesLoader.LoadXML(stylesNode)
	if err != nil {
		return err
	}

	transNode, err := node.GetChild("transitions")
	if err != nil {
		return err
	}

	err = sl.transLoader.LoadXML(transNode)
	if err != nil {
		return err
	}

	sldNode, err := node.GetChild("slides")
	if err != nil {
		return err
	}

	return sl.LoadXMLSlides(sldNode)
}

func (sl *SlideLoader) LoadXMLLayout(node *Node) error {
	var err error

	sl.layoutX, err = strconv.ParseFloat(node.GetAttr("width"), 64)
	if err != nil {
		sl.layoutX = sl.displayX
	}

	sl.layoutY, err = strconv.ParseFloat(node.GetAttr("height"), 64)
	if err != nil {
		sl.layoutY = sl.displayY
	}

	sl.scale = math.Min(sl.displayX/sl.layoutX, sl.displayY/sl.layoutY)

	bgColor := node.GetAttr("background-color")
	if bgColor != "" {
		sl.backgroundColor, err = sl.stylesLoader.ParseColor(bgColor)
		if err != nil {
			return err
		}
	} else {
		sl.backgroundColor = color.Black
	}

	return nil
}

func (sl *SlideLoader) LoadXMLSlides(node *Node) error {
	sl.slides = make([]golem.Entity, 0)

	for _, sNode := range node.Children {
		if sNode.GetName() != "slide" {
			return fmt.Errorf("invalid node node: \"%s\", \"slide\" expected", sNode.GetName())
		}

		slide := entity.NewScene(sl.layer, sNode.GetAttr("name"))
		err := sl.transLoader.ApplyTransition(slide, sNode.GetAttr("transition"))
		if err != nil {
			return err
		}

		entities, err := sl.GetXmlSlideEntities(sNode)
		if err != nil {
			return err
		}

		slide.Lifecycle = component.NewLifecycle(
			func() {
				slide.World.AddLayers(sl.layer)
				slide.World.AddEntities(entities...)
				slide.World.AddSystems(
					system.NewAnimation(),
					system.NewSpriteRenderer(),
					system.NewTextRenderer(),
				)
			},
			func() {
				slide.World.Clear()
			},
		)

		sl.slides = append(sl.slides, slide)
	}
	return nil
}

func (sl *SlideLoader) GetXmlSlideEntities(node *Node) ([]golem.Entity, error) {
	entities := make([]golem.Entity, 0)

	for _, eNode := range node.Children {
		switch eNode.GetName() {
		case "text":
			en := entity.NewText(sl.layer, eNode.GetContent())
			err := sl.stylesLoader.ApplyStyle(en, eNode.GetAttr("style"))
			if err != nil {
				return nil, err
			}
			entities = append(entities, en)
		}
	}

	return entities, nil
}

func (sl *SlideLoader) GetSlides(l golem.LayerID) []golem.Entity {
	return sl.slides
}
