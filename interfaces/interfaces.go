package interfaces

type RequestInterface interface {
	GetRepository() RepositoryInterface
	GetUser() (UserInterface, error)
}

type RepositoryInterface interface {
	GetName() string
}
type UserInterface interface {
	GetUsername() string
	GetName() string
	GetAvatarURL() string
}
