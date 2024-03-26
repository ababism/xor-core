package repo_models

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type PurchaseRequest struct {
	UUID       uuid.UUID   `db:"uuid"`
	Sender     uuid.UUID   `db:"sender"`
	Receiver   uuid.UUID   `db:"receiver"`
	Products   []uuid.UUID `db:"products"`
	WebhookURL string      `db:"webhook_url"`
	CreatedAt  time.Time   `db:"created_at"`
}

func CreateToPurchaseRequestPostgres(model *domain.PurchaseRequestCreate) *PurchaseRequest {
	id, _ := uuid.NewUUID()
	return &PurchaseRequest{
		UUID:       id,
		Sender:     model.Sender,
		Receiver:   model.Receiver,
		Products:   model.Products,
		WebhookURL: model.WebhookURL,
		CreatedAt:  time.Now(),
	}
}

func ToPurchaseRequestDomain(model *PurchaseRequest) *domain.PurchaseRequestGet {
	return &domain.PurchaseRequestGet{
		UUID:       model.UUID,
		Sender:     model.Sender,
		Receiver:   model.Receiver,
		Products:   model.Products,
		WebhookURL: model.WebhookURL,
		ReceivedAt: model.CreatedAt,
	}
}
