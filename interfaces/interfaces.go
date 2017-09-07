package interfaces

type Request interface {
	GetRepository() Repository
	GetUser() (User, error)
}

type Repository interface {
	GetName() string
}
type User interface {
	GetUsername() string
	GetName() string
	GetAvatarURL() string
}
