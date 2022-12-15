package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-passwords-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type passwordsCommandableHttpClientV1Test struct {
	client  *version1.PasswordsCommandableHttpClientV1
	fixture *PasswordsClientFixtureV1
}

func newPasswordsCommandableHttpClientV1Test() *passwordsCommandableHttpClientV1Test {
	return &passwordsCommandableHttpClientV1Test{}
}

func (c *passwordsCommandableHttpClientV1Test) setup(t *testing.T) *PasswordsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewPasswordsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewPasswordsClientFixtureV1(c.client)

	return c.fixture
}

func (c *passwordsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpRecoverPassword(t *testing.T) {
	c := newPasswordsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestRecoverPassword(t)
}

func TestHttpChangePassword(t *testing.T) {
	c := newPasswordsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestChangePassword(t)
}

func TestHttpSigninWithWrongPassword(t *testing.T) {
	c := newPasswordsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSigninWithWrongPassword(t)
}
