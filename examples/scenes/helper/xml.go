package helper

import (
	"encoding/xml"
	"io"
	"strings"
)

type Node struct {
	XMLName  xml.Name   `xml:""`
	Attrs    []xml.Attr `xml:",any,attr"`
	Children []*Node    `xml:",any"`
	Content  []byte     `xml:",chardata"`
}

func ParseXML(r io.Reader) *Node {
	var dom *Node
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(b, &dom)
	if err != nil {
		panic(err)
	}

	return dom
}

func NewEmptyNode() *Node {
	return &Node{
		Attrs:    make([]xml.Attr, 0),
		Children: make([]*Node, 0),
	}
}

func (n *Node) GetName() string {
	return strings.ToLower(n.XMLName.Local)
}

func (n *Node) GetAttr(name string) string {
	for _, attr := range n.Attrs {
		if strings.ToLower(attr.Name.Local) == name {
			return attr.Value
		}
	}

	return ""
}

func (n *Node) GetChild(name string) *Node {
	for _, child := range n.Children {
		if child.GetName() == name {
			return child
		}
	}

	return NewEmptyNode()
}

func (n *Node) GetChildWithAttrs(name string, attrs ...string) *Node {
	if len(attrs)%2 != 0 {
		panic("attrs must be a list of key-value pairs")
	}

	for _, child := range n.Children {
		if child.GetName() == name {
			found := true
			for i := 0; i < len(attrs); i += 2 {
				if child.GetAttr(attrs[i]) != attrs[i+1] {
					found = false
				}
			}
			if found {
				return child
			}
		}
	}

	return NewEmptyNode()
}

func (n *Node) GetContent() string {
	return strings.Trim(string(n.Content), " \n\t")
}
