package build

import (
	clients1 "github.com/pip-services-users2/client-passwords-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type PasswordsClientFactory struct {
	cbuild.Factory
}

func NewPasswordsClientFactory() *PasswordsClientFactory {
	c := &PasswordsClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("pip-services-sasswords", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("pip-services-sasswords", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-passwords", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-passwords", "client", "grpc", "*", "1.0")
	memoryClientDescriptor := cref.NewDescriptor("service-passwords", "client", "memory", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewPasswordsNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewPasswordsDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewPasswordsCommandableHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewPasswordGrpcClientV1)
	c.RegisterType(memoryClientDescriptor, clients1.NewPasswordsMemoryClientV1)

	return c
}
