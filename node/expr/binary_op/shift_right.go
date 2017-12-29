package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftRight struct {
	BinaryOp
}

func NewShiftRight(variable node.Node, expression node.Node) node.Node {
	return ShiftRight{
		BinaryOp{
			"BinaryShiftRight",
			variable,
			expression,
		},
	}
}

func (n ShiftRight) Name() string {
	return "ShiftRight"
}

func (n ShiftRight) Attributes() map[string]interface{} {
	return nil
}

func (n ShiftRight) Walk(v node.Visitor) {
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
