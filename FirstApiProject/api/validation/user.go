package validtion

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {
	value , ok := fld.Field().Interface().(string)
	
	if !ok {
		return false
	}
	result , err := regexp.MatchString("^",value)
	if err != nil {
		log.Print(err.Error())		
	}
	return result
}