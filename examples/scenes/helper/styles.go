package helper

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"image/color"
	"reflect"
	"strconv"
)

const (
	DefaultFont         = "embd://font.otf"
	DefaultSize float64 = 12
)

type Style struct {
	PosRelX      *float64    `xml:"pos-x"`
	PosRelY      *float64    `xml:"pos-y"`
	PosOriginX   *float64    `xml:"pos-orig-x"`
	PosOriginY   *float64    `xml:"pos-orig-y"`
	Scale        *float64    `xml:"scale"`
	ScaleOriginX *float64    `xml:"scale-orig-x"`
	ScaleOriginY *float64    `xml:"scale-orig-y"`
	Opacity      *float64    `xml:"opacity"`
	Color        color.Color `xml:"color"`
	FontSource   *string     `xml:"font"`
	FontSize     *float64    `xml:"font-size"`
}

type StyleLoader struct {
	styles map[string]*Style
	fonts  map[string]*sfnt.Font
	faces  map[string]text.Face
	scale  float64
}

func NewStyleLoader() *StyleLoader {
	return &StyleLoader{
		styles: make(map[string]*Style),
		fonts:  make(map[string]*sfnt.Font),
		faces:  make(map[string]text.Face),
		scale:  1,
	}
}

func (sl *StyleLoader) SetScale(scale float64) {
	sl.scale = scale
}

func (sl *StyleLoader) LoadXML(node *Node) error {
	for _, child := range node.Children {
		name := child.GetAttr("name")
		if name == "" {
			return fmt.Errorf("missing name attribute in style")
		}

		st, err := sl.GetXMLStyle(child)
		if err != nil {
			return err
		}

		sl.styles[name] = st
	}

	return nil
}

func (sl *StyleLoader) GetXMLStyle(node *Node) (*Style, error) {
	s := Style{}
	st := reflect.TypeOf(s)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		tag := field.Tag.Get("xml")
		if tag == "" {
			continue
		}

		v := node.GetAttr(tag)
		if len(v) == 0 {
			continue
		}

		if field.Type.Kind() == reflect.Pointer {
			switch field.Type.Elem().Kind() {
			case reflect.Float64:
				f, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, err
				}
				reflect.ValueOf(&s).Elem().Field(i).Set(reflect.ValueOf(&f))
			case reflect.String:
				reflect.ValueOf(&s).Elem().Field(i).Set(reflect.ValueOf(&v))
			default:
				return nil, fmt.Errorf("unsupported field type: %s", field.Type.Kind())
			}

			continue
		}

		if field.Type.Kind() == reflect.Interface {
			// todo this is only valid for color and not safe if we add more interfaces
			col, err := sl.ParseColor(v)
			if err != nil {
				return nil, err
			}
			reflect.ValueOf(&s).Elem().Field(i).Set(reflect.ValueOf(col))
		}
	}

	return &s, nil
}

func (sl *StyleLoader) GetNamedStyle(name string) (*Style, error) {
	if len(name) == 0 {
		return nil, nil
	}

	st, ok := sl.styles[name]
	if !ok {
		return nil, fmt.Errorf("style not found: %s", name)
	}

	return st, nil
}

func (sl *StyleLoader) ApplyStyles(e golem.Entity, sts ...*Style) error {
	st := sl.MergeStyles(sts)

	pos := component.GetPosition(e)
	if pos != nil {
		if st.PosRelX != nil {
			pos.RelX = *st.PosRelX / 100
		}
		if st.PosRelY != nil {
			pos.RelY = *st.PosRelY / 100
		}
		if st.PosOriginX != nil {
			pos.OriginX = *st.PosOriginX / 100
		}
		if st.PosOriginY != nil {
			pos.OriginY = *st.PosOriginY / 100
		}
	}

	txt := component.GetText(e)
	if txt != nil {
		face, err := sl.getFontFace(*st.FontSource, *st.FontSize)
		if err != nil {
			return err
		}
		txt.Face = face
	}

	scl := component.GetScale(e)
	if scl != nil {
		scale := 1.0
		if txt == nil {
			// on text components, the scale is applied to the font size
			scale = sl.scale
		}
		if st.Scale != nil {
			scale *= *st.Scale
		}
		scl.Value = scale
		if st.ScaleOriginX != nil {
			scl.OriginX = *st.ScaleOriginX
		}
		if st.ScaleOriginY != nil {
			scl.OriginY = *st.ScaleOriginY
		}
	}
	op := component.GetOpacity(e)
	if op != nil && st.Opacity != nil {
		op.Value = float32(*st.Opacity)
	}

	col := component.GetColor(e)
	if col != nil && st.Color != nil {
		col.Value = st.Color
	}

	return nil
}

func (sl *StyleLoader) ParseColor(v string) (color.Color, error) {
	// todo there must be a way to factorize this
	if v[0] == '#' {
		if len(v) == 7 {
			r, err := strconv.ParseUint(v[1:3], 16, 8)
			if err != nil {
				return nil, err
			}
			g, err := strconv.ParseUint(v[3:5], 16, 8)
			if err != nil {
				return nil, err
			}
			b, err := strconv.ParseUint(v[5:7], 16, 8)
			if err != nil {
				return nil, err
			}
			return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff}, nil
		}
		if len(v) == 4 {
			r, err := strconv.ParseUint(v[1:2], 16, 8)
			if err != nil {
				return nil, err
			}
			g, err := strconv.ParseUint(v[2:3], 16, 8)
			if err != nil {
				return nil, err
			}
			b, err := strconv.ParseUint(v[3:4], 16, 8)
			if err != nil {
				return nil, err
			}
			return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff}, nil
		}
	}
	if c, ok := colornames.Map[v]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("invalid color value: %s", v)
}

func (sl *StyleLoader) getFontFace(path string, size float64) (text.Face, error) {
	fKey := fmt.Sprintf("%s-%f", path, size)
	if f, ok := sl.faces[fKey]; ok {
		return f, nil
	}

	ft, ok := sl.fonts[path]
	if !ok {
		bts, err := ReadFile(path)
		if err != nil {
			return nil, err
		}

		ft, err = sfnt.Parse(bts)
		if err != nil {
			return nil, err
		}

		sl.fonts[path] = ft
	}

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size:    size * sl.scale,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	sl.faces[fKey] = text.NewGoXFace(face)

	return sl.faces[fKey], nil
}

func (sl *StyleLoader) MergeStyles(styles []*Style) *Style {
	if len(styles) == 0 {
		return nil
	}

	if len(styles) == 1 {
		return styles[0]
	}

	s := Style{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if field.Type.Kind() != reflect.Ptr && field.Type.Kind() != reflect.Interface {
			continue
		}

		for _, style := range styles {
			if style == nil {
				continue
			}
			if reflect.ValueOf(style).Elem().Field(i).IsNil() {
				continue
			}

			reflect.ValueOf(&s).Elem().Field(i).Set(reflect.ValueOf(reflect.ValueOf(style).Elem().Field(i).Interface()))
		}
	}

	if s.FontSource == nil {
		s.FontSource = new(string)
		*s.FontSource = DefaultFont
	}

	if s.FontSize == nil {
		s.FontSize = new(float64)
		*s.FontSize = DefaultSize
	}

	return &s
}
