package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-passwords-go/version1"
)

type passwordsMockClientV1Test struct {
	client  *version1.PasswordsMockClientV1
	fixture *PasswordsClientFixtureV1
}

func newPasswordsMockClientV1Test() *passwordsMockClientV1Test {
	return &passwordsMockClientV1Test{}
}

func (c *passwordsMockClientV1Test) setup(t *testing.T) *PasswordsClientFixtureV1 {

	c.client = version1.NewPasswordsMockClientV1()
	c.fixture = NewPasswordsClientFixtureV1(c.client)
	return c.fixture
}

func (c *passwordsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMockRecoverPassword(t *testing.T) {
	c := newPasswordsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestRecoverPassword(t)
}

func TestMockChangePassword(t *testing.T) {
	c := newPasswordsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestChangePassword(t)
}

func TestMockSigninWithWrongPassword(t *testing.T) {
	c := newPasswordsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSigninWithWrongPassword(t)
}
