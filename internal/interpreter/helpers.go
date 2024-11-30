package interpreter

func Stringify(obj any) (string, error) {
	return "", nil
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
