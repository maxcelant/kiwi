package interpreter

import "fmt"

type ComparatorFunc func(a any, b any) bool

func Compare(a any, b any, comparators ...ComparatorFunc) bool {
	for _, c := range comparators {
		if ok := c(a, b); ok {
			return ok
		}
	}
	return false
}

func WithString() ComparatorFunc {
	return func(a any, b any) bool {
		a, ok := a.(string)
		if !ok {
			return false
		}
		b, ok = b.(string)
		return ok
	}
}

func WithInt() ComparatorFunc {
	return func(a any, b any) bool {
		a, ok := a.(int)
		if !ok {
			return false
		}
		b, ok = b.(int)
		return ok
	}
}

func WithBool() ComparatorFunc {
	return func(a any, b any) bool {
		a, ok := a.(bool)
		if !ok {
			return false
		}
		b, ok = b.(bool)
		return ok
	}
}

func Stringify(obj any) string {
	if obj == nil {
		return "nil"
	}
	if v, ok := obj.(int); ok {
		return fmt.Sprintf("%d", v)
	}
	if v, ok := obj.(float64); ok {
		return fmt.Sprintf("%f", v)
	}
	if v, ok := obj.(string); ok {
		return v
	}
	if v, ok := obj.(bool); ok {
		return fmt.Sprintf("%t", v)
	}
	return fmt.Sprintf("unsupported type: %T", obj)
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
