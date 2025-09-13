package entity

import (
	"database/sql/driver"
)

type ApprovalStatus string

const (
	Pending  ApprovalStatus = "pending"
	Approved ApprovalStatus = "approved"
	Rejected ApprovalStatus = "rejected"
)

func (a ApprovalStatus) Value() (driver.Value, error) {
	return string(a), nil
}

func (a *ApprovalStatus) Scan(value interface{}) error {
	*a = ApprovalStatus(value.(string))
	return nil
}

type Approval struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title" gorm:"not null"`
	Description *string        `json:"description"`
	Status      ApprovalStatus `json:"status" gorm:"type:approval_status;default:pending;not null"`
}
