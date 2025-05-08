package order

import "github.com/Vractos/ecoffe-go/entity"

type UseCase interface {
	CreateOrder(input CreateOrderDtoInput) (*entity.ID, error)
	RetrieveAllOrders() (*[]entity.Order, error)
}

/*
#########################################
#########################################
---------------REPOSITORY---------------
#########################################
#########################################
*/

type RepoWriter interface {
	CreateOrder(o *entity.Order) error
}

type RepoReader interface {
	RetrieveAllOrders() (*[]entity.Order, error)
}

type Repository interface {
	RepoWriter
	RepoReader
}
