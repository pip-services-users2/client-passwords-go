package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-passwords-go/version1"
)

type passwordsMemoryClientV1Test struct {
	client  *version1.PasswordsMemoryClientV1
	fixture *PasswordsClientFixtureV1
}

func newPasswordsMemoryClientV1Test() *passwordsMemoryClientV1Test {
	return &passwordsMemoryClientV1Test{}
}

func (c *passwordsMemoryClientV1Test) setup(t *testing.T) *PasswordsClientFixtureV1 {

	c.client = version1.NewPasswordsMemoryClientV1()
	c.fixture = NewPasswordsClientFixtureV1(c.client)
	return c.fixture
}

func (c *passwordsMemoryClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMemoryRecoverPassword(t *testing.T) {
	c := newPasswordsMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestRecoverPassword(t)
}

func TestMemoryChangePassword(t *testing.T) {
	c := newPasswordsMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestChangePassword(t)
}

func TestMemorySigninWithWrongPassword(t *testing.T) {
	c := newPasswordsMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSigninWithWrongPassword(t)
}
