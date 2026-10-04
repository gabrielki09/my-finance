package financialaccountrepository

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	constantsdbcode "finance/internal/constants/db"
	financialaccountrequest "finance/internal/http/request/financial/financial_account"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	financialmodel "finance/models/financial"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FinancialAccountRepository struct {
	db *pgxpool.Pool
}

func NewFinancialAccountRepository(db *pgxpool.Pool) *FinancialAccountRepository {
	return &FinancialAccountRepository{
		db: db,
	}
}

func (f *FinancialAccountRepository) GetAll(ctx context.Context) ([]financialmodel.FinancialAccountModel, error) {
	var financialAccounts []financialmodel.FinancialAccountModel

	financialAccountsRows, err := f.db.Query(
		ctx,
		`
			SELECT 
				id,
				name,
				type,
				initial_balance,
				opened_at,
				created_at,
    			updated_at
			FROM
				financial_accounts
			WHERE	
				deleted_at IS NULL
		`,
	)

	if err != nil {
		logger.Error("Erro ao exceutar o select:", err)
		return []financialmodel.FinancialAccountModel{}, err
	}

	defer financialAccountsRows.Close()

	for financialAccountsRows.Next() {
		var financialAccount financialmodel.FinancialAccountModel

		if err := financialAccountsRows.Scan(
			&financialAccount.Id,
			&financialAccount.Name,
			&financialAccount.Type,
			&financialAccount.InitialBalance,
			&financialAccount.OpenedAt,
			&financialAccount.CreatedAt,
			&financialAccount.UpdatedAt,
		); err != nil {
			logger.Error("Erro ao ler os dados do select:", err)
			return []financialmodel.FinancialAccountModel{}, err
		}

		financialAccounts = append(financialAccounts, financialAccount)

	}

	if err := financialAccountsRows.Err(); err != nil {
		logger.Error("Erro ao iterar os dados do select:", err)
		return []financialmodel.FinancialAccountModel{}, err
	}

	return financialAccounts, nil
}

func (f *FinancialAccountRepository) Create(ctx context.Context, model financialaccountrequest.FinancialAccountRequest) (financialmodel.FinancialAccountModel, error) {
	var pgErr *pgconn.PgError
	var financialAccount financialmodel.FinancialAccountModel

	err := f.db.QueryRow(
		ctx,
		`
			INSERT INTO financial_accounts 
				(name, type, initial_balance, opened_at)
			VALUES
				($1, $2, $3, $4)
			RETURNING
				id,
				name,
				type,
				initial_balance,
				opened_at,
				created_at,
				updated_at
		`,
		model.Name,
		model.Type,
		model.InitialBalance,
		model.OpenedAt,
	).Scan(
		&financialAccount.Id,
		&financialAccount.Name,
		&financialAccount.Type,
		&financialAccount.InitialBalance,
		&financialAccount.OpenedAt,
		&financialAccount.CreatedAt,
		&financialAccount.UpdatedAt,
	)

	if err != nil {

		logger.Error("Erro ao criar a conta financeira:", err)
		if errors.As(err, &pgErr) {

			if pgErr.Code == constantsdbcode.UniqueViolationCode {
				logger.Error("erro do banco de dados", pgErr)
				logger.Error("erro do banco de dados", apperrors.ErrUniqueConstraint)

				return financialmodel.FinancialAccountModel{}, apperrors.ErrUniqueConstraint
			}

		}

		return financialmodel.FinancialAccountModel{}, err
	}

	return financialAccount, nil
}

func (f *FinancialAccountRepository) FindById(ctx context.Context, financialAccountId int) (financialmodel.FinancialAccountModel, error) {
	var financialAccount financialmodel.FinancialAccountModel

	err := f.db.QueryRow(
		ctx,
		`
			SELECT 
				id,
				name,
				type,
				initial_balance,
				opened_at,
				created_at
			FROM
				financial_accounts
			WHERE
		 		id = $1
		`,
		financialAccountId,
	).Scan(
		&financialAccount.Id,
		&financialAccount.Name,
		&financialAccount.Type,
		&financialAccount.InitialBalance,
		&financialAccount.OpenedAt,
		&financialAccount.CreatedAt,
	)

	if err != nil {
		logger.Error("Erro ao ler os dados da consulta:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.Error("O registro não foi localizado:", err)
			return financialmodel.FinancialAccountModel{}, apperrors.ErrNotFound
		}

		return financialmodel.FinancialAccountModel{}, err
	}

	return financialAccount, nil
}

func (f *FinancialAccountRepository) Update(ctx context.Context, model financialaccountrequest.FinancialAccountRequest, financialAccountId int) (financialAccount financialmodel.FinancialAccountModel, err error) {
	if err = f.db.QueryRow(
		ctx,
		`
			UPDATE
				financial_accounts
			SET
				name = $2,
				type = $3,
				initial_balance = $4,
				opened_at = $5,
				updated_at = now()
			WHERE
				id = $1
			RETURNING
				id,
				name,
				type,
				initial_balance,
				opened_at,
				created_at,
				updated_at
		`,
		financialAccountId,
		model.Name,
		model.Type,
		model.InitialBalance,
		model.OpenedAt,
	).Scan(
		&financialAccount.Id,
		&financialAccount.Name,
		&financialAccount.Type,
		&financialAccount.InitialBalance,
		&financialAccount.OpenedAt,
		&financialAccount.CreatedAt,
		&financialAccount.UpdatedAt,
	); err != nil {
		logger.Error("Erro ao alterar os dados da conta bancaria:", err)
		return financialAccount, err
	}

	return financialAccount, err
}

func (f *FinancialAccountRepository) Delete(ctx context.Context, financialAccountId int) error {
	if _, err := f.db.Exec(
		ctx,
		`
			UPDATE
				financial_accounts
			SET
				updated_at = now(),
				deleted_at = now()
			WHERE
				id = $1	
		`,
		financialAccountId,
	); err != nil {
		logger.Error("Erro ao deletar os a conta financeira:", err)
		return err
	}

	return nil
}

func (f *FinancialAccountRepository) Active(ctx context.Context, financialAccountId int) error {
	if _, err := f.db.Exec(
		ctx,
		`
			UPDATE
				financial_accounts
			SET
				updated_at = now(),
				deleted_at = NULL
			WHERE
				id = $1
		`,
		financialAccountId,
	); err != nil {
		logger.Error("Erro ao ativar a conta financeira:", err)
		return err
	}

	return nil
}

func (f *FinancialAccountRepository) VerifyExistsFinancialAccountName(ctx context.Context, financialAccountName string) (*financialmodel.FinancialAccountModel, error) {
	var financialAccount financialmodel.FinancialAccountModel

	err := f.db.QueryRow(
		ctx,
		`
			SELECT 
				id,
				name
			FROM
				financial_accounts
			WHERE
		 		name = $1
		`,
		financialAccountName,
	).Scan(
		&financialAccount.Id,
		&financialAccount.Name,
	)

	if err != nil {
		logger.Error("Erro ao ler os dados da consulta:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.Info("A conta financeira não existe")
			return nil, nil
		}

		return nil, err
	}

	return &financialAccount, nil
}

func (f *FinancialAccountRepository) GetCurrentBalance(ctx context.Context, financialAccountId int) (financialresponse.FinancialCurrentBalanceResponse, error) {
	var currentBalanceBody financialresponse.FinancialCurrentBalanceResponse

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT
				fa.id,
				fa.name,
				fa.initial_balance 
				+
				COALESCE(SUM(
					CASE
						WHEN ft.movement_type = 'entry' THEN ft.amount
						ELSE 0
					END
				), 0)
				-
				COALESCE(SUM(
						CASE
						WHEN ft.movement_type = 'exit' THEN ft.amount
						ELSE 0
					END
				), 0) AS balance
			FROM		
				financial_accounts fa
			LEFT JOIN financial_transactions ft 
				ON ft.financial_account_id = fa.id
				AND ft.canceled_at IS NULL
				AND ft.operation_type IN ('original', 'adjustment')
			WHERE
				fa.deleted_at IS NULL
				AND fa.id = $1
			GROUP BY
				fa.id,
				fa.name,
				fa.initial_balance 
		`,
		financialAccountId,
	).Scan(
		&currentBalanceBody.Id,
		&currentBalanceBody.Name,
		&currentBalanceBody.Balance,
	); err != nil {
		logger.Error("Erro ao consultar o balanço total da conta financeira:", err)

		return financialresponse.FinancialCurrentBalanceResponse{}, err
	}

	return currentBalanceBody, nil
}
