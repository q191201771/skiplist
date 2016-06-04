package skip_list

import (
	"reflect"
)

func CommonCompare(l, r interface{}) int {
	switch reflect.TypeOf(l).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IntCompare(l.(int), r.(int))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return UintCompare(l.(uint), r.(uint))
	case reflect.Float32, reflect.Float64:
		return FloatCompare(l.(float64), r.(float64))
	case reflect.String:
		return StringCompare(l.(string), r.(string))
	}
	return 0
}

func IntCompare(l, r int) int {
	if l < r {
		return -1
	} else if l == r {
		return 0
	}
	return 1

}

func UintCompare(l, r uint) int {
	if l < r {
		return -1
	} else if l == r {
		return 0
	}
	return 1

}

func FloatCompare(l, r float64) int {
	if l < r {
		return -1
	} else if l == r {
		return 0
	}
	return 1

}

func StringCompare(l, r string) int {
	if l < r {
		return -1
	} else if l == r {
		return 0
	}
	return 1
}
