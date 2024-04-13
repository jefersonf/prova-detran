package storage

type Repository interface {
	Save()
	Find()
	Count()
}
