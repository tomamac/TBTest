package approval

import (
	"github.com/tomamac/tbtest/backend/internal/entity"
	"gorm.io/gorm"
)

type IRepository interface {
	Get() ([]entity.Approval, error)
	Update(approval UpdateApprovalRequest) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{db: db}
}

func (r *repository) Get() ([]entity.Approval, error) {
	var approvals []entity.Approval

	res := r.db.Order("id").Find(&approvals)

	if res.Error != nil {
		return nil, res.Error
	}

	return approvals, nil
}

func (r *repository) Update(approvals UpdateApprovalRequest) error {
	res := r.db.Model(entity.Approval{}).Where("id IN ?", approvals.IDs).
		Updates(entity.Approval{Description: &approvals.Description, Status: approvals.Status})

	if res.Error != nil {
		return res.Error
	}

	return nil
}
