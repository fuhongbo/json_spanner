package json_spanner

import (
	"fmt"
	"testing"
)

func BenchmarkCondition_Match(b *testing.B) {
	eng, _ := NewFilter("test.a.b=1 and (b=2 or c=4)")

	for i := 0; i < b.N; i++ {

		result := eng.Valid(`{"test":{"a":{"b":1}},"b":3,"c":4}`)

		if !result {
			b.Error("条件不匹配")
		}
	}
}

func BenchmarkCondition_Transform(b *testing.B) {
	eng, _ := NewQuery("select test.a as x.t, 6.8 as x.d,b as x.b,c where test.a=1 and (b=2 or c=4)")

	for i := 0; i < b.N; i++ {

		result, _, err := eng.Transform(`{"test":{"a":1},"b":3,"c":4}`)

		if !result {
			fmt.Printf("The result is %v \n", result)
			return
		}

		if err != nil {
			fmt.Printf("has error %v \n", err)
			return
		}

		//fmt.Println(json)
	}
}
