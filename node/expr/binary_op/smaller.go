package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Smaller struct {
	BinaryOp
}

func NewSmaller(variable node.Node, expression node.Node) node.Node {
	return Smaller{
		BinaryOp{
			"BinarySmaller",
			variable,
			expression,
		},
	}
}

func (n Smaller) Name() string {
	return "Smaller"
}

func (n Smaller) Attributes() map[string]interface{} {
	return nil
}

func (n Smaller) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
