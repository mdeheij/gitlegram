package interfaces

type RequestInterface interface {
	GetRepository() RepositoryInterface
}

type RepositoryInterface interface {
	GetName() string
}
