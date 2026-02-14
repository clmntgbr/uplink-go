package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Endpoint struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	BaseUri string    `gorm:"not null" json:"baseUri"`
	Path string    `gorm:"not null" json:"path"`
	Method string    `gorm:"not null" json:"method"`
	Timeout int `gorm:"not null" json:"timeout"`

	Header JSON `json:"header" gorm:"type:json"`
	Body   JSON `json:"body" gorm:"type:json"`
	Query  JSON `json:"query" gorm:"type:json"`

	ProjectID uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Endpoint) TableName() string {
	return "endpoints"
}

type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	
	return json.Unmarshal(bytes, j)
}