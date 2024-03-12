package dto_models

import (
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/handler/generated"
)

func ToTripResponse(t domain.Trip) generated.Trip {
	return generated.Trip{
		DriverId: t.DriverId,
		From:     ToLatLngLiteralResponse(t.From),
		Id:       &t.Id,
		Price:    ToMoneyResponse(t.Price),
		Status:   ToTripStatusResponse(t.Status),
		To:       ToLatLngLiteralResponse(t.To),
	}
}

func ToTripsResponse(t []domain.Trip) []generated.Trip {
	var tripsResponse []generated.Trip

	for _, trip := range t {
		tripsResponse = append(tripsResponse, ToTripResponse(trip))
	}

	return tripsResponse
}

func ToLatLngLiteralResponse(lll *domain.LatLngLiteral) *generated.LatLngLiteral {
	return &generated.LatLngLiteral{
		Lat: lll.Lat,
		Lng: lll.Lng,
	}
}

func ToMoneyResponse(m *domain.Money) *generated.Money {
	return &generated.Money{
		Amount:   m.Amount,
		Currency: m.Currency,
	}
}
func ToTripStatusResponse(ts *domain.TripStatus) *generated.TripStatus {
	var res generated.TripStatus
	switch *ts {
	case "CANCELED":
		res = generated.CANCELED
	case "DRIVER_FOUND":
		res = generated.DRIVERFOUND
	case "DRIVER_SEARCH":
		res = generated.DRIVERSEARCH
	case "ENDED":
		res = generated.ENDED
	case "ON_POSITION":
		res = generated.ONPOSITION
	case "STARTED":
		res = generated.STARTED
	default:
		return nil
	}
	return &res
}
