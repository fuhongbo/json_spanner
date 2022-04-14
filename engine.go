package json_spanner

import (
	"errors"
	"github.com/fuhongbo/json_spanner/parser"
	"strings"

	//"github.com/tidwall/gjson"
	"github.com/fuhongbo/json_spanner/gjson"
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

func (e *Engine) Update(query string) error {
	lex := parser.NewLex(query)
	statements := parser.Parse(lex)
	if lex.ErrorInfo != nil {
		return lex.ErrorInfo
	}
	e.statement = statements
	e.selectStatement = parser.Info(statements[0])
	return nil
}

func (e *Engine) valid(json *gjson.Result) bool {
	if e.selectStatement.WhereCondition == nil {
		return true
	}
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

		b := NewBuilder()
		//resultJson := `{}`

		for _, item := range e.selectStatement.Fields {
			switch item.Type {
			case "ALL":
				return true, json, nil
			case "NAME":
				if item.Alias != "" {
					b.Set(item.Alias, jsonPath.Get(item.Name).Value())
					//resultJson, _ = sjson.Set(resultJson, item.Alias, jsonPath.Get(item.Name).Value())
				} else {
					b.Set(item.Name, jsonPath.Get(item.Name).Value())
					//resultJson, _ = sjson.Set(resultJson, item.Name, jsonPath.Get(item.Name).Value())
				}
			case "STRING":
				if strings.Contains(item.ValueString, ".#(") {
					ar := strings.Split(item.ValueString, ".#(")
					if jsonPath.Get(ar[0]).IsArray() {
						b.Set(item.Alias, jsonPath.Get(item.ValueString).Value())
					} else {
						b.Set(item.Alias, jsonPath.Get(ar[0]).Value())
					}

				} else {
					b.Set(item.Alias, item.ValueString)
				}

				//resultJson, _ = sjson.Set(resultJson, item.Alias, item.ValueString)
			case "FLOAT":
				b.Set(item.Alias, item.ValueFloat)
				//resultJson, _ = sjson.Set(resultJson, item.Alias, item.ValueFloat)
			case "INT":
				b.Set(item.Alias, item.ValueInt)
				//resultJson, _ = sjson.Set(resultJson,)
			}
		}

		return true, b.ToJson(), nil

	} else {
		return false, "", errors.New("Data does not match filter criteria . ")
	}
}

func (e *Engine) PrintTravel() {
	parser.Travel(e.statement)
}
