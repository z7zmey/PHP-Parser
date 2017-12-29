package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type PropertyList struct {
	name       string
	modifiers  []node.Node
	properties []node.Node
}

func NewPropertyList(modifiers []node.Node, properties []node.Node) node.Node {
	return PropertyList{
		"PropertyList",
		modifiers,
		properties,
	}
}

func (n PropertyList) Name() string {
	return "PropertyList"
}

func (n PropertyList) Attributes() map[string]interface{} {
	return nil
}

func (n PropertyList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.modifiers != nil {
		vv := v.GetChildrenVisitor("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.properties != nil {
		vv := v.GetChildrenVisitor("properties")
		for _, nn := range n.properties {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
