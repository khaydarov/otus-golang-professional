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
	//nolint:exhaustive
	switch field.Kind() {
	case reflect.Int:
		return validateInt(field, tag)
	case reflect.String:
		return validateString(field, tag)
	case reflect.Slice:
		intSlice, ok := field.Interface().([]int)
		if ok {
			return validateIntSlice(intSlice, tag)
		}

		stringSlice, ok := field.Interface().([]string)
		if ok {
			return validateStringSlice(stringSlice, tag)
		}

		return errors.New("unsupported slice type")
	default:
		return errors.New("unsupported type")
	}
}

func validateInt(field reflect.Value, tag string) error {
	fieldValue := field.Int()
	ruleSet := strings.Split(tag, "|")
	for _, rule := range ruleSet {
		condition := strings.Split(rule, ":")
		switch condition[0] {
		case Min:
			minValue, _ := strconv.Atoi(condition[1])
			if fieldValue < int64(minValue) {
				return fmt.Errorf("value is less than %d", minValue)
			}
		case Max:
			maxValue, _ := strconv.Atoi(condition[1])
			if fieldValue > int64(maxValue) {
				return fmt.Errorf("value is greater than %d", maxValue)
			}
		case In:
			list := strings.Split(condition[1], ",")
			if !slices.Contains(list, strconv.Itoa(int(fieldValue))) {
				return fmt.Errorf("value is not in %v", list)
			}
		}
	}

	return nil
}

func validateString(field reflect.Value, tag string) error {
	fieldValue := field.String()
	ruleSet := strings.Split(tag, "|")
	for _, rule := range ruleSet {
		condition := strings.Split(rule, ":")
		switch condition[0] {
		case Len:
			length, _ := strconv.Atoi(condition[1])
			if len(fieldValue) > length {
				return fmt.Errorf("length is greater than %d", length)
			}
		case In:
			list := strings.Split(condition[1], ",")
			if !slices.Contains(list, fieldValue) {
				return fmt.Errorf("value is not in %v", list)
			}
		case Regexp:
			re, _ := regexp.Compile(condition[1])
			match := re.MatchString(fieldValue)
			if !match {
				return fmt.Errorf("value does not match %s", condition[1])
			}
		}
	}
	return nil
}

func validateIntSlice(intSlice []int, tag string) error {
	ruleSet := strings.Split(tag, "|")
	for _, rule := range ruleSet {
		condition := strings.Split(rule, ":")
		switch condition[0] {
		case Len:
			length, _ := strconv.Atoi(condition[1])
			if len(intSlice) > length {
				return fmt.Errorf("length is greater than %d", length)
			}
		case Min:
			minValue, _ := strconv.Atoi(condition[1])

			for _, v := range intSlice {
				if v < minValue {
					return fmt.Errorf("slice value %d is less than %d", v, minValue)
				}
			}
		case Max:
			maxValue, _ := strconv.Atoi(condition[1])

			for _, v := range intSlice {
				if v > maxValue {
					return fmt.Errorf("slice value %d is greater than %d", v, maxValue)
				}
			}
		case In:
			list := strings.Split(condition[1], ",")
			for _, v := range intSlice {
				if !slices.Contains(list, strconv.Itoa(v)) {
					return fmt.Errorf("slice value %d is not in %v", v, list)
				}
			}
		}
	}

	return nil
}

func validateStringSlice(stringSlice []string, tag string) error {
	ruleSet := strings.Split(tag, "|")
	for _, rule := range ruleSet {
		condition := strings.Split(rule, ":")

		switch condition[0] {
		case Len:
			length, _ := strconv.Atoi(condition[1])
			if len(stringSlice) > length {
				return fmt.Errorf("length is greater than %d", length)
			}
		case In:
			list := strings.Split(condition[1], ",")

			for _, v := range stringSlice {
				if !slices.Contains(list, v) {
					return fmt.Errorf("value is not in %v", list)
				}
			}
		case Regexp:
			re, _ := regexp.Compile(condition[1])

			for _, v := range stringSlice {
				match := re.MatchString(v)
				if !match {
					return fmt.Errorf("value does not match %s", condition[1])
				}
			}
		}
	}

	return nil
}
