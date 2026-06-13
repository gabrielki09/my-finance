package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func DecodeErrorMessage(err error) map[string]any {
	var syntaxError *json.SyntaxError
	var typeError *json.UnmarshalTypeError

	switch {
	case errors.Is(err, io.EOF):
		return map[string]any{
			"message": "Payload vazio.",
		}

	case errors.As(err, &syntaxError):
		return map[string]any{
			"message": "JSON inválido",
			"details": fmt.Sprintf("Erro próximo ao byte %d", syntaxError.Offset),
		}

	case errors.As(err, &typeError):
		field := typeError.Field

		if field == "" {
			field = "campo_desconhecido"
		}

		return map[string]any{
			"message": "Tipo de dado inválido",
			"fields": map[string][]string{
				field: {
					fmt.Sprintf(
						"O campo %s deve ser do tipo %s",
						field,
						typeError.Type.String(),
					),
				},
			},
		}
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		field := strings.TrimPrefix(err.Error(), "json: unknown field ")
		field = strings.Trim(field, `"`)

		return map[string]any{
			"message": "Campo desconhecido no JSON.",
			"fields": map[string][]string{
				field: {
					"Este campo não é permetido.",
				},
			},
		}

	default:
		return map[string]any{
			"message": "Erro ao ler o JSON.",
			"details": err.Error(),
		}
	}
}
