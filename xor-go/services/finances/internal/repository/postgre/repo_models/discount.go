package repo_models

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type Discount struct {
	UUID          uuid.UUID `db:"uuid"`
	CreatedBy     uuid.UUID `db:"created_by"`
	Percent       float32   `db:"percent"`
	StandAlone    bool      `db:"stand_alone"`
	StartedAt     time.Time `db:"started_at"`
	EndedAt       time.Time `db:"ended_at"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

func CreateToDiscountPostgres(model *domain.DiscountCreate) *Discount {
	id, _ := uuid.NewUUID()
	return &Discount{
		UUID:       id,
		CreatedBy:  model.CreatedBy,
		Percent:    model.Percent,
		StandAlone: model.StandAlone,
		StartedAt:  model.StartedAt,
		EndedAt:    model.EndedAt,
		Status:     model.Status,
		CreatedAt:  time.Now(),
	}
}

func UpdateToDiscountPostgres(model *domain.DiscountUpdate) *Discount {
	return &Discount{
		UUID:       model.UUID,
		CreatedBy:  model.CreatedBy,
		Percent:    model.Percent,
		StandAlone: model.StandAlone,
		StartedAt:  model.StartedAt,
		EndedAt:    model.EndedAt,
		Status:     model.Status,
	}
}

func ToDiscountDomain(model *Discount) *domain.DiscountGet {
	return &domain.DiscountGet{
		UUID:         model.UUID,
		CreatedBy:    model.CreatedBy,
		Percent:      model.Percent,
		StandAlone:   model.StandAlone,
		StartedAt:    model.StartedAt,
		EndedAt:      model.EndedAt,
		Status:       model.Status,
		CreatedAt:    model.CreatedAt,
		LastUpdateAt: model.LastUpdatedAt,
	}
}
