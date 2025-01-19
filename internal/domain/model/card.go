package model

type Card struct {
	ID   UUID[Card]
	Name string
	Text string
}

type CreateCard struct {
	Name string
	Text string
}
