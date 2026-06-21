package filterv1

import (
	"errors"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
)

func ParseFinancialObligationFilters(r *http.Request) (string, []any, error) {
	logger.General.Info.Println("Vai conferir os filtros inseridos na rota")

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
		Where("deleted_at IS NULL").
		PlaceholderFormat(squirrel.Dollar)

	if c := q.Get("category_id"); c != "" {
		categoryId, err := strconv.Atoi(c)

		if err != nil {
			logger.General.Error.Println("Erro ao converter o ID da categoria:", err)

			return "", nil, errors.New("o filtro category_id deve ser um número inteiro.")
		}

		qb = qb.Where(squirrel.Eq{"category_id": categoryId})
	}

	if t := q.Get("type"); t != "" {

		if !financialobligationrequest.ValidateFinancialObligationsTypes(financialmodel.FinancialObligationsTypes(t)) {
			return "", nil, errors.New("o filtro type deve ser estar em um formato válido.")
		}

		qb = qb.Where(squirrel.Eq{`"type"`: t})
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

	query, args, err := qb.ToSql()

	if err != nil {
		logger.General.General.Println("Erro ao converter o builder para SQL:", err)
		return "", nil, err
	}

	logger.General.Info.Println("SQL gerado pelo http_filter:", query)

	return query, args, nil
}
