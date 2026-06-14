package categoryrepository

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	constantsdbcode "finance/internal/constants/db"
	transactionhelper "finance/internal/helpers/transaction"
	categoryrequest "finance/internal/http/request/category"
	"finance/internal/logger"
	categorymodel "finance/models/category"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (c *CategoryRepository) GetAll(ctx context.Context) ([]categorymodel.CategoryModel, error) {
	var categories []categorymodel.CategoryModel

	categoriesRows, err := c.db.Query(
		ctx,
		`
			SELECT 
				id,
				parent_id,
				name,
				type,
				created_at,
				updated_at
			FROM
				categories
			WHERE	
				deleted_at IS NULL
		`,
	)

	if err != nil {
		logger.General.Error.Println("Erro ao exceutar o select:", err)
		return []categorymodel.CategoryModel{}, err
	}

	defer categoriesRows.Close()

	for categoriesRows.Next() {
		var category categorymodel.CategoryModel

		if err := categoriesRows.Scan(
			&category.Id,
			&category.ParentId,
			&category.Name,
			&category.Type,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			logger.General.Error.Println("Erro ao ler os dados do select:", err)
			return []categorymodel.CategoryModel{}, err
		}

		categories = append(categories, category)

	}

	if err := categoriesRows.Err(); err != nil {
		logger.General.Error.Println("Erro ao iterar os dados do select:", err)
		return []categorymodel.CategoryModel{}, err
	}

	return categories, nil
}

func (c *CategoryRepository) Create(ctx context.Context, model categoryrequest.CategoryRequest) (category categorymodel.CategoryModel, err error) {
	var pgErr *pgconn.PgError

	err = c.db.QueryRow(
		ctx,
		`
			INSERT INTO categories 
				(parent_id, name, type)
			VALUES
				($1, $2, $3)
			RETURNING
				id,
				parent_id,
				name,
				type,
				created_at,
				updated_at
		`,
		model.ParentId,
		model.Name,
		model.Type,
	).Scan(
		&category.Id,
		&category.ParentId,
		&category.Name,
		&category.Type,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {

		logger.General.Error.Println("Erro ao criar a categoria:", err)
		if errors.As(err, &pgErr) {

			if pgErr.Code == constantsdbcode.UniqueViolationCode {
				logger.General.Error.Println(pgErr)
				logger.General.Error.Println(apperrors.ErrUniqueConstraint.Error())

				return category, apperrors.ErrUniqueConstraint
			}

		}

		return category, err
	}

	return category, err
}

func (c *CategoryRepository) FindById(ctx context.Context, categoryId int) (categorymodel.CategoryModel, error) {
	var category categorymodel.CategoryModel

	err := c.db.QueryRow(
		ctx,
		`
			SELECT 
				id,
				parent_id,
				name,
				type,
				created_at,
				updated_at
			FROM
				categories
			WHERE
		 		id = $1 AND
				deleted_at IS NULL
		`,
		categoryId,
	).Scan(
		&category.Id,
		&category.ParentId,
		&category.Name,
		&category.Type,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		logger.General.Error.Println("Erro ao ler os dados da consulta:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.General.Error.Println("O registro não foi localizado:", err)
			return categorymodel.CategoryModel{}, apperrors.ErrNotFound
		}

		return categorymodel.CategoryModel{}, err
	}

	return category, nil
}

func (c *CategoryRepository) Update(ctx context.Context, payload categoryrequest.CategoryRequest, categoryId int) (category categorymodel.CategoryModel, err error) {
	err = transactionhelper.WithTransaction(
		ctx,
		c.db,
		func(tx pgx.Tx) error {
			if err := c.db.QueryRow(
				ctx,
				`
					UPDATE
						categories
					SET
						parent_id = $2,
						name = $3,
						type = $4,
						updated_at = now()
					WHERE
						id = $1
					RETURNING
						id,
						parent_id,
						name,
						type,
						created_at,
						updated_at
				`,
				categoryId,
				payload.ParentId,
				payload.Name,
				payload.Type,
			).Scan(
				&category.Id,
				&category.ParentId,
				&category.Name,
				&category.Type,
				&category.CreatedAt,
				&category.UpdatedAt,
			); err != nil {
				logger.General.Error.Println("Erro ao alterar os dados da categoria:", err)
				return err
			}

			return nil
		})

	return category, err
}

func (c *CategoryRepository) Delete(ctx context.Context, categoryId int) error {
	if _, err := c.db.Exec(
		ctx,
		`
			UPDATE
				categories
			SET
				updated_at = now(),
				deleted_at = now()
			WHERE
				id = $1	
		`,
		categoryId,
	); err != nil {
		logger.General.Error.Println("Erro ao deletar os a categoria:", err)
		return err
	}

	return nil
}

func (c *CategoryRepository) Active(ctx context.Context, categoryId int) error {
	if _, err := c.db.Exec(
		ctx,
		`
			UPDATE
				categories
			SET
				updated_at = now(),
				deleted_at = NULL
			WHERE
				id = $1
		`,
		categoryId,
	); err != nil {
		logger.General.Error.Println("Erro ao ativar a categoria:", err)
		return err
	}

	return nil
}

func (c *CategoryRepository) VerifyParentId(ctx context.Context, parentId int) (bool, error) {
	logger.General.Info.Println("CategoryRepository - VerifyParentId called")

	var exists bool

	if err := c.db.QueryRow(
		ctx,
		`SELECT EXISTS 
			( 
			SELECT
				id 
			FROM
				categories 
			WHERE 
				id = $1
			ORDER BY 
				id
			DESC 
				LIMIT 1
			)
		`,
		parentId,
	).Scan(&exists); err != nil {
		logger.General.Error.Println("Erro ao conferir se a categoria pai existe:", err)
		return false, err
	}

	if exists {
		logger.General.Info.Println("A categoria pai existe")
		return true, nil
	}

	logger.General.Info.Println("A categoria pai não existe")
	return false, nil
}

func (c *CategoryRepository) VerifyExistsCategoryName(ctx context.Context, categoryName string) (*categorymodel.CategoryModel, error) {
	var category categorymodel.CategoryModel

	if err := c.db.QueryRow(
		ctx,
		`
			SELECT
				id,
				name
			FROM
				categories
			WHERE
				name = $1
		`,
		categoryName,
	).Scan(
		&category.Id,
		&category.Name,
	); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			logger.General.Info.Println("A categoria não existe")
			return nil, nil

		}

		return nil, err
	}

	return &category, nil
}
