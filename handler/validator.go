package handler

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	StatusIsRequired    = "is required"
	StatusMinLengthIs   = "min length is"
	StatusMinValueIs    = "min value is"
	StatusMaxLengthIs   = "max length is"
	StatusMaxValueIs    = "max value is"
	StatusInvalidFormat = "format is invalid"
)

func ValidateStruct(s interface{}) []generated.ErrorDetail {
	var errs []generated.ErrorDetail

	t := reflect.TypeOf(s)
	values := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tagValidate := field.Tag.Get("validate")
		if len(tagValidate) > 0 {

			//Get json tag name
			name := field.Tag.Get("json")
			if name == "" {
				name = field.Name
			}

			//Split validation for multiple validation on 1 field
			validations := strings.Split(tagValidate, ",")
			for _, v := range validations {
				err := validateTag(name, v, values.Field(i))

				// only get the first validation founded on each field
				if err != nil {
					//If validation error is exist, then append to return the message
					errs = append(errs, *err)

					break
				}
			}

		}
	}

	return errs
}

func validateTag(fieldName string, tag string, value reflect.Value) *generated.ErrorDetail {
	msg := ""

	// Get element if kind of pointer
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if tag == "required" && (value.IsZero()) {
		// if struct has required tag name and value is empty
		msg = StatusIsRequired
	} else if strings.Contains(tag, "min=") {
		// if struct field has min tag, then validate
		sep := strings.Split(tag, "min=")
		if len(sep) == 2 {
			minValue, _ := strconv.Atoi(sep[1])

			if value.Kind() == reflect.String && len(value.String()) < minValue {
				msg = fmt.Sprintf("%s %d", StatusMinLengthIs, minValue)
			} else if value.Kind() == reflect.Int && int(value.Int()) < minValue {
				msg = fmt.Sprintf("%s %d", StatusMinValueIs, minValue)
			}
		}

	} else if strings.Contains(tag, "max=") {
		// if struct field contain max tag, then validate
		sep := strings.Split(tag, "max=")
		if len(sep) == 2 {
			maxValue, _ := strconv.Atoi(sep[1])

			if value.Kind() == reflect.String && len(value.String()) > maxValue {
				msg = fmt.Sprintf("%s %d", StatusMaxLengthIs, maxValue)
			} else if value.Kind() == reflect.Int && int(value.Int()) > maxValue {
				msg = fmt.Sprintf("%s %d", StatusMaxValueIs, maxValue)
			}
		}

	} else if tag == "phone" && !value.IsZero() && !validPhoneNumber(value.String()) {
		// if struct field has phone tag, and format is not valid
		msg = StatusInvalidFormat
	} else if tag == "secure_password" && !value.IsZero() {
		// check password security
		passCheck, errMsg := isSecurePassword(value.String())

		if !passCheck {
			// if struct field has secure_password tag, and is not secure enough
			msg = errMsg
		}
	}

	if len(msg) > 0 {
		// generate error bad request details
		return &generated.ErrorDetail{
			Title:   fieldName,
			Message: msg,
		}
	}
	return nil

}

func validPhoneNumber(phone string) bool {
	pattern := `^\+62\d{7,10}$`
	return regexp.MustCompile(pattern).MatchString(phone)
}

func isSecurePassword(input string) (bool, string) {
	var msg []string

	// make regex description
	reg := map[string]string{
		`.{6,}`: "min length is 6 char",
		`[A-Z]`: "must contains at least one capital letter",
		`\d`:    "must containts at least one number",
		`\W`:    "must contains at least one special char",
	}

	secure := true
	for m, v := range reg {
		if !regexp.MustCompile(m).MatchString(input) {
			secure = false
			msg = append(msg, v)
		}
	}

	return secure, strings.Join(msg, ", ")
}
