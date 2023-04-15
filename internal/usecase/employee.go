// INFO: Errors do not wrapped, there is no logic, nothing will happen.
// INFO: Pointers are not used for the same reason.

package usecase

import (
	"context"

	"github.com/v1adhope/web-service-employees/internal/entity"
)

type UseCase struct {
	Repo Repo
}

func New(r Repo) *UseCase {
	return &UseCase{r}
}

func (uc *UseCase) Create(ctx context.Context, emp entity.Employee) (string, error) {
	id, err := uc.Repo.Insert(ctx, emp)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (uc *UseCase) DeleteByID(ctx context.Context, id string) error {
	if err := uc.Repo.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetByCompany(ctx context.Context, companyID int) ([]entity.Employee, error) {
	emp, err := uc.Repo.GetByCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (uc *UseCase) GetByDepartament(ctx context.Context, deportment string) ([]entity.Employee, error) {
	emp, err := uc.Repo.GetByDepartament(ctx, deportment)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (uc *UseCase) UpdateByID(ctx context.Context, emp entity.Employee) error {
	if err := uc.Repo.UpdateByID(ctx, emp); err != nil {
		return err
	}

	return nil
}
