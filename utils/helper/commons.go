package helper

import (
	"encoding/json"
	"errors"
	"recything/utils/constanta"
	"reflect"
	"strings"
	"time"

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
	structA := reflect.ValueOf(a)
	structB := reflect.ValueOf(b)

	for _, fieldName := range fields {
		fieldA := structA.FieldByName(fieldName)
		fieldB := structB.FieldByName(fieldName)

		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	return true
}

func ChangeStatusMission(endDate string) (string, error) {
	var status string
	endDateValid, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return status, err
	}
	currentTime := time.Now().Truncate(24 * time.Hour)
	if endDateValid.Before(currentTime) {
		status = constanta.OVERDUE
	} else {
		status = constanta.ACTIVE
	}
	return status, nil
}
