package reflection

import (
	"reflect"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {
	/*
		val := getValue(x)

		numOfVals := 0
		var getField func(int) reflect.Value

		switch val.Kind() {
		case reflect.String:
			fn(val.String())
		case reflect.Slice, reflect.Array:
			numOfVals = val.Len()
			getField = val.Index
		case reflect.Struct:
			numOfVals = val.NumField()
			getField = val.Field
		case reflect.Map:
			for _, key := range val.MapKeys() {
				//fn(val.MapIndex(key).String())
				walk(val.MapIndex(key).Interface(), fn)
			}
		}
		for i := 0; i < numOfVals; i++ {
			walk(getField(i).Interface(), fn)
		}
	*/
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, result := range valFnResult {
			walkValue(result)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// cant use NumField at pointers
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
