package kafkaproducer

import (
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
	"time"
)

func ToCommandTypeKafka(commandType domain.CommandType) CommandType {
	return CommandType(commandType)
}

func ToTripCommand(trip domain.Trip, commandType domain.CommandType, reason *string) Command {
	command := Command{
		DriverId:        *trip.DriverId,
		Source:          domain.Source,
		Type:            ToCommandTypeKafka(commandType),
		DataContentType: "application/json",
		Time:            time.Now(),
		Data: Data{
			TripID: trip.Id.String(),
			Reason: reason,
		},
	}

	return command
}
