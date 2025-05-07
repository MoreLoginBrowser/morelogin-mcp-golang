package utils

import (
	"fmt"
	"math"
	"reflect"
)

// SafelyConvertToInt
func SafelyConvertToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float32:
		if float32(int(v)) == v {
			return int(v), nil
		}
		return 0, NewParamError("number", "must be an integer value")
	case float64:
		if float64(int(v)) == v {
			return int(v), nil
		}
		return 0, NewParamError("number", "must be an integer value")
	case string:
		var intValue int
		if _, err := fmt.Sscanf(v, "%d", &intValue); err == nil {
			return intValue, nil
		}
		var floatValue float64
		if _, err := fmt.Sscanf(v, "%f", &floatValue); err == nil {
			if math.Floor(floatValue) == floatValue {
				return int(floatValue), nil
			}
		}
		return 0, NewParamError("number", "must be a valid integer")
	default:
		return 0, NewParamError("number", fmt.Sprintf("unsupported type: %v", reflect.TypeOf(value)))
	}
}
