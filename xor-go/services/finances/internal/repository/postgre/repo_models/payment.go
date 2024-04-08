package repo_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type PaymentData struct {
}

func (a *PaymentData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *PaymentData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Payment struct {
	UUID      uuid.UUID   `db:"uuid"`
	Sender    uuid.UUID   `db:"sender"`
	Receiver  uuid.UUID   `db:"receiver"`
	Data      PaymentData `db:"data"`
	URL       string      `db:"url"`
	Status    string      `db:"status"`
	EndedAt   time.Time   `db:"ended_at"`
	CreatedAt time.Time   `db:"created_at"`
}

func CreateToPaymentPostgres(model *domain.PaymentCreate) *Payment {
	id, _ := uuid.NewUUID()
	return &Payment{
		UUID:      id,
		Sender:    model.Sender,
		Receiver:  model.Receiver,
		Data:      ToPaymentDataPostgres(model.Data),
		URL:       model.URL,
		Status:    model.Status,
		EndedAt:   model.EndedAt,
		CreatedAt: time.Now(),
	}
}

func ToPaymentDataPostgres(model domain.PaymentData) PaymentData {
	return PaymentData{}
}

func ToPaymentDomain(model *Payment) *domain.PaymentGet {
	return &domain.PaymentGet{
		UUID:      model.UUID,
		Sender:    model.Sender,
		Receiver:  model.Receiver,
		Data:      ToPaymentDataDomain(model.Data),
		URL:       model.URL,
		Status:    model.Status,
		EndedAt:   model.EndedAt,
		CreatedAt: model.CreatedAt,
	}
}

func ToPaymentDataDomain(model PaymentData) domain.PaymentData {
	return domain.PaymentData{}
}
