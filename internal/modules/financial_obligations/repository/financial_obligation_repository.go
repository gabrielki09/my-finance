package financialobligationrepository

import (
	"context"
	financialmodel "finance/models/financial"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FinancialObligationRepository struct {
	db *pgxpool.Pool
}

func NewFinancialObligationRepository(db *pgxpool.Pool) *FinancialObligationRepository {
	return &FinancialObligationRepository{
		db: db,
	}
}

func (f FinancialObligationRepository) GetAll(ctx context.Context) ([]financialmodel.FinancialObligationModel, error) {
	var financialObligation []financialmodel.FinancialObligationModel

	return financialObligation, nil
}
