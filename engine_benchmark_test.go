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

func BenchmarkCondition_complex_Transform(b *testing.B) {
	eng, _ := NewQuery("select b,c,'test.#(a~!0315dd67-47df-4c33-aff2-8220f8f8ef36,8c0e023f-cc19-4197-ad01-c05dbcbab8dc)#' as test where 'test.#(a~!0315dd67-47df-4c33-aff2-8220f8f8ef36,8c0e023f-cc19-4197-ad01-c05dbcbab8dc)#'>0 and (b=2 or c=4)")

	for i := 0; i < b.N; i++ {

		result, reJson, err := eng.Transform(`{"test":[{"a":"8c0e023f-cc19-4197-ad01-c05dbcbab8dc"},{"a":"39adf134-2ef1-441b-9186-41575daa0417"},{"a":"0315dd67-47df-4c33-aff2-8220f8f8ef36"}],"b":3,"c":4}`)

		if !result {
			fmt.Printf("The result is %v \n", result)
			return
		}

		if err != nil {
			fmt.Printf("has error %v \n", err)
			return
		}

		if reJson != `{"b":3,"c":4,"test":[{"a":"39adf134-2ef1-441b-9186-41575daa0417"}]}` {
			fmt.Println("返回出错")
		}

		//fmt.Println(json)
	}
}
