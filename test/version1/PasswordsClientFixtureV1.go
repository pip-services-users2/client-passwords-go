package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-passwords-go/version1"
	"github.com/stretchr/testify/assert"
)

type PasswordsClientFixtureV1 struct {
	Client version1.IPasswordsClientV1
}

var USER_PWD = &version1.UserPasswordV1{
	Id:       "1",
	Password: "password123",
}

func NewPasswordsClientFixtureV1(client version1.IPasswordsClientV1) *PasswordsClientFixtureV1 {
	return &PasswordsClientFixtureV1{
		Client: client,
	}
}

func (c *PasswordsClientFixtureV1) clear() {
	c.Client.DeletePassword(context.Background(), "", USER_PWD.Id)
}

func (c *PasswordsClientFixtureV1) TestChangePassword(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create a new user
	err := c.Client.SetPassword(context.Background(), "", USER_PWD.Id, USER_PWD.Password)
	assert.Nil(t, err)

	// Recover password
	err = c.Client.RecoverPassword(context.Background(), "", USER_PWD.Id)
	assert.Nil(t, err)
}

func (c *PasswordsClientFixtureV1) TestRecoverPassword(t *testing.T) {
	c.clear()
	defer c.clear()

	// Sign up
	err := c.Client.SetPassword(context.Background(), "", USER_PWD.Id, USER_PWD.Password)
	assert.Nil(t, err)

	// Change password
	err = c.Client.ChangePassword(context.Background(), "", USER_PWD.Id, USER_PWD.Password, "xxx123")
	assert.Nil(t, err)

	// Sign in with new password
	auth, err1 := c.Client.Authenticate(context.Background(), "", USER_PWD.Id, "xxx123")
	assert.Nil(t, err1)
	assert.True(t, auth)

	// Delete password
	err = c.Client.DeletePassword(context.Background(), "", USER_PWD.Id)
	assert.Nil(t, err)
}

func (c *PasswordsClientFixtureV1) TestSigninWithWrongPassword(t *testing.T) {
	c.clear()
	defer c.clear()

	// Sign up
	err := c.Client.SetPassword(context.Background(), "", USER_PWD.Id, USER_PWD.Password)
	assert.Nil(t, err)

	// Sign in with wrong password
	auth, err := c.Client.Authenticate(context.Background(), "", USER_PWD.Id, "xxx")
	assert.NotNil(t, err)
	assert.False(t, auth)
}
