package ast

import (
	"fmt"
	"reflect"
)

func EncodeValue(input interface{}) (Argument, error) {
	switch input := input.(type) {
	case Argument:
		return input, nil
	case string:
		return NewText(input), nil
	case float64, float32:
		return NewNumber(reflect.ValueOf(input).Float()), nil
	case int, int8, int32, int64:
		// loss casting
		return NewNumber(float64(reflect.ValueOf(input).Int())), nil
	case bool:
		if input {
			return True(), nil
		} else {
			return False(), nil
		}
	}
	return nil, fmt.Errorf("unable to encode object of type %T into Argument", input)
}

var (
	trueSym = Symbol{
		sym: "true",
	}
	falseSym = Symbol{
		sym: "false",
	}
)

func True() Symbol {
	return trueSym
}

func False() Symbol {
	return falseSym
}
