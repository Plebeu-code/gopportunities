package schemas

import (
	"gorm.io/gorm"
)

// Opening representa uma vaga de emprego.
type Opening struct {
	gorm.Model
	Role     string `gorm:"not null"` // Adicionado not null como exemplo
	Company  string `gorm:"not null"` // Adicionado not null como exemplo
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
