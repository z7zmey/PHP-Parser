package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type EncapsedStringPart struct {
	name       string
	attributes map[string]interface{}
}

func NewEncapsedStringPart(value string) node.Node {
	return EncapsedStringPart{
		"EncapsedStringPart",
		map[string]interface{}{
			"value": value,
		},
	}
}

func (n EncapsedStringPart) Name() string {
	return "EncapsedStringPart"
}

func (n EncapsedStringPart) Attributes() map[string]interface{} {
	return nil
}

func (n EncapsedStringPart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
