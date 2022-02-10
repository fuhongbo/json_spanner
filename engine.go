package json_spanner

import (
	"errors"
	"github.com/fuhongbo/json_spanner/parser"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Engine struct {
	workType        int // 0 is filter ,1 is query
	statement       parser.Statements
	selectStatement parser.SelectStatement
}

func NewFilter(filter string) (engine *Engine, errs error) {

	lex := parser.NewLex("where " + filter)

	statements := parser.Parse(lex)

	if lex.ErrorInfo != nil {
		return nil, lex.ErrorInfo
	}

	return &Engine{statement: statements, selectStatement: parser.Info(statements[0]), workType: 0}, nil
}

func NewQuery(query string) (engine *Engine, errs error) {
	lex := parser.NewLex(query)
	statements := parser.Parse(lex)
	if lex.ErrorInfo != nil {
		return nil, lex.ErrorInfo
	}
	return &Engine{statement: statements, selectStatement: parser.Info(statements[0]), workType: 1}, nil
}

func (e *Engine) Valid(json string) bool {

	jsonRef := gjson.Parse(json)
	return e.valid(&jsonRef)

}

func (e *Engine) valid(json *gjson.Result) bool {
	return Valid(e.selectStatement.WhereCondition.Left, e.selectStatement.WhereCondition.Right, e.selectStatement.WhereCondition.Op, json)
}

func (e *Engine) GetSource() string {
	if e.selectStatement.Source != nil {
		return e.selectStatement.Source.Name
	} else {
		return ""
	}
}

func (e *Engine) Transform(json string) (match bool, result string, err error) {
	if e.workType != 1 {
		return false, "", errors.New("transform can only be used in query mode . ")
	}

	jsonPath := gjson.Parse(json)

	if e.valid(&jsonPath) {

		resultJson := `{}`

		for _, item := range e.selectStatement.Fields {
			switch item.Type {
			case "NAME":
				if item.Alias != "" {
					resultJson, _ = sjson.Set(resultJson, item.Alias, jsonPath.Get(item.Name).Value())
				} else {
					resultJson, _ = sjson.Set(resultJson, item.Name, jsonPath.Get(item.Name).Value())
				}
			case "STRING":
				resultJson, _ = sjson.Set(resultJson, item.Alias, item.ValueString)
			case "FLOAT":
				resultJson, _ = sjson.Set(resultJson, item.Alias, item.ValueFloat)
			case "INT":
				resultJson, _ = sjson.Set(resultJson, item.Alias, item.ValueInt)
			}
		}

		return true, resultJson, nil

	} else {
		return false, "", errors.New("Data does not match filter criteria . ")
	}
}

func (e *Engine) PrintTravel() {
	parser.Travel(e.statement)
}
