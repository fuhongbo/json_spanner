package json_spanner

import (
	"fmt"
	"github.com/tidwall/gjson"
	parser "json_spanner/parser"
	"json_spanner/wildcard"
	"reflect"
)

var floatType = reflect.TypeOf(float64(0))

func Valid(left *parser.ExprNode, right *parser.ExprNode, op int, json *gjson.Result) bool {

	if left == nil && right == nil {
		return true
	} else {
		switch op {

		case parser.AND: // AND
			return Valid(left.Left, left.Right, left.Op, json) && Valid(right.Left, right.Right, right.Op, json)
		case parser.OR: //  OR
			return Valid(left.Left, left.Right, left.Op, json) || Valid(right.Left, right.Right, right.Op, json)
		default:
			return compare(getValue(left, json), getValue(right, json), op)
		}
	}

}

func compare(a, b interface{}, op int) bool {

	switch a.(type) {
	case float64:
		tempA, err := getFloat(a)
		if err != nil {
			return false
		}
		tempB, err := getFloat(b)
		if err != nil {
			return false
		}

		switch op {
		case parser.EQ: //   =
			return tempA == tempB
		case parser.NEQ: //  <>
			return tempA != tempB
		case parser.GT: //   >
			return tempA > tempB
		case parser.GTE: //  >=
			return tempA >= tempB
		case parser.LT: //   <
			return tempA < tempB
		case parser.LTE: // <=
			return tempA <= tempB
		default:
			return false
		}

	default:
		switch op {
		case parser.EQ: //   =
			return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
		case parser.NEQ: //  <>
			return fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b)
		case parser.GT: //   >
			return fmt.Sprintf("%v", a) > fmt.Sprintf("%v", b)
		case parser.GTE: //  >=
			return fmt.Sprintf("%v", a) >= fmt.Sprintf("%v", b)
		case parser.LT: //   <
			return fmt.Sprintf("%v", a) < fmt.Sprintf("%v", b)
		case parser.LTE: // <=
			return fmt.Sprintf("%v", a) <= fmt.Sprintf("%v", b)
		case parser.TLIKE: // =~
			rs := wildcard.Match(fmt.Sprintf("%v", a), fmt.Sprintf("%v", b))
			if rs == nil {
				return false
			} else {
				return true
			}
		default:
			return false
		}

	}

}

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

func getValue(node *parser.ExprNode, json *gjson.Result) interface{} {
	switch node.Type {
	case parser.FieldNode:
		return json.Get(node.Name).Value()
	case parser.FloatNode:
		return node.FloVal
	case parser.IntegerNode:
		return node.IntVal
	case parser.StringNode:
		return node.StrVal
	}
	return nil
}
