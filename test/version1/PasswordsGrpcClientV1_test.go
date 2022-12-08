package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-passwords-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type passwordsGrpcCommandableClientV1Test struct {
	client  *version1.PasswordGrpcClientV1
	fixture *PasswordsClientFixtureV1
}

func newPasswordsGrpcCommandableClientV1Test() *passwordsGrpcCommandableClientV1Test {
	return &passwordsGrpcCommandableClientV1Test{}
}

func (c *passwordsGrpcCommandableClientV1Test) setup(t *testing.T) *PasswordsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewPasswordGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewPasswordsClientFixtureV1(c.client)

	return c.fixture
}

func (c *passwordsGrpcCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestRecoverPassword(t *testing.T) {
	c := newPasswordsGrpcCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestRecoverPassword(t)
}

func TestChangePassword(t *testing.T) {
	c := newPasswordsGrpcCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestChangePassword(t)
}

func TestSigninWithWrongPassword(t *testing.T) {
	c := newPasswordsGrpcCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSigninWithWrongPassword(t)
}
