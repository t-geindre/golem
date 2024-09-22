package helper

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/pkg/golem"
	"image"
	"strconv"
	"strings"
	"time"
)

type EntityBuilder struct {
	images map[string]*ebiten.Image
}

func NewEntityBuilder() *EntityBuilder {
	return &EntityBuilder{
		images: make(map[string]*ebiten.Image),
	}
}

func (eb *EntityBuilder) GetImage(path string) (*ebiten.Image, error) {
	img, ok := eb.images[path]
	if !ok {
		file, err := OpenFile(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		img, _, err = ebitenutil.NewImageFromReader(file)
		if err != nil {
			return nil, err
		}

		eb.images[path] = img
	}

	return img, nil
}

func (eb *EntityBuilder) BuildFromXMLNode(l golem.LayerID, node *Node) (golem.Entity, error) {
	switch node.GetName() {
	case "text":
		return eb.buildText(l, node)
	case "img":
		return eb.buildImage(l, node)
	case "animation":
		return eb.buildAnimation(l, node)
	case "list":
		return eb.buildList(l, node)
	}
	return nil, fmt.Errorf("invalid node name: \"%s\"", node.GetName())
}

func (eb *EntityBuilder) buildList(l golem.LayerID, node *Node) (golem.Entity, error) {
	text := ""
	for _, child := range node.Children {
		text += fmt.Sprintf("• %s\n", child.GetContent())
	}
	return entity.NewText(l, strings.Trim(text, "\n")), nil
}

func (eb *EntityBuilder) buildText(l golem.LayerID, node *Node) (golem.Entity, error) {
	return entity.NewText(l, node.GetContent()), nil
}

func (eb *EntityBuilder) buildImage(l golem.LayerID, node *Node) (golem.Entity, error) {
	img, err := eb.GetImage(node.GetAttr("src"))
	if err != nil {
		return nil, err
	}
	return entity.NewImage(l, img), nil
}

func (eb *EntityBuilder) buildAnimation(l golem.LayerID, node *Node) (golem.Entity, error) {
	fs := make([]component.Frame, 0)

	img, err := eb.GetImage(node.GetAttr("src"))
	if err != nil {
		return nil, err
	}

	tw, err := strconv.Atoi(node.GetAttr("tw"))
	if err != nil {
		return nil, err
	}

	th, err := strconv.Atoi(node.GetAttr("th"))
	if err != nil {
		return nil, err
	}

	for _, fNode := range node.Children {
		d, err := strconv.Atoi(fNode.GetAttr("duration"))
		if err != nil {
			return nil, err
		}

		tx, err := strconv.Atoi(fNode.GetAttr("tx"))
		if err != nil {
			return nil, err
		}

		ty, err := strconv.Atoi(fNode.GetAttr("ty"))
		if err != nil {
			return nil, err
		}

		fImg := img.SubImage(image.Rect(tx*tw, ty*th, (tx+1)*tw, (ty+1)*th)).(*ebiten.Image)

		fs = append(fs, component.NewFrame(fImg, time.Millisecond*time.Duration(d)))
	}

	return entity.NewAnimation(l, fs...), nil
}
