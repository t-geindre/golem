package helper

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Node struct {
	XMLName  xml.Name   `xml:""`
	Attrs    []xml.Attr `xml:",any,attr"`
	Children []*Node    `xml:",any"`
	Content  []byte     `xml:",chardata"`
}

func ParseXML(r io.Reader) (*Node, error) {
	var dom *Node
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(b, &dom)
	if err != nil {
		return nil, err
	}

	return dom, nil
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

func (n *Node) GetChild(name string) (*Node, error) {
	for _, child := range n.Children {
		if child.GetName() == name {
			return child, nil
		}
	}

	return nil, fmt.Errorf("child node not found: %s", name)
}

func (n *Node) GetChildWithAttrs(name string, attrs ...string) (*Node, error) {
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
				return child, nil
			}
		}
	}

	// Todo improve error message to display the expected attributes
	return nil, fmt.Errorf("child node not found: %s", name)
}

func (n *Node) GetContent() string {
	return strings.Trim(string(n.Content), " \n\t")
}
