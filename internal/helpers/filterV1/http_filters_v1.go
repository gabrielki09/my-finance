package filterv1

import (
	"finance/internal/logger"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
	_ "github.com/Masterminds/squirrel"
)

type FinancialObligationFilters struct {
	CategoryId   *int
	Type         *string
	Status       *string
	StartDueDate *string
	EndDueDate   *string
}

func ParseFinancialObligationFilters(r *http.Request) (string, error) {
	q := r.URL.Query()

	qb := squirrel.
		Select(
			"id",
			"category_id",
			"description",
			"type",
			"status",
			"original_amount",
			"due_date",
			"competence_date",
			"notes",
			"created_at",
			"updated_at",
		).
		From("financial_obligations").
		Where("deleted_at IS NULL")

	if c := q.Get("category_id"); c != "" {
		categoryId, err := strconv.Atoi(c)

		if err != nil {
			logger.General.Error.Println("Erro ao converter o ID da categoria:", err)
		}

		//filter.CategoryId = &categoryId
		qb = qb.Where(squirrel.Eq{"category_id": categoryId})
	}

	if t := q.Get("type"); t != "" {
		qb = qb.Where(squirrel.Eq{"status": t})
	}

	if s := q.Get("status"); s != "" {
		qb = qb.Where(squirrel.Eq{"status": s})
	}

	sd := q.Get("start_due_date")
	ed := q.Get("end_due_date")

	if sd != "" && ed != "" {
		//qb = qb.Where(squirrel.{"status": s})
		qb = qb.Where(squirrel.Expr("due_date BETWEEN $1 AND $2"), sd, ed)
	}

	logger.General.Info.Println("SQL gerado pelo http_filter:", qb)

	query, _, err := qb.ToSql()

	if err != nil {
		logger.General.General.Println("Erro ao converter o builder para SQL:", err)
		return "", err
	}

	return query, nil
}
