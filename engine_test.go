package json_spanner

import (
	"encoding/json"
	"fmt"
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

func TestEngine_Platform_Filter_Valid(t *testing.T) {

	Convey("Given a filter engine", t, func() {
		eng, _ := NewFilter("header.`ability-messagetype`='timeSeries'")

		Convey("When give a json match this condition", func() {
			result := eng.Valid(`{
    "header": {
        "ability-messagetype": "timeSeries",
        "x": "y"
    },
    "body": {
        "objectId": "5ee55388-ab19-43ce-a1e6-074f68ae6edb",
        "timestamp": "2019-05-30T09:21:14.7629753Z",
        "model": "abb.ability.device",
        "variable": "skinTemperature",
        "value": 300,
        "quality": 1023
    }
}`)
			So(result, ShouldBeTrue)
		})

		Convey("When give a json not match this condition", func() {
			result := eng.Valid(`{
    "header": {
        "ability-messagetype": "event",
        "x": "y"
    },
    "body": {
        "objectId": "5ee55388-ab19-43ce-a1e6-074f68ae6edb",
        "timestamp": "2019-05-30T09:21:14.7629753Z",
        "model": "abb.ability.device",
        "variable": "skinTemperature",
        "value": 300,
        "quality": 1023
    }
}`)
			So(result, ShouldBeFalse)
		})
	})
}

func TestEngine_Platform_Filter_Transform(t *testing.T) {

	Convey("Given a filter engine", t, func() {
		eng, _ := NewQuery("SELECT 'bfae53f0-e0f7-44d7-879b-b94266871929' as instanceId,'https://env01.abilitylite02.abb.com.cn/iotapi' as hubApi,clientid as deviceId, header as data.header, body as data.body where header.`ability-messagetype`='timeSeries'")

		Convey("When give a json match this condition", func() {
			result, jsonStr, err := eng.Transform(`{
    "header": {
        "ability-messagetype": "timeSeries"
    },
    "body": {
        "objectId": "5ee55388-ab19-43ce-a1e6-074f68ae6edb",
        "timestamp": "2019-05-30T09:21:14.7629753Z",
        "model": "abb.ability.device",
        "variable": "skinTemperature",
        "value": 300,
        "quality": 1023
    }
}`)
			So(result, ShouldBeTrue)
			So(err, ShouldBeNil)

			var json1 map[string]interface{}
			var json2 map[string]interface{}
			_ = json.Unmarshal([]byte(jsonStr), &json1)
			_ = json.Unmarshal([]byte(`{
    "data": {
        "body": {
            "model": "abb.ability.device",
            "objectId": "5ee55388-ab19-43ce-a1e6-074f68ae6edb",
            "quality": 1023,
            "timestamp": "2019-05-30T09:21:14.7629753Z",
            "value": 300,
            "variable": "skinTemperature"
        },
        "header": {
            "ability-messagetype": "timeSeries"
        }
    },
    "deviceId": null,
    "hubApi": "https://env01.abilitylite02.abb.com.cn/iotapi",
    "instanceId": "bfae53f0-e0f7-44d7-879b-b94266871929"
}`), &json2)
			So(reflect.DeepEqual(json1, json2), ShouldBeTrue)
		})

	})
}

func Test_Debug(t *testing.T) {
	eng, err := NewFilter("header.`ability-messagetype`='timeSeries'")

	if err != nil {
		t.Error(err)
		return
	}

	result := eng.Valid(`{
    "header": {
        "ability-messagetype": "timeSeries",
        "x": "y"
    },
    "body": {
        "objectId": "5ee55388-ab19-43ce-a1e6-074f68ae6edb",
        "timestamp": "2019-05-30T09:21:14.7629753Z",
        "model": "abb.ability.device",
        "variable": "skinTemperature",
        "value": 300,
        "quality": 1023
    }
}`)
	t.Log(result)
}
