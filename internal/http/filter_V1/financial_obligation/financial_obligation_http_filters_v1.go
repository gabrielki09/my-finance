package financialobligationfiltersv1

import (
	"finance/internal/apperrors"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
)

func ParseFinancialObligationFilters(r *http.Request) (string, []any, error) {
	logger.General.Info.Println("Vai conferir os filtros inseridos na rota")

	errors := apperrors.ValidationErrors{}

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

			errors["category_id"] = append(errors["category_id"], "Tipo da categoria incoerente com o tipo da obrigação financeira")
		} else {
			qb = qb.Where(squirrel.Eq{"category_id": categoryId})
		}
	}

	if t := q.Get("type"); t != "" {
		if !financialobligationrequest.ValidateFinancialObligationsTypes(financialmodel.FinancialObligationsTypes(t)) {
			errors["type"] = append(errors["type"], "o filtro type deve estar em um formato válido.")
		} else {
			qb = qb.Where(squirrel.Eq{`"type"`: t})

		}
	}

	if s := q.Get("status"); s != "" {
		validFinancialObligationStatus := func(t financialmodel.FinancialObligationsStatus) bool {
			switch t {
			case financialmodel.PENDING,
				financialmodel.PARTIALLY_SETTLED,
				financialmodel.SETTLED,
				financialmodel.CANCELED:

				return true

			default:
				return false
			}
		}

		if !validFinancialObligationStatus(financialmodel.FinancialObligationsStatus(s)) {
			errors["status"] = append(errors["status"], "o filtro status deve estar em um formato válido.")
		} else {
			qb = qb.Where(squirrel.Eq{"status": s})

		}
	}

	startDueDate := q.Get("start_due_date")
	endDueDate := q.Get("end_due_date")

	logger.General.Info.Printf("startDueDate: %s, endDueDate: %s", startDueDate, endDueDate)

	var haveOneErr bool

	if startDueDate != "" && endDueDate == "" {
		errors["end_due_date"] = append(errors["end_due_date"], "A data final deve ser informada quando a data de inicio for preenchida.")
	} else if startDueDate == "" && endDueDate != "" {
		errors["start_due_date"] = append(errors["start_due_date"], "A data de inicio deve ser informada quando a data final for preenchida.")
	}

	if startDueDate != "" && endDueDate != "" {
		if _, err := time.Parse("2006-01-02", startDueDate); err != nil {
			logger.General.Error.Println("Erro ao convertar a data inicial:", err)
			errors["start_due_date"] = append(errors["start_due_date"], fmt.Sprintln("Erro ao converter a data para YYYY-MM-DD:", err))
			haveOneErr = true

		} else if _, err := time.Parse("2006-01-02", endDueDate); err != nil {
			logger.General.Error.Println("Erro ao convertar a data final:", err)
			errors["end_due_date"] = append(errors["end_due_date"], fmt.Sprintln("Erro ao converter a data para YYYY-MM-DD:", err))
			haveOneErr = true
		}
	}

	if !haveOneErr {
		logger.General.Info.Println("Datas validadas")
		qb = qb.Where(squirrel.Expr("due_date BETWEEN ? AND ?", startDueDate, endDueDate))

	}

	if len(errors) > 0 {
		return "", nil, apperrors.NewValidationError(errors)
	}

	query, args, err := qb.ToSql()

	if err != nil {
		logger.General.Error.Println("Erro ao converter o builder para SQL:", err)
		return "", nil, err
	}

	logger.General.Info.Printf("SQL gerado pelo http_filter %s, filtros: %s", query, args)

	return query, args, nil
}
