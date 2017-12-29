package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Global struct {
	name string
	vars []node.Node
}

func NewGlobal(vars []node.Node) node.Node {
	return Global{
		"Global",
		vars,
	}
}

func (n Global) Name() string {
	return "Global"
}

func (n Global) Attributes() map[string]interface{} {
	return nil
}

func (n Global) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.vars != nil {
		vv := v.GetChildrenVisitor("vars")
		for _, nn := range n.vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
