package util

// First returns the first element of a slice or a zero value if the slice is empty.
func First[T any](s []T) T {
	if len(s) > 0 {
		return s[0]
	}
	var zero T
	return zero
}
