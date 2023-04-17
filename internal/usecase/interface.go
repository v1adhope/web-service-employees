package usecase

import (
	"context"

	"github.com/v1adhope/web-service-employees/internal/entity"
)

type (
	Repo interface {
		Insert(ctx context.Context, emp entity.Employee) (string, error)
		DeleteByID(ctx context.Context, id string) error
		GetByCompanyID(ctx context.Context, companyID int) ([]entity.Employee, error)
		GetByDeportamentName(ctx context.Context, deportment string) ([]entity.Employee, error)
		UpdateByID(ctx context.Context, emp entity.Employee) error
	}

	Employee interface {
		Create(ctx context.Context, emp entity.Employee) (string, error)
		DeleteByID(ctx context.Context, id string) error
		GetByCompanyID(ctx context.Context, companyID int) ([]entity.Employee, error)
		GetByDeportamentName(ctx context.Context, deportment string) ([]entity.Employee, error)
		UpdateByID(ctx context.Context, emp entity.Employee) error
	}
)
