package hw09structvalidator_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	hw09structvalidator "github.com/khaydarov/otus-golang-professional/hw09_struct_validator"
	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	Request struct {
		IDs []int `validate:"in:1,2,3"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "123",
				Name:   "Vasya",
				Age:    25,
				Email:  "vasya@gmail.com",
				Role:   "admin",
				Phones: []string{"+79111234567"},
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "1.10.0",
			},
			expectedErr: hw09structvalidator.ValidationErrors{
				hw09structvalidator.ValidationError{
					Field: "Version",
					Err:   errors.New("length is greater than 5"),
				},
			},
		},
		{
			in: Token{
				Header:    []byte("header"),
				Payload:   []byte("payload"),
				Signature: []byte("signature"),
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				Body: "OK",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 600,
				Body: "OK",
			},
			expectedErr: hw09structvalidator.ValidationErrors{
				hw09structvalidator.ValidationError{
					Field: "Code",
					Err:   errors.New("value is not in [200 404 500]"),
				},
			},
		},
		{
			in: Request{
				IDs: []int{1, 2, 5},
			},
			expectedErr: hw09structvalidator.ValidationErrors{
				hw09structvalidator.ValidationError{
					Field: "IDs",
					Err:   errors.New("slice value 5 is not in [1 2 3]"),
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := hw09structvalidator.Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)

			_ = tt
		})
	}
}
