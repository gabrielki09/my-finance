package categoryrepository

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	constantsdbcode "finance/internal/constants/db"
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
				id::text,
				name,
				created_at,
				updated_at,
				deleted_at
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
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
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

func (c *CategoryRepository) Create(ctx context.Context, model categorymodel.CategoryModel) (categorymodel.CategoryModel, error) {
	var pgErr *pgconn.PgError
	var category categorymodel.CategoryModel

	err := c.db.QueryRow(
		ctx,
		`
			INSERT INTO categories 
				(name)
			VALUES
				($1)
			RETURNING
				id::text,
				name,
				created_at,
				updated_at,
				deleted_at
		`,
		model.Name,
	).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	)

	if err != nil {

		logger.General.Error.Println("Erro ao criar a categoria:", err)
		if errors.As(err, &pgErr) {

			if pgErr.Code == constantsdbcode.UniqueViolationCode {
				logger.General.Error.Println(pgErr)
				logger.General.Error.Println(apperrors.ErrUniqueConstraint.Error())

				return categorymodel.CategoryModel{}, apperrors.ErrUniqueConstraint
			}

		}

		return categorymodel.CategoryModel{}, err
	}

	return category, nil
}

func (c *CategoryRepository) FindById(ctx context.Context, id string) (categorymodel.CategoryModel, error) {
	var category categorymodel.CategoryModel

	err := c.db.QueryRow(
		ctx,
		`
			SELECT 
				id::text,
				name,
				created_at,
				updated_at,
				deleted_at
			FROM
				categories
			WHERE
		 		id = $1 AND
				deleted_at IS NULL
		`,
		id,
	).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
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

func (c *CategoryRepository) Update(ctx context.Context, category categorymodel.CategoryModel) (categorymodel.CategoryModel, error) {
	if err := c.db.QueryRow(
		ctx,
		`
			UPDATE
				categories
			SET
				name = $2,
				updated_at = now()
			WHERE
				id = $1
			RETURNING
				id::text,
				name,
				created_at,
				updated_at,
				deleted_at
		`,
		category.Id,
		category.Name,
	).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt,
	); err != nil {
		logger.General.Error.Println("Erro ao alterar os dados da categoria:", err)
		return categorymodel.CategoryModel{}, err
	}

	return category, nil
}

func (c *CategoryRepository) Delete(ctx context.Context, id string) error {
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
		id,
	); err != nil {
		logger.General.Error.Println("Erro ao deletar os a categoria:", err)
		return err
	}

	return nil
}

func (c *CategoryRepository) Active(ctx context.Context, id string) error {
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
		id,
	); err != nil {
		logger.General.Error.Println("Erro ao ativar a categoria:", err)
		return err
	}

	return nil
}
