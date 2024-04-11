package financesclient

import (
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/financesclient/generated"
)

func ToProductCreateRequest(product domain.Product) generated.ProductCreate {
	return generated.ProductCreate{
		Info:        product.ID.String(),
		IsAvailable: true,
		Name:        product.Item.String(),
		Price:       product.Price,
	}
}

//
//func ToDriverLocationDomain(r generated.Driver) (*domain.DriverLocation, error) {
//	if r.Id == nil {
//		return nil, errors.New("driver id is nil")
//	}
//	driverId, err := uuid.Parse(*r.Id)
//	if err != nil {
//		return nil, err
//	}
//	res := domain.DriverLocation{
//		DriverId: driverId,
//		Coordinates: domain.LatLngLiteral{
//			Lat: r.Lat,
//			Lng: r.Lng,
//		},
//	}
//	return &res, nil
//}
//
//func ToDriverLocationsDomain(dLocations []generated.Driver) ([]domain.DriverLocation, error) {
//	driversDomain := make([]domain.DriverLocation, len(dLocations))
//
//	for i, driverLocation := range dLocations {
//		dlDom, err := ToDriverLocationDomain(driverLocation)
//		if err != nil {
//			return nil, err
//		}
//		driversDomain[i] = *dlDom
//	}
//
//	return driversDomain, nil
//}
