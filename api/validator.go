package api

import (
	"github.com/Sandhya-Pratama/simple-bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is valid or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
