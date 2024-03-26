package repo_models

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type PayoutRequestData struct {
}

type PayoutRequest struct {
	UUID      uuid.UUID         `db:"uuid"`
	Receiver  uuid.UUID         `db:"receiver"`
	Amount    float32           `db:"amount"`
	Data      PayoutRequestData `db:"data"`
	CreatedAt time.Time         `db:"created_at"`
}

func CreateToPayoutRequestDataPostgres(model domain.PayoutRequestData) PayoutRequestData {
	return PayoutRequestData{}
}

func CreateToPayoutRequestPostgres(model *domain.PayoutRequestCreate) *PayoutRequest {
	id, _ := uuid.NewUUID()
	return &PayoutRequest{
		UUID:      id,
		Receiver:  model.Receiver,
		Amount:    model.Amount,
		Data:      CreateToPayoutRequestDataPostgres(model.Data),
		CreatedAt: time.Now(),
	}
}

func ToPayoutRequestDataDomain(model PayoutRequestData) domain.PayoutRequestData {
	return domain.PayoutRequestData{}
}

func ToPayoutRequestDomain(model *PayoutRequest) *domain.PayoutRequestGet {
	return &domain.PayoutRequestGet{
		UUID:       model.UUID,
		Receiver:   model.Receiver,
		Amount:     model.Amount,
		ReceivedAt: model.CreatedAt,
		Data:       ToPayoutRequestDataDomain(model.Data),
	}
}
