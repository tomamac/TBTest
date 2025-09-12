package approval

import "github.com/tomamac/tbtest/backend/internal/entity"

type IService interface {
	Get() ([]entity.Approval, error)
	Update(approval UpdateApprovalRequest) error
}

type service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &service{repo: repo}
}

func (s *service) Get() ([]entity.Approval, error) {
	return s.repo.Get()
}

func (s *service) Update(approval UpdateApprovalRequest) error {
	return s.repo.Update(approval)
}
