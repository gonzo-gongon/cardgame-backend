package model

type UUID[T any] string

func (u *UUID[T]) String() string {
	return string(*u)
}
