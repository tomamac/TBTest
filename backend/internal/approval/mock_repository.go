package approval

import "github.com/tomamac/tbtest/backend/internal/entity"

type mockrepository struct {
	approvals []entity.Approval
}

func NewMockRepository() IRepository {
	return &mockrepository{
		approvals: []entity.Approval{
			{ID: 1, Title: "doc1", Status: entity.Pending},
			{ID: 2, Title: "doc2", Status: entity.Pending},
			{ID: 3, Title: "doc3", Status: entity.Pending},
			{ID: 4, Title: "doc4", Status: entity.Pending},
			{ID: 5, Title: "doc5", Status: entity.Pending},
			{ID: 6, Title: "doc6", Status: entity.Pending},
			{ID: 7, Title: "doc7", Status: entity.Pending},
			{ID: 8, Title: "doc8", Status: entity.Pending},
			{ID: 9, Title: "doc9", Status: entity.Pending},
			{ID: 10, Title: "doc10", Status: entity.Pending},
		},
	}
}

func (r *mockrepository) Get() ([]entity.Approval, error) {
	return r.approvals, nil
}

func (r *mockrepository) Update(approvals UpdateApprovalRequest) error {
	for i := range approvals.IDs {
		r.approvals[i].Description = &approvals.Description
		r.approvals[i].Status = approvals.Status
	}

	return nil
}
