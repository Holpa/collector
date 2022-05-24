package utils

import "math"

type FloatConvertible interface {
	~int | ~uint | float64
}

func Min[T FloatConvertible](values ...T) T {
	if len(values) == 0 {
		return 0
	}
	absMin := values[0]

	for _, value := range values {
		absMin = T(math.Min(float64(absMin), float64(value)))
	}

	return absMin
}
