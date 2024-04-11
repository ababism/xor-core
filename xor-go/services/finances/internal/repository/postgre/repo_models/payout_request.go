package repo_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type PayoutRequestData struct {
}

func (a *PayoutRequestData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *PayoutRequestData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type PayoutRequest struct {
	UUID      uuid.UUID         `db:"uuid"`
	Receiver  uuid.UUID         `db:"receiver"`
	Status    string            `db:"status"`
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
		Status:    "created",
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
		UUID:      model.UUID,
		Receiver:  model.Receiver,
		Status:    model.Status,
		Amount:    model.Amount,
		CreatedAt: model.CreatedAt,
		Data:      ToPayoutRequestDataDomain(model.Data),
	}
}
