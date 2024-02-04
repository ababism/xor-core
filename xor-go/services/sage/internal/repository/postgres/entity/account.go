package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type AccountPostgres struct {
	Uuid         uuid.UUID         `xdb:"uuid"`
	Login        string            `xdb:"login"`
	PasswordHash string            `xdb:"password_hash"`
	Contacts     *ContactsPostgres `xdb:"contacts"`
	Active       bool              `xdb:"active"`
}

type ContactsPostgres struct {
	TgLogin string `json:"tg_login,omitempty"`
}

func (r *ContactsPostgres) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ContactsPostgres) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &r)
}
