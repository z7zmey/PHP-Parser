package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Foreach struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
	key        node.Node
	variable   node.Node
	stmt       node.Node
}

func NewForeach(expr node.Node, key node.Node, variable node.Node, stmt node.Node, byRef bool) node.Node {
	return Foreach{
		"Foreach",
		map[string]interface{}{
			"byRef": byRef,
		},
		expr,
		key,
		variable,
		stmt,
	}
}

func (n Foreach) Name() string {
	return "Foreach"
}

func (n Foreach) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Foreach) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	if n.key != nil {
		vv := v.GetChildrenVisitor("key")
		n.key.Walk(vv)
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
