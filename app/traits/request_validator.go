package traits

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/goravel/framework/contracts/http"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{Validator: validator.New()}
}

func (rv *RequestValidator) BindAndValidate(ctx http.Context, req interface{}) error {
	method := ctx.Request().Origin().Method

	if method == http.MethodGet || method == http.MethodPut || method == http.MethodDelete {
		if err := ctx.Request().Bind(req); err != nil {
			return fmt.Errorf("invalid query params: %v", err)
		}
	} else {
		if err := json.NewDecoder(ctx.Request().Origin().Body).Decode(req); err != nil {
			if err == io.EOF {
				return fmt.Errorf("empty request body")
			}
			return err
		}
	}

	if err := rv.Validator.Struct(req); err != nil {
		return err
	}

	return nil
}
