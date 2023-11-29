package helper

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

func DecodeJSON(e echo.Context, input interface{}) error {

	decoder := json.NewDecoder(e.Request().Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(input); err != nil {
		return errors.New("input salah, periksa kembali")
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}

	return nil
}

func BindFormData(c echo.Context, input interface{}) error {

	if err := c.Bind(input); err != nil {
		return err
	}
	// if err := c.Bind(input); err != nil {
	// 	return err
	// }

	decoder := schema.NewDecoder()
	if err := decoder.Decode(input, c.Request().Form); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}

	return nil
}

func HttpResponseCondition(err error, Messages ...string) bool {
	for _, Message := range Messages {
		if strings.Contains(err.Error(), Message) {
			return true
		}
	}
	return false
}

func FieldsEqual(a, b interface{}, fields ...string) bool {
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)

	for _, fieldName := range fields {
		fieldA := valueA.FieldByName(fieldName)
		fieldB := valueB.FieldByName(fieldName)

		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	return true
}
