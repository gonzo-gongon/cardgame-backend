package model

type Card struct {
	ID        UUID[Card]
	Name      string
	Text      string
	CreatedBy UUID[User]
	UpdatedBy UUID[User]
}

type CreateCard struct {
	Name string
	Text string
}
