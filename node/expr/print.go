package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Print struct {
	name string
	expr node.Node
}

func NewPrint(expression node.Node) node.Node {
	return Print{
		"Print",
		expression,
	}
}

func (n Print) Name() string {
	return "Print"
}

func (n Print) Attributes() map[string]interface{} {
	return nil
}

func (n Print) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
