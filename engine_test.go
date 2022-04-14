package json_spanner

import (
	"encoding/json"
	"fmt"
	"github.com/fuhongbo/json_spanner/gjson"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestEngine_SIMPLE_FILTER_Valid(t *testing.T) {

	Convey("Given a filter engine", t, func() {
		eng, _ := NewFilter("test.a.b=1 and (b=2 or c=4)")

		Convey("When give a json match this condition", func() {
			result := eng.Valid(`{"test":{"a":{"b":1}},"b":3,"c":4}`)
			So(result, ShouldBeTrue)
		})

		Convey("When give a json not match this condition", func() {
			result := eng.Valid(`{"test":{"a":1},"b":3,"c":4}`)
			So(result, ShouldBeFalse)
		})
	})

}

func TestEngine_SIMPLE_FILTER_DEC_Valid(t *testing.T) {

	//当前不支持
	Convey("Given a filter engine", t, func() {
		eng, err := NewFilter("c-b=1")

		So(err, ShouldNotBeNil)
		if eng != nil {
			eng.PrintTravel()
		}

	})

}

func TestEngine_SIMPLE_QUERY_Valid(t *testing.T) {

	Convey("Given a filter engine", t, func() {
		eng, _ := NewQuery("select test.a as x.t, 6.8 as x.d where test.a.b=1 and (b=2 or c=4)")

		Convey("When give a json match this condition", func() {
			result := eng.Valid(`{"test":{"a":{"b":1}},"b":3,"c":4}`)
			So(result, ShouldBeTrue)
		})

		Convey("When give a json not match this condition", func() {
			result := eng.Valid(`{"test":{"a":1},"b":3,"c":4}`)
			So(result, ShouldBeFalse)
		})
	})

}

func TestEngine_SIMPLE_QUERY_with_source_Valid(t *testing.T) {

	Convey("Given a filter engine", t, func() {
		eng, err := NewQuery("select test.a as x.t, 6.8 as x.d from 'test/+/bar' where test.a.b=1 and (b=2 or c=4)")

		if err != nil {
			t.Error(err.Error())
			return
		}

		Convey("When give a json match this condition", func() {
			result := eng.Valid(`{"test":{"a":{"b":1}},"b":3,"c":4}`)
			fmt.Println(eng.GetSource())
			So(result, ShouldBeTrue)
		})
	})

}

func TestEngine_SIMPLE_QUERY_Transform(t *testing.T) {

	Convey("Given a query for engine", t, func() {
		eng, _ := NewQuery("select test.a as x.t, 6.8 as x.d,b as x.b,c where test.a=1 and (b=2 or c=4)")

		Convey("When give a json match this condition", func() {
			result, jsonStr, err := eng.Transform(`{"test":{"a":1},"b":3,"c":4}`)
			So(result, ShouldBeTrue)
			So(err, ShouldBeNil)

			var json1 map[string]interface{}
			var json2 map[string]interface{}
			_ = json.Unmarshal([]byte(jsonStr), &json1)
			_ = json.Unmarshal([]byte(`{
    "c": 4,
    "x": {
        "b": 3,
        "d": 6.8,
        "t": 1
    }
}`), &json2)
			So(reflect.DeepEqual(json1, json2), ShouldBeTrue)

		})
	})
}

func TestEngine_SIMPLE_QUERY_Transform_SELETCTALL(t *testing.T) {

	Convey("Given a query for engine", t, func() {
		eng, _ := NewQuery("select * where test.a=1 and (b=2 or c=4)")

		Convey("When give a json match this condition", func() {
			result, jsonStr, err := eng.Transform(`{"test":{"a":1},"b":3,"c":4}`)
			So(result, ShouldBeTrue)
			So(err, ShouldBeNil)

			var json1 map[string]interface{}
			var json2 map[string]interface{}
			_ = json.Unmarshal([]byte(jsonStr), &json1)
			_ = json.Unmarshal([]byte(`{
    "c": 4,
    "x": {
        "b": 3,
        "d": 6.8,
        "t": 1
    }
}`), &json2)
			So(reflect.DeepEqual(json1, json2), ShouldBeTrue)

		})
	})
}

func TestEngine_SIMPLE_QUERY_Array_Transform_SELETCTALL(t *testing.T) {

	Convey("Given a query for engine", t, func() {
		eng, err := NewQuery("SELECT '1c75d89474a06b2c1625b040d0901d14' as instanceId,'https://env01.abilitylite01.abb.com.cn/iotapi' as hubApi,deviceId,header as data.header,'body.#(objectId~c4558a26-1089-43c1-9e53-61bf87b20706)#' as data.body,tenantId,solutionId WHERE (header.msgType='event' or header.`ability-messagetype`='event' or header.msgType='variable' or header.`ability-messagetype`='timeSeries' or header.`ability-messagetype`='variable' or header.msgType='alarm' or header.`ability-messagetype`='alarm') and 'body.#(objectId~c4558a26-1089-43c1-9e53-61bf87b20706)#'>0 and tenantId='09d7ce82-a6d8-4465-989f-d0dece0de1fd'")

		if err != nil {
			println(err.Error())
		}

		Convey("When give a json match this condition", func() {
			result, jsonStr, err := eng.Transform(`{
    "body": [
      {
        "model": "abb.ability.device",
        "objectId": "c4558a26-1089-43c1-9e53-61bf87b20706",
        "timestamp": "2022-04-14T00:12:02.845543Z",
        "value": "bd62356a-64c8-4bd5-b5b6-80cefa4c03c7",
        "variable": "var1"
      },
      {
        "model": "abb.ability.device",
        "objectId": "bbe4ab5e-eb31-4d48-9684-11a13fb4cb84",
        "timestamp": "2022-04-14T00:12:02.845545Z",
        "value": "6d65a584-98a1-4184-bda8-e418ba2dc239",
        "variable": "var2"
      },
      {
        "model": "abb.ability.device",
        "objectId": "bbe4ab5e-eb31-4d48-9684-11a13fb4cb84",
        "timestamp": "2022-04-14T00:12:02.845547Z",
        "value": "94bf6d88-1f26-43ac-b58a-27031931d4dd",
        "variable": "var3"
      },
      {
        "model": "abb.ability.device",
        "objectId": "bbe4ab5e-eb31-4d48-9684-11a13fb4cb84",
        "timestamp": "2022-04-14T00:12:02.845550Z",
        "value": "06815277-c095-4d66-af08-8b39898b369c",
        "variable": "var4"
      },
      {
        "model": "abb.ability.device",
        "objectId": "bbe4ab5e-eb31-4d48-9684-11a13fb4cb84",
        "timestamp": "2022-04-14T00:12:02.845551Z",
        "value": "0a99eb35-de1f-45b9-a5fc-d4aaaf3f2701",
        "variable": "var5"
      }
    ],
    "header": {
      "ability-messagetype": "timeSeries"
    }
"solutionId": "32184099-4676-4df2-aa96-794e3bd2ecc3",
  "tenantId": "09d7ce82-a6d8-4465-989f-d0dece0de1fd"
  }`)
			So(result, ShouldBeTrue)
			So(err, ShouldBeNil)

			var json1 map[string]interface{}
			var json2 map[string]interface{}
			_ = json.Unmarshal([]byte(jsonStr), &json1)
			_ = json.Unmarshal([]byte(`{
    "c": 4,
    "x": {
        "b": 3,
        "d": 6.8,
        "t": 1
    }
}`), &json2)
			So(reflect.DeepEqual(json1, json2), ShouldBeTrue)

		})
	})
}

func TestEngine_SIMPLE_QUERY_Object_Transform_SELETCTALL(t *testing.T) {

	Convey("Given a query for engine", t, func() {
		eng, err := NewQuery("SELECT '1c75d89474a06b2c1625b040d0901d14' as instanceId,'https://env01.abilitylite01.abb.com.cn/iotapi' as hubApi,deviceId,header as data.header,'body.#(objectId~c4558a26-1089-43c1-9e53-61bf87b20706)#' as data.body,tenantId,solutionId WHERE (header.msgType='event' or header.`ability-messagetype`='event' or header.msgType='variable' or header.`ability-messagetype`='timeSeries' or header.`ability-messagetype`='variable' or header.msgType='alarm' or header.`ability-messagetype`='alarm') and ('body.#(objectId~c4558a26-1089-43c1-9e53-61bf87b20706)#'>0 or body.objectId='c4558a26-1089-43c1-9e53-61bf87b20706') and tenantId='09d7ce82-a6d8-4465-989f-d0dece0de1fd'")

		if err != nil {
			println(err.Error())
		}

		Convey("When give a json match this condition", func() {
			result, jsonStr, err := eng.Transform(`{
    "body": 
      {
        "model": "abb.ability.device",
        "objectId": "c4558a26-1089-43c1-9e53-61bf87b20706",
        "timestamp": "2022-04-14T00:12:02.845543Z",
        "value": "bd62356a-64c8-4bd5-b5b6-80cefa4c03c7",
        "variable": "var1"
      },
    "header": {
      "ability-messagetype": "timeSeries"
    }
"solutionId": "32184099-4676-4df2-aa96-794e3bd2ecc3",
  "tenantId": "09d7ce82-a6d8-4465-989f-d0dece0de1fd"
  }`)
			So(result, ShouldBeTrue)
			So(err, ShouldBeNil)

			var json1 map[string]interface{}
			var json2 map[string]interface{}
			_ = json.Unmarshal([]byte(jsonStr), &json1)
			_ = json.Unmarshal([]byte(`{
    "c": 4,
    "x": {
        "b": 3,
        "d": 6.8,
        "t": 1
    }
}`), &json2)
			So(reflect.DeepEqual(json1, json2), ShouldBeTrue)

		})
	})
}

func TestSS(t *testing.T) {

	js := `{"test":[{"a":1},{"a":2},{"a":3}],"b":3,"c":4}`
	val := gjson.Get(js, `test.#(a~2,3)#`).Array()
	if len(val) > 0 {
		fmt.Println(len(val))
		fmt.Printf("%v\n", val)
	}
}
