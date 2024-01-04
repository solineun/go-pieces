package stubs

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(userID string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

type User struct {
	Id string
}

type Person struct {
	User
	Name string
}

type Pet struct {
	Type string
}