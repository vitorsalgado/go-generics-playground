package arrays

type Arrays[T comparable] struct {
	Value []T
}

func Of[T comparable](list []T) *Arrays[T] {
	return &Arrays[T]{Value: list}
}

func (sa *Arrays[T]) Map(fn func(i T) T) *Arrays[T] {
	arr := make([]T, len(sa.Value))

	for i, v := range sa.Value {
		arr[i] = fn(v)
	}

	return Of(arr)
}

func (sa *Arrays[T]) Filter(pred func(i T) bool) *Arrays[T] {
	arr := make([]T, 0)

	for i, v := range sa.Value {
		if pred(v) {
			arr = append(arr, sa.Value[i])
		}
	}

	return Of(arr)
}

func (sa *Arrays[T]) Every(pred func(i T) bool) bool {
	for _, v := range sa.Value {
		if !pred(v) {
			return false
		}
	}

	return true
}

func (sa *Arrays[T]) Collect() []T {
	return sa.Value
}
