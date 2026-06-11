package apperrors

import "errors"

var (
	ErrNotFound            = errors.New("registro não encontrado.")
	ErrInvalidUUID         = errors.New("o id informado está fora do formato esperado.")
	ErrUnprocessableEntity = errors.New("campos obrigatórios ausentes ou inválidos.")
	// DB errors
	ErrUniqueConstraint = errors.New("valores duplicados.")
	ErrCheckViolation   = errors.New("valores inseridos fora do padrão.")
)

type ValidationErrors map[string][]string

type ValidationError struct {
	Errors ValidationErrors
}

func NewValidationError(errors ValidationErrors) *ValidationError {
	return &ValidationError{
		Errors: errors,
	}
}
func (e *ValidationError) Error() string {
	return "erro de validação"
}
