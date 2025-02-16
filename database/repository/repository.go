package repository

type Repository[T any] interface {
	FindById(ID uint) (*T, error)
	Update(object *T) error
	Add(object *T) error
}
