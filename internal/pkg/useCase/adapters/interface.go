package adapters

// Storage Interface for interacting with the use case repository.
type Storage interface {
	User
	Text
	Binary
	Card
	Account
}
