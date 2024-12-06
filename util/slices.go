package util

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Reduce[T Numeric, V any](arr []V, fn func(curr V) T, initial T) T {
	acc := initial
	for _, curr := range arr {
		acc += fn(curr)
	}
	return acc
}
