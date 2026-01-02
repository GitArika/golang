// Package examples
package examples

func Contains[T comparable](s []T, compareTo T) bool {
	for _, str := range s {
		if str == compareTo {
			return true
		}
	}
	return false
}
