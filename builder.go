package json_spanner

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Builder struct {
	jsonMap map[string]interface{}
}

func NewBuilder() *Builder {
	return &Builder{jsonMap: make(map[string]interface{})}
}

func (b *Builder) Set(key string, value interface{}) {
	b.jsonMap = b.setInternal(key, b.jsonMap, value)
}

func (b *Builder) setInternal(key string, p map[string]interface{}, value interface{}) map[string]interface{} {
	keys := strings.Split(key, ".")
	if p == nil {
		p = make(map[string]interface{})
	}
	if len(keys) == 1 {
		p[keys[0]] = value
		return p
	} else {
		if _, ok := p[keys[0]]; ok {
			if reflect.TypeOf(p[keys[0]]) == reflect.TypeOf(p) {
				p[keys[0]] = b.setInternal(strings.Join(keys[1:], "."), p[keys[0]].(map[string]interface{}), value)
				return p
			} else {
				p[keys[0]] = make(map[string]interface{})
				p[keys[0]] = b.setInternal(strings.Join(keys[1:], "."), p[keys[0]].(map[string]interface{}), value)
				return p
			}

		} else {
			p[keys[0]] = make(map[string]interface{})
			p[keys[0]] = b.setInternal(strings.Join(keys[1:], "."), p[keys[0]].(map[string]interface{}), value)
			return p
		}
	}
}

func (b *Builder) ToJson() string {
	jsonBytes, _ := json.Marshal(b.jsonMap)
	return string(jsonBytes)
}
