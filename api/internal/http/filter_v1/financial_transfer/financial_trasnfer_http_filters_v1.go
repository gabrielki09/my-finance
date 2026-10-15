package financialtransferfiltersv1

import (
	"finance/internal/apperrors"

	"finance/internal/logger"

	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
)

func ParseFinancialTransferFilters(r *http.Request) (string, []any, error) {
	logger.Info("Vai conferir os filtros inseridos na rota")

	errors := apperrors.ValidationErrors{}

	q := r.URL.Query()

	qb := squirrel.
		Select(
			"id",
			"transfer_date",
			"source_account_id",
			"destination_account_id",
			"idempotency_key",
			"created_at",
			"canceled_at",
		).
		From("financial_transfer").
		PlaceholderFormat(squirrel.Dollar)

	if sourceAccountID := q.Get("source_account_id"); sourceAccountID != "" {
		id, err := strconv.Atoi(sourceAccountID)

		if err != nil {
			logger.Error("Erro ao converter o ID da conta de origem:", err)

			errors["source_account_id"] = append(errors["source_account_id"], "ID fora do padrão esperado.")
		} else {
			qb = qb.Where(squirrel.Eq{"source_account_id": id})
		}
	}

	if destinationAccountID := q.Get("destination_account_id"); destinationAccountID != "" {
		id, err := strconv.Atoi(destinationAccountID)

		if err != nil {
			logger.Error("Erro ao converter o ID da conta de destino:", err)

			errors["destination_account_id"] = append(errors["destination_account_id"], "ID fora do padrão esperado.")
		} else {
			qb = qb.Where(squirrel.Eq{"destination_account_id": id})
		}
	}

	if len(errors) > 0 {
		return "", nil, apperrors.NewValidationError(errors)
	}

	query, args, err := qb.ToSql()

	if err != nil {
		logger.Error("Erro ao converter o builder para SQL:", err)
		return "", nil, err
	}

	logger.Info("SQL gerado pelo http_filter %s, filtros: %s", query, args)

	return query, args, nil
}
