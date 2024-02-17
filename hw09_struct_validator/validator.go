package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	Len    = "len"
	Min    = "min"
	Max    = "max"
	In     = "in"
	Regexp = "regexp"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return fmt.Sprint([]ValidationError(v))
}

func parseTag(tag string) map[string]string {
	rules := make(map[string]string)
	for _, rule := range strings.Split(tag, "|") {
		condition := strings.Split(rule, ":")
		rules[condition[0]] = condition[1]
	}
	return rules
}

func Validate(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	var validationErrors ValidationErrors
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := value.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}

		if err := validateField(field, tag); err != nil {
			validationErrors = append(validationErrors, ValidationError{
				Field: value.Type().Field(i).Name,
				Err:   err,
			})
		}
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return validationErrors
}

func validateField(field reflect.Value, tag string) error {
	rulesMap := parseTag(tag)

	//nolint:exhaustive
	switch field.Kind() {
	case reflect.Int:
		intValue := int(field.Int())
		return validateInt(intValue, rulesMap)
	case reflect.String:
		stringValue := field.String()
		return validateString(stringValue, rulesMap)
	case reflect.Slice:
		slice := make([]interface{}, field.Len())
		for i := 0; i < field.Len(); i++ {
			slice[i] = field.Index(i).Interface()
		}
		return validateSlice(slice, rulesMap)
	default:
		return errors.New("unsupported type")
	}
}

func validateInt(fieldValue int, rulesMap map[string]string) error {
	for ruleType, ruleValue := range rulesMap {
		switch ruleType {
		case Min:
			minValue, err := strconv.Atoi(ruleValue)
			if err != nil {
				return err
			}
			if fieldValue < minValue {
				return fmt.Errorf("value is less than %d", minValue)
			}
		case Max:
			maxValue, err := strconv.Atoi(ruleValue)
			if err != nil {
				return err
			}
			if fieldValue > maxValue {
				return fmt.Errorf("value is greater than %d", maxValue)
			}
		case In:
			list := strings.Split(ruleValue, ",")
			if !slices.Contains(list, strconv.Itoa(fieldValue)) {
				return fmt.Errorf("value %d is not in %v", fieldValue, list)
			}
		}
	}

	return nil
}

func validateString(fieldValue string, rulesMap map[string]string) error {
	for ruleType, ruleValue := range rulesMap {
		switch ruleType {
		case Len:
			length, err := strconv.Atoi(ruleValue)
			if err != nil {
				return err
			}
			if len(fieldValue) > length {
				return fmt.Errorf("length is greater than %d", length)
			}
		case In:
			list := strings.Split(ruleValue, ",")
			if !slices.Contains(list, fieldValue) {
				return fmt.Errorf("value %s is not in %v", fieldValue, list)
			}
		case Regexp:
			re, err := regexp.Compile(ruleValue)
			if err != nil {
				return err
			}
			match := re.MatchString(fieldValue)
			if !match {
				return fmt.Errorf("value does not match %s", ruleValue)
			}
		}
	}
	return nil
}

func validateSlice(slice []interface{}, rulesMap map[string]string) error {
	for ruleType, ruleValue := range rulesMap {
		switch ruleType {
		case Len:
			length, err := strconv.Atoi(ruleValue)
			if err != nil {
				return err
			}

			if len(slice) > length {
				return fmt.Errorf("slice length is greater than %d", length)
			}
		default:
			for _, sliceValue := range slice {
				switch v := sliceValue.(type) {
				case string:
					if err := validateString(v, rulesMap); err != nil {
						return err
					}
				case int:
					if err := validateInt(v, rulesMap); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
