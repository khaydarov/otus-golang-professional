package tests

import (
	"github.com/brianvoe/gofakeit/v6"
	ssov1 "github.com/khaydarov/otus-golang-professional/sample_projects/protos/gen/go/sso"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/tests/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	pass := randomFakePassword()

	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token)
	require.Equal(t, "token123", token)
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}
