package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null" json:"name"`

	Endpoints []Endpoint `gorm:"foreignKey:ProjectID" json:"endpoints"`
	Workflows []Workflow `gorm:"foreignKey:ProjectID" json:"workflows"`
	Users     []User     `gorm:"many2many:user_projects;" json:"users,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	IsActive bool `gorm:"-"`
}

func (Project) TableName() string {
	return "projects"
}

func (p *Project) AfterFind(tx *gorm.DB) error {
	if activeProjectID, ok := tx.Statement.Context.Value("activeProjectID").(*uuid.UUID); ok {
		if activeProjectID != nil && *activeProjectID == p.ID {
			p.IsActive = true
		}
	}
	return nil
}