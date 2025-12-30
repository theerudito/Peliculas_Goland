package helpers

import (
	"fmt"
	"strconv"
)

func ConvertToInt(val interface{}) (int, error) {
	switch v := val.(type) {
	case float64:
		return int(v), nil
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return n, nil
	default:
		return 0, fmt.Errorf("invalid type: %T", val)
	}
}

func ConvertToUInt(val interface{}) (uint, error) {
	switch v := val.(type) {
	case float64:
		return uint(v), nil
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return uint(n), nil
	default:
		return 0, fmt.Errorf("invalid type: %T", val)
	}
}
