package model

type UUID[T any] string

func (u *UUID[T]) String() string {
	return string(*u)
}

func UUIDFromString[T any](s string) UUID[T] {
	return UUID[T](s)
}

func UUIDsFromString[T any](s []string) []UUID[T] {
	r := make([]UUID[T], len(s))

	for i, v := range s {
		r[i] = UUIDFromString[T](v)
	}

	return r
}
