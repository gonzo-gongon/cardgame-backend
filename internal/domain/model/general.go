package model

import "time"

type Created struct {
	By *User
	At *time.Time
}

type Updated struct {
	By *User
	At *time.Time
}

type Deleted struct {
	By *User
	At *time.Time
}
