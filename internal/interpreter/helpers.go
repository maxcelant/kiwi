package interpreter

import "fmt"

func Stringify(obj any) (string, error) {
	if obj == nil {
		return "nil", nil
	}
	if v, ok := obj.(int); ok {
		return fmt.Sprintf("%d", v), nil
	}
	if v, ok := obj.(float64); ok {
		return fmt.Sprintf("%f", v), nil
	}
	if v, ok := obj.(string); ok {
		return v, nil
	}
	if v, ok := obj.(bool); ok {
		return fmt.Sprintf("%t", v), nil
	}
	return "", fmt.Errorf("unsupported type: %T", obj)
}

func IsTruthy(v any) bool {
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return true
}

func isNumber(v any) (int, bool) {
	switch v.(type) {
	case int, float64:
		return v.(int), true
	default:
		return 0.0, false
	}
}

func BothOperandsNumbers(a any, b any) (int, int, bool) {
	left, ok := isNumber(a)
	if !ok {
		return 0, 0, false
	}
	right, ok := isNumber(b)
	if !ok {
		return 0, 0, false
	}
	return left, right, true
}

func isString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}
