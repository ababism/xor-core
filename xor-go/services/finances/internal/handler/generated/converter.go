package generated

import "gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"

// TODO import OpenAPI constants

func ToTODOTypeDomain(ts TripStatus) domain.TripStatus {
	converted := domain.TripStatus(ts)
	return converted
}

func ScrapeStatusesConstants() {
	domain.InitTripStatusCollection(
		ToTODOTypeDomain(CANCELED),
		ToTODOTypeDomain(DRIVERFOUND),
		ToTODOTypeDomain(DRIVERSEARCH),
		ToTODOTypeDomain(ENDED),
		ToTODOTypeDomain(ONPOSITION),
		ToTODOTypeDomain(STARTED))
}
