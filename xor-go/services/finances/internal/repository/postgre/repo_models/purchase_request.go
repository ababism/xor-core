package repo_models

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type PurchaseRequest struct {
	UUID       uuid.UUID   `db:"uuid"`
	Sender     *uuid.UUID  `db:"sender"`
	Receiver   *uuid.UUID  `db:"receiver"`
	Amount     float32     `db:"amount"`
	Products   []uuid.UUID `db:"products"`
	WebhookURL string      `db:"webhook_url"`
	CreatedAt  time.Time   `db:"created_at"`
}

func CreateToPurchaseRequestPostgres(model *domain.PurchaseRequestCreate, amount float32) *PurchaseRequest {
	id, _ := uuid.NewUUID()
	return &PurchaseRequest{
		UUID:       id,
		Sender:     model.Sender,
		Receiver:   model.Receiver,
		Amount:     amount,
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
		Amount:     model.Amount,
		Products:   model.Products,
		WebhookURL: model.WebhookURL,
		CreatedAt:  model.CreatedAt,
	}
}
