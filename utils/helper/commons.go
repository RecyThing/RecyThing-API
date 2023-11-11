package helper

import (
	"encoding/json"
	"errors"
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
		return errors.New("input salah, periksa kembali")
	}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(input, c.Request().Form); err != nil {
		return errors.New("input salah, periksa kembali")
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}

	return nil
}
