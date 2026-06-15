package financialaccountrepository

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	constantsdbcode "finance/internal/constants/db"
	financialaccountrequest "finance/internal/http/request/financial/financial_account"
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
		logger.General.Error.Println("Erro ao exceutar o select:", err)
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
			logger.General.Error.Println("Erro ao ler os dados do select:", err)
			return []financialmodel.FinancialAccountModel{}, err
		}

		financialAccounts = append(financialAccounts, financialAccount)

	}

	if err := financialAccountsRows.Err(); err != nil {
		logger.General.Error.Println("Erro ao iterar os dados do select:", err)
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
				created_at
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

		logger.General.Error.Println("Erro ao criar a conta financeira:", err)
		if errors.As(err, &pgErr) {

			if pgErr.Code == constantsdbcode.UniqueViolationCode {
				logger.General.Error.Println(pgErr)
				logger.General.Error.Println(apperrors.ErrUniqueConstraint.Error())

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
		logger.General.Error.Println("Erro ao ler os dados da consulta:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.General.Error.Println("O registro não foi localizado:", err)
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
		logger.General.Error.Println("Erro ao alterar os dados da conta bancaria:", err)
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
		logger.General.Error.Println("Erro ao deletar os a conta financeira:", err)
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
		logger.General.Error.Println("Erro ao ativar a conta financeira:", err)
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
		financialAccount.Id,
		financialAccount.Name,
	)

	if err != nil {
		logger.General.Error.Println("Erro ao ler os dados da consulta:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.General.Info.Println("A conta financeira não existe")
			return nil, nil
		}

		return nil, err
	}

	return &financialAccount, nil
}
