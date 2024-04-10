package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductSmall struct {
	UUID  uuid.UUID
	Name  string
	Info  string
	Price float32
}

type ProductGet struct {
	UUID          uuid.UUID
	Name          string
	Info          string
	Price         float32
	IsAvailable   bool
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}

type ProductCreate struct {
	Name        string
	Info        string
	Price       float32
	IsAvailable bool
}

type ProductUpdate struct {
	UUID        uuid.UUID
	Name        string
	Info        string
	Price       float32
	IsAvailable bool
}

type ProductFilter struct {
	UUID        *uuid.UUID
	Name        *string
	Price       *float32
	IsAvailable *bool
}

func ConvertProductToSmall(product ProductGet) ProductSmall {
	return ProductSmall{
		UUID:  product.UUID,
		Name:  product.Name,
		Info:  product.Info,
		Price: product.Price,
	}
}

func ConvertProductsToSmall(products []ProductGet) []ProductSmall {
	res := make([]ProductSmall, len(products))
	for i, product := range products {
		res[i] = ConvertProductToSmall(product)
	}
	return res
}
